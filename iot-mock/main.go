package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand" // 仅用于自动上报的数值抖动，非安全用途（凭证/令牌不使用）
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// iot-mock：向本项目设备侧 HTTP 接入层模拟上报，用于联调大屏/告警/物模型校验。
//
//	POST {base}/api/iot/telemetry   body {"values":{...},"ts":<毫秒>}
//	POST {base}/api/iot/heartbeat   body {"ts":<毫秒>}
//	凭证头 X-Device-Sn + X-Access-Token；产品接入协议须为 http（或留空）。
//
// 交互：终端输入数值回车 → 写入 -key 标识符并立即推送；hb 发送心跳；q 退出。
func main() {
	var (
		base     = flag.String("base", envOr("IOT_BASE", "http://127.0.0.1:18000"), "服务基地址")
		sn       = flag.String("sn", os.Getenv("IOT_SN"), "设备编号 X-Device-Sn")
		token    = flag.String("token", os.Getenv("IOT_TOKEN"), "设备接入凭证 X-Access-Token")
		key      = flag.String("key", "value", "输入数值时写入的属性标识符")
		values   = flag.String("values", "", "初始属性键值，如 temperature=26.3,humidity=55")
		interval = flag.Duration("interval", 0, "自动上报间隔，0=仅手动推送")
		jitter   = flag.Float64("jitter", 0, "自动模式下对 -key 值叠加的随机幅度(±绝对值)")
		withHB   = flag.Bool("heartbeat", false, "每次自动上报前额外发送一次心跳")
		timeout  = flag.Duration("timeout", 5*time.Second, "单次 HTTP 请求超时")
	)
	flag.Parse()

	if *sn == "" || *token == "" {
		fmt.Fprintln(os.Stderr, "必须提供 -sn 与 -token（或环境变量 IOT_SN / IOT_TOKEN）")
		flag.Usage()
		os.Exit(1)
	}

	m := &mock{
		base:     strings.TrimRight(*base, "/"),
		sn:       *sn,
		token:    *token,
		key:      *key,
		interval: *interval,
		jitter:   *jitter,
		hb:       *withHB,
		client:   &http.Client{Timeout: *timeout},
	}
	m.initValues(*values)

	logf("启动 base=%s sn=%s key=%s interval=%s", m.base, m.sn, m.key, m.interval)
	logf("命令：<数值> 设置并立即推送 | hb 心跳 | q 退出")

	stop := make(chan struct{})
	var once sync.Once
	stopOnce := func() { once.Do(func() { close(stop) }) }

	if m.interval > 0 {
		go m.autoLoop(stop)
	}
	m.readStdin(stopOnce)
	stopOnce()
}

// mock 持有一次设备模拟所需的全部配置与当前属性值。
type mock struct {
	base     string
	sn       string
	token    string
	key      string
	interval time.Duration
	jitter   float64
	hb       bool
	client   *http.Client

	mu     sync.Mutex
	values map[string]any
}

// initValues 解析 -values（k=v 逗号分隔），并保证 -key 存在（缺省 0）。
func (m *mock) initValues(s string) {
	m.values = map[string]any{}
	for _, pair := range strings.Split(s, ",") {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			continue
		}
		m.values[strings.TrimSpace(kv[0])] = parseScalar(strings.TrimSpace(kv[1]))
	}
	if _, ok := m.values[m.key]; !ok {
		m.values[m.key] = float64(0)
	}
}

func (m *mock) setValue(v float64) {
	m.mu.Lock()
	m.values[m.key] = v
	m.mu.Unlock()
}

// applyJitter 给 -key 数值叠加 ±jitter 的随机扰动。
func (m *mock) applyJitter() {
	if m.jitter <= 0 {
		return
	}
	m.mu.Lock()
	if cur, ok := m.values[m.key].(float64); ok {
		m.values[m.key] = cur + (rand.Float64()*2-1)*m.jitter
	}
	m.mu.Unlock()
}

func (m *mock) snapshot() map[string]any {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make(map[string]any, len(m.values))
	for k, v := range m.values {
		out[k] = v
	}
	return out
}

func (m *mock) sendTelemetry() {
	vals := m.snapshot()
	logf("[上报] %s", jsonCompact(vals))
	m.post("/api/iot/telemetry", map[string]any{"values": vals, "ts": time.Now().UnixMilli()})
}

func (m *mock) sendHeartbeat() {
	m.post("/api/iot/heartbeat", map[string]any{"ts": time.Now().UnixMilli()})
}

// post 发送一次设备请求并打印接入层应答。
func (m *mock) post(path string, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		logf("[构造请求失败] %v", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, m.base+path, bytes.NewReader(data))
	if err != nil {
		logf("[构造请求失败] %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Sn", m.sn)
	req.Header.Set("X-Access-Token", m.token)

	resp, err := m.client.Do(req)
	if err != nil {
		logf("[请求失败] %s -> %v", path, err)
		return
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)

	var r struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(raw, &r)
	if resp.StatusCode == http.StatusOK && r.Code == 200 {
		logf("[成功] %s %s", path, oneLine(string(raw)))
	} else {
		logf("[失败] %s http=%d %s", path, resp.StatusCode, oneLine(string(raw)))
	}
}

// autoLoop 按 interval 周期性自动上报。
func (m *mock) autoLoop(stop <-chan struct{}) {
	t := time.NewTicker(m.interval)
	defer t.Stop()
	for {
		select {
		case <-stop:
			return
		case <-t.C:
			if m.hb {
				m.sendHeartbeat()
			}
			m.applyJitter()
			m.sendTelemetry()
		}
	}
}

// readStdin 读取交互命令，阻塞直到退出或 EOF。
func (m *mock) readStdin(stop func()) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		switch line {
		case "":
			continue
		case "q", "quit", "exit":
			stop()
			return
		case "hb", "heartbeat":
			m.sendHeartbeat()
		default:
			if v, err := strconv.ParseFloat(line, 64); err == nil {
				m.setValue(v)
				m.sendTelemetry()
			} else {
				logf("[输入无效] %q（数值 / hb / q）", line)
			}
		}
	}
	stop()
}

var stdoutLog = log.New(os.Stdout, "", log.Ltime)

func logf(format string, a ...any) { stdoutLog.Printf(format, a...) }

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

// parseScalar 数值优先解析为 float64，否则保留为字符串。
func parseScalar(s string) any {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return s
}

func jsonCompact(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(b)
}

func oneLine(s string) string { return strings.Join(strings.Fields(s), " ") }
