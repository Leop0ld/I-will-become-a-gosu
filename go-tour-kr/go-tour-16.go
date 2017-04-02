package main

import "fmt"

/*
Go 에는 반복문이 for 문 밖에 없다
C 랑 비슷한데 소괄호가 필요없다는 게 다른 점
근데 중괄호는 필요함
 */

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum: ", sum)
}