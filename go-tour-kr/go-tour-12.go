package main

/*
변수 선언과 동시에 초기화는 다음과 같이 가능
초기화를 하는 경우 타입 생략가능. 초기화 할 값에 따라 타입 결정
 */

import "fmt"

var x1, y1, z1 int = 1, 2, 3

var c1, python1, java1 = true, false, "no!"

func main() {
	fmt.Println(x1, y1, z1, c1, python1, java1)
}
