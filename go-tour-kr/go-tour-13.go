package main

import "fmt"

/*
함수 내에서 := 을 사용하면 var 생략 가능 및 타입 명시 생략 가능
그러나 함수 밖에서는 := 사용 불가
 */

func main() {
	var x2, y2, z2 int = 4, 5, 6
	c, python, java := true, false, "Yes!"

	fmt.Println(x2, y2, z2, c, python, java)
}