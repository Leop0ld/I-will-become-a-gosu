package main

import (
	"fmt"
)

/*
_ 를 이용해서 인덱스(index)나 값(value)를 무시 가능
*/

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
