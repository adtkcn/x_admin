package captcha_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"log"
	"x_admin/plugin/aj-captcha-go/captcha_config"
	"x_admin/plugin/aj-captcha-go/model/vo"
	"x_admin/plugin/aj-captcha-go/util"
	img "x_admin/plugin/aj-captcha-go/util/image"
)

const (
	TEXT = "的一了是我不在人们有来他这上着个地到大里说就去子得也和那要下看天时过出小么起你都把好还多没为又可家学只以主会样年想生同老中十从自面前头道它后然走很像见两用她国动进成回什边作对开而己些现山民候经发工向事命给长水几义三声于高手知理眼志点心战二问但身方实吃做叫当住听革打呢真全才四已所敌之最光产情路分总条白话东席次亲如被花口放儿常气五第使写军吧文运再果怎定许快明行因别飞外树物活部门无往船望新带队先力完却站代员机更九您每风级跟笑啊孩万少直意夜比阶连车重便斗马哪化太指变社似士者干石满日决百原拿群究各六本思解立河村八难早论吗根共让相研今其书坐接应关信觉步反处记将千找争领或师结块跑谁草越字加脚紧爱等习阵怕月青半火法题建赶位唱海七女任件感准张团屋离色脸片科倒睛利世刚且由送切星导晚表够整认响雪流未场该并底深刻平伟忙提确近亮轻讲农古黑告界拉名呀土清阳照办史改历转画造嘴此治北必服雨穿内识验传业菜爬睡兴形量咱观苦体众通冲合破友度术饭公旁房极南枪读沙岁线野坚空收算至政城劳落钱特围弟胜教热展包歌类渐强数乡呼性音答哥际旧神座章帮啦受系令跳非何牛取入岸敢掉忽种装顶急林停息句区衣般报叶压慢叔背细"
)

func NewClickWordCaptchaService(factory *CaptchaServiceFactory) *ClickWordCaptchaService {
	img.SetUp()
	return &ClickWordCaptchaService{factory: factory}
}

type ClickWordCaptchaService struct {
	factory *CaptchaServiceFactory
}

// 校验点击文字验证码
func (c *ClickWordCaptchaService) Check(token string, pointJson string) error {
	cache, err := c.factory.GetCache()
	if err != nil {
		return err
	}
	codeKey := fmt.Sprintf(captcha_config.CodeKeyPrefix, token)

	cachePointInfo := cache.Get(codeKey)

	if cachePointInfo == "" {
		return errors.New("验证码已失效")
	}

	// 解析结构体
	var cachePoint []vo.PointVO

	var userPoint []vo.PointVO

	err = json.Unmarshal([]byte(cachePointInfo), &cachePoint)

	if err != nil {
		return err
	}

	// 解密前端传递过来的数据
	userPointJson := util.AesDecrypt(pointJson, cachePoint[0].SecretKey)

	err = json.Unmarshal([]byte(userPointJson), &userPoint)

	if err != nil {
		return err
	}

	// 校验用户提交的点数是否与期望数量一致
	if len(userPoint) < len(cachePoint) {
		cache.Delete(codeKey)
		return errors.New("验证失败：点击数量不足")
	}

	fontSize := c.factory.config.ClickWord.FontSize
	for i, pointVO := range cachePoint {
		userTargetPoint := userPoint[i]
		startX := pointVO.X - c.factory.config.ClickWord.XOffset
		endX := pointVO.X + fontSize + c.factory.config.ClickWord.XOffset

		startY := pointVO.Y - c.factory.config.ClickWord.YOffset
		endY := pointVO.Y + fontSize + c.factory.config.ClickWord.YOffset
		if userTargetPoint.X >= startX && userTargetPoint.X <= endX && userTargetPoint.Y >= startY && userTargetPoint.Y <= endY {

		} else {
			cache.Delete(codeKey)
			return errors.New("验证失败")
		}
	}

	return nil
}

