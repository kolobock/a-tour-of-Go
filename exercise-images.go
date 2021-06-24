package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		res_el := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// res_el[x] = uint8((x + y) / 2)
			// res_el[x] = uint8(x ^ y)
			res_el[x] = uint8(x * y)
		}
		res[y] = res_el
	}
	return res
}

type Image struct {
	w   int
	h   int
	pic [][]uint8
}

func (im *Image) init_defaults() {
	im.w = 100
	im.h = 100
	im.pic = Pic(im.w, im.h)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{im.pic[x][0], im.pic[x][y], 255, 255}
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func main() {
	m := Image{}
	m.init_defaults()
	pic.ShowImage(m)
}
