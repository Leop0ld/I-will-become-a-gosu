package main

import "fmt"

/*
슬라이스는 배열의 값을 가리킵니다(point). 그리고 배열의 길이를 가지고 있습니다.
[]T 는 타입 T 를 가지는 요소의 슬라이스(slice) 입니다.
*/

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}
