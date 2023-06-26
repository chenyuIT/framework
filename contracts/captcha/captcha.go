package captcha

import (
	"image/color"
)

type CaptchaComplex int

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{X: x, Y: y}
}

type Captcha interface {
	NewCaptchaImage() (img string, err error)

	ClearImage()
	DrawHollowLine()
	DrawSineLine()
	DrawLine(num int)
	DrawBeeline(point1 Point, point2 Point, lineColor color.RGBA)
	DrawBorder(borderColor color.RGBA)
	DrawNoise(complex CaptchaComplex)
	DrawTextNoise(complex CaptchaComplex)
	DrawText(text string)
	SaveImage() (img string)
}
