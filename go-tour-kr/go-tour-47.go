package main

import (
	"fmt"
	"time"
)

/*
switch case 문에서 조건을 생략하면 true 와 같은 뜻임
긴 if else 문을 쓰는 대신 쓸 수 있음
*/

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("좋은 아침!")
	case t.Hour() < 17:
		fmt.Println("좋은 오후!")
	default:
		fmt.Println("좋은 저녁!")
	}
}
