package ordered_queue

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json/v2"
	"errors"
	"hash/fnv"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// blpopTimeout 空轮超时，必须有限。用 0 永久阻塞会让 Close() 的 wg.Wait 挂到下一条消息才醒，
	// 给 2s 就把「感知停机」的代价钉死在 2s 内。
	blpopTimeout = 2 * time.Second
	// leaseTTL 分片租约时长：持有者进程崩溃后，其他实例最迟 TTL 接管该分片。
	leaseTTL = 15 * time.Second
	// leaseRenew 续租间隔，取 TTL/3，留出连续两次失败的重试余量。
	leaseRenew = leaseTTL / 3
	// reclaimWait 抢锁失败后的退避。租约剩余最长为 TTL，取一半折中：既不让接管干等满 TTL，也不空转打 Redis。
	reclaimWait = leaseTTL / 2
)

// 续租与释放都先校验持有者，防止「本地以为还持有、实际已被别家抢走」时误伤别人的租约。
// 单键脚本，不碰 Redis 4.0+ 才有的语法，也不依赖 unpack 之外的 Lua 特性。
var (
	renewScript = redis.NewScript(`
if redis.call('GET', KEYS[1]) == ARGV[1] then
    return redis.call('PEXPIRE', KEYS[1], ARGV[2])
end
return 0
`)
	releaseScript = redis.NewScript(`
if redis.call('GET', KEYS[1]) == ARGV[1] then
    return redis.call('DEL', KEYS[1])
end
return 0
`)
)

// shardResult consumeShard 的退出原因，决定认领协程是收摊还是重新抢锁。
type shardResult int

const (
	shardStop shardResult = iota // ctx 取消或 Close()：本实例不再消费任何分片
	shardLost                    // 租约易主/续租失败：交回认领循环，等下一轮再抢
)

// redisBackend 分片有序队列：RPUSH 右进 + BLPOP 左出合成 FIFO；
// 每分片一把租约，全局只有持有者一人串行消费，故同分片键严格有序。
type redisBackend struct {
	client *redis.Client
	prefix string
	shards int
	// holder 本实例的租约持有者标识，续租/释放按它做归属校验。
	holder string

	// stopCh：关闭后停止取新消息（已在途的仍处理完）。由 Close() 关闭。
	stopCh chan struct{}
	// wg：统计在途（含进行中的 Enqueue 与处理中的消息），Close() 等待其归零。
	wg sync.WaitGroup
	// closeOnce：保证 stopCh 只关闭一次。
	closeOnce sync.Once

	mu   sync.Mutex
	live map[string]bool // 已启动消费的 name：同一 name 起两套消费者会立刻打乱顺序，必须拒掉
}

func newRedisBackend(client *redis.Client, prefix string, shards int) *redisBackend {
	return &redisBackend{
		client: client,
		prefix: prefix,
		shards: shards,
		holder: newHolderID(),
		stopCh: make(chan struct{}),
		live:   make(map[string]bool, 4),
	}
}

// newHolderID 随机 16 字节 hex 作为持有者标识。
// 不复用 util/redis_lock：util 依赖 core，本包被 core 引用会成循环依赖。
func newHolderID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "holder-unavailable"
	}
	return hex.EncodeToString(b)
}

func (b *redisBackend) shardKey(name string, idx int) string {
	return b.prefix + name + ":" + strconv.Itoa(idx)
}

func (b *redisBackend) leaseKey(name string, idx int) string {
	return b.prefix + name + ":lease:" + strconv.Itoa(idx)
}

// shardOf FNV-1a(32) 取模：输入相同则结果恒定。
// 不能用 Go 内置 maphash——它每次进程启动换随机种子，重启后同一设备会落到别的分片，
// 与仍在旧分片排队的消息当场断开顺序关系。
func shardOf(shardKey string, shards int) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(shardKey))
	return int(h.Sum32() % uint32(shards))
}

// Enqueue 按 shardKey 哈希投递到对应分片尾部。入队动作计入 wg，与 queue 同套路。
func (b *redisBackend) Enqueue(name string, shardKey string, payload any) error {
	bs, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	b.wg.Add(1)
	defer b.wg.Done()
	return b.client.RPush(context.Background(), b.shardKey(name, shardOf(shardKey, b.shards)), bs).Err()
}

// Consume 为队列 name 起 shards 个认领协程，每片一个；本实例实际拿到几片由抢锁结果决定。
func (b *redisBackend) Consume(ctx context.Context, name string, handler Handler) error {
	b.mu.Lock()
	if b.live[name] {
		b.mu.Unlock()
		return errors.New("ordered_queue: 队列 " + strconv.Quote(name) + " 已在本实例启动消费，重复调用会产生第二个消费者并破坏有序性")
	}
	b.live[name] = true
	b.mu.Unlock()

	for idx := 0; idx < b.shards; idx++ {
		go b.claimLoop(ctx, name, idx, handler)
	}
	return nil
}

// claimLoop 单分片的「认领 → 串行消费 → 失锁重抢」循环。
// 抢不到锁不报错：那说明这片正被其他实例持有，退避后再排队，实例间自动分摊。
func (b *redisBackend) claimLoop(ctx context.Context, name string, idx int, handler Handler) {
	lk := b.leaseKey(name, idx)
	for {
		if b.stopped(ctx) {
			return
		}
		ok, err := b.client.SetNX(ctx, lk, b.holder, leaseTTL).Result()
		if err != nil {
			if b.stopped(ctx) {
				return
			}
			log.Printf("[ordered_queue:redis] 认领分片 %s#%d 失败(将重试): %v", name, idx, err)
			if sleepWhile(ctx, b.stopCh, leaseRenew) {
				return
			}
			continue
		}
		if !ok {
			if sleepWhile(ctx, b.stopCh, reclaimWait) {
				return
			}
			continue
		}

		res := b.consumeShard(ctx, name, idx, handler)
		// 无论失锁还是停机都尝试交回租约（脚本校验归属，不是自己的锁不会误删）；
		// 主动释放让别的实例不必等满 TTL 就能接管。
		b.release(lk)
		if res == shardStop {
			return
		}
	}
}

