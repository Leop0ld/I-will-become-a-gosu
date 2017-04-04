package main

import (
	"fmt"
	"math"
)

/*
메소드는 아무 타입에나 붙일 수 있음
다만 다른 패키지에 있는 타입이나 기본 타입들에 메소드를 붙이는 것은 불가능함
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
