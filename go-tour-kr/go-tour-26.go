package main

import "fmt"

/*
구조체 안의 데이터에는 .(dot) 으로 접근 가능
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