// consumeShard 串行消费一个分片：取一条、处理一条，处理期间不取下一条。
// 「一片一消费者」与「取一条处理一条」两条缺一不可，少任何一条都会退化成乱序。
// 续租由 startRenewer 独立协程按固定节奏执行，不依赖本循环推进：否则一条卡过
// leaseTTL 的慢消息会让租约到期被别家抢走，同片出现两个消费者而当场断序。
func (b *redisBackend) consumeShard(ctx context.Context, name string, idx int, handler Handler) shardResult {
	sk := b.shardKey(name, idx)
	lk := b.leaseKey(name, idx)

	lost := make(chan struct{})
	stopRenew := b.startRenewer(ctx, name, idx, lk, lost)
	defer stopRenew()

	for {
		select {
		case <-lost: // 续租协程判定租约已失，本协程交回认领循环
			return shardLost
		default:
		}
		if b.stopped(ctx) {
			return shardStop
		}

		res, err := b.client.BLPop(ctx, blpopTimeout, sk).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				continue // 空轮超时：无消息，回循环查停机位与失锁位
			}
			if b.stopped(ctx) {
				return shardStop
			}
			log.Printf("[ordered_queue:redis] 读取分片 %s#%d 失败(将重试): %v", name, idx, err)
			if sleepWhile(ctx, b.stopCh, 100*time.Millisecond) { // 网络抖动退避
				return shardStop
			}
			continue
		}
		if len(res) < 2 {
			continue
		}
		// 取出即出队，即使本协程已在失锁临界也必须处理完这一条，否则消息凭空丢失
		b.handleOne(ctx, name, idx, []byte(res[1]), handler)
	}
}

// startRenewer 起一个按 leaseRenew 续租的后台协程：一旦续租失败或租约易主，就关闭
// lost 通知消费循环让位。返回的 stop 用于结束该协程（consumeShard 退出时必调）。
// 残留窗口只剩「发现失锁前最长 blpopTimeout + 一条在途 handler」，且失锁后不再取新消息。
func (b *redisBackend) startRenewer(ctx context.Context, name string, idx int, lk string, lost chan struct{}) func() {
	renewCtx, cancel := context.WithCancel(ctx)
	var once sync.Once
	flagLost := func() { once.Do(func() { close(lost) }) }

	go func() {
		t := time.NewTicker(leaseRenew)
		defer t.Stop()
		for {
			select {
			case <-renewCtx.Done():
				return
			case <-t.C:
				ok, err := b.renew(renewCtx, lk)
				if err != nil {
					log.Printf("[ordered_queue:redis] 分片 %s#%d 续租失败(让位): %v", name, idx, err)
					flagLost()
					return
				}
				if !ok {
					log.Printf("[ordered_queue:redis] 分片 %s#%d 租约已易主，停止消费", name, idx)
					flagLost()
					return
				}
			}
		}
	}()
	return cancel
}

// handleOne 在本协程内就地处理（本协程即该分片唯一 worker），逐条 recover，单条失败不拖垮消费者。
// 取出即出队：处理中进程崩溃会丢这一条，与 queue 同为 at-most-once 语义。
func (b *redisBackend) handleOne(ctx context.Context, name string, idx int, body []byte, handler Handler) {
	b.wg.Add(1)
	defer b.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ordered_queue:redis] handle %s#%d panicked: %v", name, idx, r)
		}
	}()
	if err := handler(ctx, body); err != nil {
		log.Printf("[ordered_queue:redis] handle message from %s#%d failed: %v", name, idx, err)
	}
}

func (b *redisBackend) renew(ctx context.Context, lk string) (bool, error) {
	n, err := renewScript.Run(ctx, b.client, []string{lk}, b.holder, leaseTTL.Milliseconds()).Int64()
	if err != nil {
		return false, err
	}
	return n == 1, nil
}

// release 主动交回租约。停机时 Redis 可能已在关连接，故用独立 ctx 并允许失败（TTL 兜底）。
func (b *redisBackend) release(lk string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if _, err := releaseScript.Run(ctx, b.client, []string{lk}, b.holder).Result(); err != nil {
		log.Printf("[ordered_queue:redis] 释放租约 %s 失败(由 TTL 兜底): %v", lk, err)
	}
}

func (b *redisBackend) stopped(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	case <-b.stopCh:
		return true
	default:
		return false
	}
}

// sleepWhile 可被停机打断的等待：返回 true 表示该收摊了。
func sleepWhile(ctx context.Context, stopCh chan struct{}, d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return true
	case <-stopCh:
		return true
	case <-timer.C:
		return false
	}
}

// Close 停取新消息并等待在途处理完。消费者最多再空转 blpopTimeout 即退出，Close 不会长挂。
// 分片 List 里没取走的消息不丢（留在 Redis），租约释放后由其他实例接着消费。客户端由 core 管理，不在此关闭。
func (b *redisBackend) Close() error {
	b.closeOnce.Do(func() {
		close(b.stopCh)
	})
	b.wg.Wait()
	return nil
}
