package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dz := make([][]uint8, dy)

	for i := range dz {
		dz[i] = make([]uint8, dx)
	}

	return dz
}

func main() {
	pic.Show(Pic)
}
