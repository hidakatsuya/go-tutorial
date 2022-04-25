package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sy := make([][]uint8, dy)
	for i := range sy {
		sy[i] = make([]uint8, dx)
		for j := range sy[i] {
			sy[i][j] = uint8(i + j)
		}
	}
	return sy
}

func main() {
	pic.Show(Pic)
}
