package main

/*
두 개 이상의 매개변수가 같은 타입(type)일 때, 같은 타입을 취하는 마지막 매개변수에만 타입을 명시하고 나머지는 생략할 수 있습니다.
 */

import "fmt"

func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add2(42, 13))
}
