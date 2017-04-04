package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v1 := &Vertex{3, 4}
	fmt.Println(v1.Abs())
	fmt.Println(v1)

	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())
	fmt.Println(v2)
}
