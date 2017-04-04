package main

import (
	"fmt"
	"time"
)

/*
조건이 참인 case를 찾으면 switch 문을 마침
*/

func main() {
	fmt.Println("토요일은 언제입니까? ")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("오늘이네!")
	case today + 1:
		fmt.Println("내일이네.")
	case today + 2:
		fmt.Println("2일 뒤야.")
	default:
		fmt.Println("아직 멀었네.")
	}
}
