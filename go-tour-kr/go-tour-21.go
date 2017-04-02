package main

import (
	"fmt"
	"math"
)

/*
if 문을 좀 더 짧게 표현할 수 있음
비교를 실행하기 전에 짧은 표현식 수행 가능
짧은 실행문을 통해 선언된 변수는 if 안쪽 범위 에서만 사용가능함
*/

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