// 校验点击文字验证码-并删除
func (c *ClickWordCaptchaService) Verification(token string, pointJson string) error {
	err := c.Check(token, pointJson)
	if err != nil {
		return err
	}
	codeKey := fmt.Sprintf(captcha_config.CodeKeyPrefix, token)
	if cache, err := c.factory.GetCache(); err == nil {
		cache.Delete(codeKey)
	}
	return nil
}

func (c *ClickWordCaptchaService) Get() (map[string]any, error) {
	// 初始化背景图片
	backgroundImage, err := img.GetClickBackgroundImage()
	if err != nil {
		return nil, err
	}
	// 为背景图片设置水印
	if c.factory.config.Watermark.Text != "" {
		backgroundImage.SetText(c.factory.config.Watermark.Text, c.factory.config.Watermark.FontSize, c.factory.config.Watermark.Color)
	}
	pointList, wordList, err := c.getImageData(backgroundImage)
	if err != nil {
		return nil, err
	}

	originalImageBase64, err := backgroundImage.Base64()

	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	data["originalImageBase64"] = originalImageBase64
	data["wordList"] = wordList
	data["secretKey"] = pointList[0].SecretKey
	data["token"] = util.GetUuid()

	codeKey := fmt.Sprintf(captcha_config.CodeKeyPrefix, data["token"])
	jsonPoint, err := json.Marshal(pointList)
	if err != nil {
		log.Printf("point json Marshal err: %v", err)
		return nil, err
	}

	cache, err := c.factory.GetCache()
	if err != nil {
		return nil, err
	}
	cache.Set(codeKey, string(jsonPoint), c.factory.config.CacheExpireSec)
	return data, nil
}

func (c *ClickWordCaptchaService) getImageData(image *util.ImageUtil) ([]vo.PointVO, []string, error) {
	cfg := c.factory.config.ClickWord
	fontNum := cfg.FontNum
	interferenceNum := cfg.InterferenceFontNum
	totalNum := fontNum + interferenceNum

	// 1. 选取点击文字
	clickWords := c.getRandomWords(fontNum)

	// 2. 选取干扰文字（排除已选中的点击文字，避免重复）
	interferenceWords := c.getRandomWordsExcluding(interferenceNum, clickWords)

	// 3. 网格均匀分布：根据图片尺寸计算网格
	allWords := make([]string, 0, totalNum)
	allWords = append(allWords, clickWords...)
	allWords = append(allWords, interferenceWords...)

	points := c.generateGridPoints(image.Width, image.Height, totalNum, cfg.FontSize)

	// 4. 打乱点位顺序（让点击文字不总是出现在固定位置）
	c.shufflePoints(points)

	// 构建本次的 secret
	key := util.RandString(16)

	var pointList []vo.PointVO
	var wordList []string

	// 5. 绘制点击文字（大字体、醒目颜色）
	for i := 0; i < fontNum; i++ {
		fontSize := util.RandomInt(cfg.FontSize-2, cfg.FontSize+1)
		points[i].SetSecretKey(key)

		clr := c.clickColor()
		if err := image.SetArtTextWithColor(allWords[i], fontSize, points[i], clr); err != nil {
			return nil, nil, err
		}

		pointList = append(pointList, points[i])
		wordList = append(wordList, allWords[i])
	}

	// 6. 绘制干扰文字（小字体、淡色）
	interferenceFontSize := cfg.InterferenceFontSize
	if interferenceFontSize <= 0 {
		interferenceFontSize = cfg.FontSize * 2 / 3 // 默认点击文字的 2/3
	}
	for i := fontNum; i < len(allWords); i++ {
		fs := util.RandomInt(interferenceFontSize-2, interferenceFontSize+1)
		points[i].SetSecretKey(key)

		clr := c.interferenceColor()
		if err := image.SetArtTextWithColor(allWords[i], fs, points[i], clr); err != nil {
			return nil, nil, err
		}
	}

	return pointList, wordList, nil
}

