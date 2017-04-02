package main

import "fmt"

/*
struct(구조체) 는 데이터들의 집합
type 키워드로 이름을 지정해줄 수 있음
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
