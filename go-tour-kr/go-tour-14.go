package main

import "fmt"

/*
상수는 const 키워드와 함께 변수처럼 선언
상수는 문자, 문자열, 불린 값 중 하나로 가능
 */

const Pi = 3.14

func main() {
	const World = "얀녕!"
	fmt.Println("Hello, ", World)
	fmt.Println("Hello ", Pi, "Day!")

	const Truth = true
	fmt.Println("Go rules? ", Truth)
}
