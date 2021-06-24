package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		res_el := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			res_el[x] = uint8((x + y) / 2)
			// res_el[x] = uint8(x ^ y)
			// res_el[x] = uint8(x * y)
		}
		res[y] = res_el
	}
	return res
}

func main() {
	pic.Show(Pic)
}