// getRandomWords 获取 count 个不重复的随机文字
func (c *ClickWordCaptchaService) getRandomWords(count int) []string {
	runesArray := []rune(TEXT)
	size := len(runesArray)

	if count > size {
		count = size
	}

	set := make(map[string]bool)
	var wordList []string

	for attempt := 0; attempt < count*3 && len(set) < count; attempt++ {
		word := runesArray[util.RandomInt(0, size)]
		set[string(word)] = true
	}
	for str := range set {
		wordList = append(wordList, str)
	}
	return wordList
}

// getRandomWordsExcluding 获取 count 个随机文字，排除 exclude 中已有的文字
func (c *ClickWordCaptchaService) getRandomWordsExcluding(count int, exclude []string) []string {
	excludeSet := make(map[string]bool, len(exclude))
	for _, w := range exclude {
		excludeSet[w] = true
	}

	runesArray := []rune(TEXT)
	size := len(runesArray)

	set := make(map[string]bool)
	var wordList []string

	for attempt := 0; attempt < count*5 && len(set) < count; attempt++ {
		word := string(runesArray[util.RandomInt(0, size)])
		if excludeSet[word] || set[word] {
			continue
		}
		set[word] = true
		wordList = append(wordList, word)
	}
	return wordList
}

// generateGridPoints 基于最小距离约束的随机采样，生成自然分布的点位
func (c *ClickWordCaptchaService) generateGridPoints(imgWidth, imgHeight, count, fontSize int) []vo.PointVO {
	padding := fontSize
	xMin, xMax := padding, imgWidth-fontSize-padding
	yMin, yMax := padding, imgHeight-fontSize-padding

	if xMax <= xMin {
		xMax = xMin + 1
	}
	if yMax <= yMin {
		yMax = yMin + 1
	}

	// 初始最小间距（根据图片和文字数量动态计算）
	baseMinDist := fontSize * 3 / 2
	points := make([]vo.PointVO, 0, count)

	for i := 0; i < count; i++ {
		placed := false
		// 逐步放宽间距约束，确保一定能放下
		for relax := 0; relax < 5; relax++ {
			minDist := baseMinDist * (5 - relax) / 5 // 100% → 80% → 60% → 40% → 20%
			maxTries := 60

			for try := 0; try < maxTries; try++ {
				x := util.RandomInt(xMin, xMax)
				y := util.RandomInt(yMin, yMax)

				// 检查与已有点位的最小距离
				tooClose := false
				for _, p := range points {
					dx := x - p.X
					dy := y - p.Y
					if dx*dx+dy*dy < minDist*minDist {
						tooClose = true
						break
					}
				}
				if !tooClose {
					points = append(points, vo.PointVO{X: x, Y: y})
					placed = true
					break
				}
			}
			if placed {
				break
			}
		}
		// 极端情况：强制放置
		if !placed {
			x := util.RandomInt(xMin, xMax)
			y := util.RandomInt(yMin, yMax)
			points = append(points, vo.PointVO{X: x, Y: y})
		}
	}

	return points
}

// shufflePoints 打乱点位顺序
func (c *ClickWordCaptchaService) shufflePoints(points []vo.PointVO) {
	for i := len(points) - 1; i > 0; i-- {
		j := util.RandomInt(0, i+1)
		points[i], points[j] = points[j], points[i]
	}
}

// clickColor 生成点击文字颜色（醒目深色）
func (c *ClickWordCaptchaService) clickColor() color.RGBA {
	return color.RGBA{
		R: uint8(util.RandomInt(10, 120)),
		G: uint8(util.RandomInt(10, 120)),
		B: uint8(util.RandomInt(10, 120)),
		A: 255,
	}
}

// interferenceColor 生成干扰文字颜色（浅色淡色）
func (c *ClickWordCaptchaService) interferenceColor() color.RGBA {
	return color.RGBA{
		R: uint8(util.RandomInt(160, 210)),
		G: uint8(util.RandomInt(160, 210)),
		B: uint8(util.RandomInt(160, 210)),
		A: 255,
	}
}
