package captcha

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/chenyuIT/framework/contracts/captcha"
	"github.com/chenyuIT/framework/facades"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	dpi        = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	r          = rand.New(rand.NewSource(time.Now().UnixNano()))
	fontFamily = make([]string, 0)
)

const txtChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

const (
	//图片格式
	ImageFormatPng ImageFormat = iota
	ImageFormatJpeg
	ImageFormatGif
)
const (
	//验证码噪点强度
	CaptchaComplexLower captcha.CaptchaComplex = iota
	CaptchaComplexMedium
	CaptchaComplexHigh
)

type CaptchaImage struct {
	nrgba       *image.NRGBA
	width       int
	height      int
	imageFormat ImageFormat
	Complex     int
	Error       error
}

type ImageFormat int

func NewCaptcha() *CaptchaImage {
	ReadFonts("fonts", "ttf")

	width := facades.Config.GetInt("captcha.width")
	height := facades.Config.GetInt("captcha.height")
	bgColor := color.RGBA{255, 0, 255, 255}

	captcha := _NewCaptchaImage(width, height, bgColor)
	return captcha
}

func (captcha *CaptchaImage) NewCaptchaImage() (img string, err error) {
	facades.Captcha.ClearImage()
	//画上三条随机直线
	facades.Captcha.DrawLine(3)

	//画边框
	facades.Captcha.DrawBorder(ColorToRGB(0x17A7A7A))

	//画随机噪点
	facades.Captcha.DrawNoise(CaptchaComplexHigh)

	//画随机文字噪点
	facades.Captcha.DrawTextNoise(CaptchaComplexLower)
	//画验证码文字，可以预先保持到Session种或其他储存容器种
	facades.Captcha.DrawText(RandText(4))
	if err != nil {
		fmt.Println(err)
	}
	// 将验证码图片转成base64串
	_img := facades.Captcha.SaveImage()

	return _img, nil
}

// 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ReadFonts(dirPth string, suffix string) (err error) {
	files := make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	SetFontFamily(files...)
	return nil
}

// 新建一个图片对象
func _NewCaptchaImage(width int, height int, bgColor color.RGBA) *CaptchaImage {

	m := image.NewNRGBA(image.Rect(0, 0, width, height))

	draw.Draw(m, m.Bounds(), &image.Uniform{C: bgColor}, image.ZP, draw.Src)

	imgFormat := facades.Config.GetInt("captcha.imageFormat")
	imageFormat := ImageFormat(imgFormat)
	return &CaptchaImage{
		nrgba:       m,
		height:      height,
		width:       width,
		imageFormat: imageFormat,
	}
}

// 清除一个图片对象内容
func (captcha *CaptchaImage) ClearImage() {
	width := facades.Config.GetInt("captcha.width")
	height := facades.Config.GetInt("captcha.height")
	bgColor := color.RGBA{255, 0, 255, 255}
	m := image.NewNRGBA(image.Rect(0, 0, width, height))

	draw.Draw(m, m.Bounds(), &image.Uniform{C: bgColor}, image.ZP, draw.Src)

	captcha.nrgba = m
}

