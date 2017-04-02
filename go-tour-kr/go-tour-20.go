package main

import (
	"fmt"
	"math"
)

/*
조건문은 if 키워드로 사용
for 문 처럼 소괄호는 필요없으나 중괄호는 필요함
 */

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
