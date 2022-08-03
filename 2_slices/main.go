package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		picture[i] = make([]uint8, dx)
	}
	for y, row := range picture {
		for x := range row {
			row[x] = uint8(x ^ y)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
