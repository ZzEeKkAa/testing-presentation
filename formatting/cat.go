package formatting

import (
	"image/color"
)

type Cat struct {
	color color.RGBA
	name  string
}

func NewCat(color color.RGBA, name string) *Cat {
	return &Cat{
		color: color,
		name:  name,
	}
}