// 保存图片对象
func (captcha *CaptchaImage) SaveImage() (img string) {
	buffer := new(bytes.Buffer)
	if captcha.imageFormat == ImageFormatPng {
		png.Encode(buffer, captcha.nrgba)
	}
	if captcha.imageFormat == ImageFormatJpeg {
		jpeg.Encode(buffer, captcha.nrgba, &jpeg.Options{Quality: 100})
	}
	if captcha.imageFormat == ImageFormatGif {
		gif.Encode(buffer, captcha.nrgba, &gif.Options{NumColors: 256})
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

// 添加一个较粗的空白直线
func (captcha *CaptchaImage) DrawHollowLine() {
	if captcha.Error != nil {
		return
	}
	first := captcha.width / 20
	end := first * 19

	lineColor := color.RGBA{R: 245, G: 250, B: 251, A: 255}

	x1 := float64(r.Intn(first))

	x2 := float64(r.Intn(first) + end)

	multiple := float64(r.Intn(5)+3) / float64(5)
	if int(multiple*10)%3 == 0 {
		multiple = multiple * -1.0
	}

	w := captcha.height / 20

	for ; x1 < x2; x1++ {

		y := math.Sin(x1*math.Pi*multiple/float64(captcha.width)) * float64(captcha.height/3)

		if multiple < 0 {
			y = y + float64(captcha.height/2)
		}
		captcha.nrgba.Set(int(x1), int(y), lineColor)

		for i := 0; i <= w; i++ {
			captcha.nrgba.Set(int(x1), int(y)+i, lineColor)
		}
	}

	return
}

// 画一条曲线.
func (captcha *CaptchaImage) DrawSineLine() {
	if captcha.Error != nil {
		return
	}
	px := 0
	var py float64 = 0

	//振幅
	a := r.Intn(captcha.height / 2)

	//Y轴方向偏移量
	b := Random(int64(-captcha.height/4), int64(captcha.height/4))

	//X轴方向偏移量
	f := Random(int64(-captcha.height/4), int64(captcha.height/4))
	// 周期
	var t float64 = 0
	if captcha.height > captcha.width/2 {
		t = Random(int64(captcha.width/2), int64(captcha.height))
	} else {
		t = Random(int64(captcha.height), int64(captcha.width/2))
	}
	w := float64((2 * math.Pi) / t)

	// 曲线横坐标起始位置
	px1 := 0
	px2 := int(Random(int64(float64(captcha.width)*0.8), int64(captcha.width)))

	c := color.RGBA{R: uint8(r.Intn(150)), G: uint8(r.Intn(150)), B: uint8(r.Intn(150)), A: uint8(255)}

	for px = px1; px < px2; px++ {
		if w != 0 {
			py = float64(a)*math.Sin(w*float64(px)+f) + b + (float64(captcha.width) / float64(5))
			i := captcha.height / 5
			for i > 0 {
				captcha.nrgba.Set(px+i, int(py), c)
				i--
			}
		}
	}

	return
}

// 画一条直线.
func (_captcha *CaptchaImage) DrawLine(num int) {
	if _captcha.Error != nil {
		return
	}
	first := _captcha.width / 10
	end := first * 9

	y := _captcha.height / 3

	for i := 0; i < num; i++ {

		point1 := captcha.Point{X: r.Intn(first), Y: r.Intn(y)}
		point2 := captcha.Point{X: r.Intn(first) + end, Y: r.Intn(y)}

		if i%2 == 0 {
			point1.Y = r.Intn(y) + y*2
			point2.Y = r.Intn(y)
		} else {
			point1.Y = r.Intn(y) + y*(i%2)
			point2.Y = r.Intn(y) + y*2
		}

		_captcha.DrawBeeline(point1, point2, randDeepColor())

	}
	return
}

// 画直线.
func (captcha *CaptchaImage) DrawBeeline(point1 captcha.Point, point2 captcha.Point, lineColor color.RGBA) {
	if captcha.Error != nil {
		return
	}
	dx := math.Abs(float64(point1.X - point2.X))

	dy := math.Abs(float64(point2.Y - point1.Y))
	sx, sy := 1, 1
	if point1.X >= point2.X {
		sx = -1
	}
	if point1.Y >= point2.Y {
		sy = -1
	}
	err := dx - dy
	//循环的画点直到到达结束坐标停止.
	for {
		captcha.nrgba.Set(point1.X, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+2, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-2, point1.Y, lineColor)
		if point1.X == point2.X && point1.Y == point2.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			point1.X += sx
		}
		if e2 < dx {
			err += dx
			point1.Y += sy
		}
	}
	return
}

// 画边框.
func (captcha *CaptchaImage) DrawBorder(borderColor color.RGBA) {
	if captcha.Error != nil {
		return
	}
	for x := 0; x < captcha.width; x++ {
		captcha.nrgba.Set(x, 0, borderColor)
		captcha.nrgba.Set(x, captcha.height-1, borderColor)
	}
	for y := 0; y < captcha.height; y++ {
		captcha.nrgba.Set(0, y, borderColor)
		captcha.nrgba.Set(captcha.width-1, y, borderColor)
	}
	return
}

// 画噪点.
func (captcha *CaptchaImage) DrawNoise(complex captcha.CaptchaComplex) {
	if captcha.Error != nil {
		return
	}
	density := 18
	if complex == CaptchaComplexLower {
		density = 28
	} else if complex == CaptchaComplexMedium {
		density = 18
	} else if complex == CaptchaComplexHigh {
		density = 8
	}
	maxSize := (captcha.height * captcha.width) / density

	for i := 0; i < maxSize; i++ {

		rw := r.Intn(captcha.width)
		rh := r.Intn(captcha.height)

		captcha.nrgba.Set(rw, rh, randColor())
		size := r.Intn(maxSize)
		if size%3 == 0 {
			captcha.nrgba.Set(rw+1, rh+1, randColor())
		}
	}
	return
}

// 画文字噪点.
func (captcha *CaptchaImage) DrawTextNoise(complex captcha.CaptchaComplex) {
	if captcha.Error != nil {
		return
	}
	density := 1500
	if complex == CaptchaComplexLower {
		density = 2000
	} else if complex == CaptchaComplexMedium {
		density = 1500
	} else if complex == CaptchaComplexHigh {
		density = 1000
	}

	maxSize := (captcha.height * captcha.width) / density

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	c := freetype.NewContext()
	c.SetDPI(*dpi)

	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)
	rawFontSize := float64(captcha.height) / (1 + float64(r.Intn(7))/float64(10))

	for i := 0; i < maxSize; i++ {

		rw := r.Intn(captcha.width)
		rh := r.Intn(captcha.height)

		text := RandText(1)
		fontSize := rawFontSize/2 + float64(r.Intn(5))

		c.SetSrc(image.NewUniform(RandLightColor()))
		c.SetFontSize(fontSize)
		f, err := RandFontFamily()

		if err != nil {
			captcha.Error = err
			return
		}
		c.SetFont(f)
		pt := freetype.Pt(rw, rh)

		_, err = c.DrawString(text, pt)
		if err != nil {
			captcha.Error = err
			return
		}
	}
	return
}

