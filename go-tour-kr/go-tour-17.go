package main

import "fmt"

/*
다른 언어의 While 문 처럼 조건문만 작성 가능
 */

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}

	/*
	조건문을 생략하면 무한 루프
	for {

	}
	*/

	fmt.Println(sum)
}
