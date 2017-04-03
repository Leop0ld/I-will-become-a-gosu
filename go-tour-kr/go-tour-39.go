package main

import "fmt"

// 가장 상위의 타입이 타입명이라면 리터럴에서 타입명을 생략해도 됨

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