// 写字.
func (captcha *CaptchaImage) DrawText(text string) {
	if captcha.Error != nil {
		return
	}
	c := freetype.NewContext()
	c.SetDPI(*dpi)

	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)

	fontWidth := captcha.width / len(text)

	for i, s := range text {

		fontSize := float64(captcha.height) / (1 + float64(r.Intn(7))/float64(9))

		c.SetSrc(image.NewUniform(randDeepColor()))
		c.SetFontSize(fontSize)
		f, err := RandFontFamily()

		if err != nil {
			captcha.Error = err
			return
		}
		c.SetFont(f)

		x := int(fontWidth)*i + int(fontWidth)/int(fontSize)

		y := 5 + r.Intn(captcha.height/2) + int(fontSize/2)

		pt := freetype.Pt(x, y)

		_, err = c.DrawString(string(s), pt)
		if err != nil {
			captcha.Error = err
			return
		}
		//pt.Y += c.PointToFixed(*size * *spacing)
		//pt.X += c.PointToFixed(*size);
	}
	return
}

// 获取所及字体.
func RandFontFamily() (*truetype.Font, error) {
	fontFile := fontFamily[r.Intn(len(fontFamily))]

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return &truetype.Font{}, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return &truetype.Font{}, err
	}
	return f, nil
}

// 随机生成深色系.
func randDeepColor() color.RGBA {

	randColor := randColor()

	increase := float64(30 + r.Intn(255))

	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))

	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

// 随机生成浅色.
func RandLightColor() color.RGBA {

	red := r.Intn(55) + 200
	green := r.Intn(55) + 200
	blue := r.Intn(55) + 200

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

// 生成随机颜色.
func randColor() color.RGBA {

	red := r.Intn(255)
	green := r.Intn(255)
	blue := r.Intn(255)
	if (red + green) > 400 {
		blue = 0
	} else {
		blue = 400 - green - red
	}
	if blue > 255 {
		blue = 255
	}
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

// 生成随机字体.
func RandText(num int) string {
	textNum := len(txtChars)
	text := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < num; i++ {
		text = text + string(txtChars[r.Intn(textNum)])
	}
	return text
}

// 添加一个字体路径到字体库.
func SetFontFamily(fontPath ...string) {
	fontFamily = append(fontFamily, fontPath...)
}

// 颜色代码转换为RGB
// input int
// output int red, green, blue.
func ColorToRGB(colorVal int) color.RGBA {

	red := colorVal >> 16
	green := (colorVal & 0x00FF00) >> 8
	blue := colorVal & 0x0000FF

	return color.RGBA{
		R: uint8(red),
		G: uint8(green),
		B: uint8(blue),
		A: uint8(255),
	}
}
