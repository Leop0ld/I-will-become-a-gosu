package main

import "fmt"

/*
슬라이스는 make 함수로 만들 수 있습니다. 이렇게 생성된 슬라이스는 0을 할당한 배열을 생성하고, 그것을 참조함.
make 함수의 세번째 매개변수로 용량(capacity)를 제한할 수 있습니다.
*/

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
