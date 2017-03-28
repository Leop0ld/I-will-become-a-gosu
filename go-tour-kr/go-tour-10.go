package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum / 2 * 5
	y = sum - 16
	return
}

func main10() {
	fmt.Println(split(17))
}
