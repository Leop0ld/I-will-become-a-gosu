package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := cmplx.Pow(x, 3)
	for z != Newtons(z, x) {
		z = Newtons(z, x)
	}
	return z
}

func Newtons(z, x complex128) complex128 {
	return z - (z*z*z-x)/(3*z*z)
}

func main() {
	fmt.Println(Cbrt(2))
}
