package main

/*
패키지를 import 하면 그 패키지가 외부로 export 한 것(메서드나 변수, 상수 등)들에 대해 접근할 수 있습니다.

Go에서는 첫 문자가 대문자로 시작하면 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됩니다.

예를 들어 Foo 와 FOO 는 외부에서 참조할 수 있지만 foo 는 참조 할 수 없습니다.
 */

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	// ERROR: fmt.Println(math.pi)
}
