package ocr

import (
	"github.com/chenyuIT/framework/contracts/ocr"
)

type Application struct {
}

func NewApplication() ocr.Ocr {
	return NewOcr()
}
