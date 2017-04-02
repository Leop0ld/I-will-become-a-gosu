package main

import (
	"fmt"
)

/*
맵은 값에 키를 지정함

맵은 반드시 사용하기 전에 make 를 명시해야 함. new 가 아님

make 를 수행하지 않은 nil 에는 값을 할당할 수 없음
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["OPEN GALLERY"] = Vertex{
		37.5122602, 127.0210934,
	}
	fmt.Println(m["OPEN GALLERY"])
}
