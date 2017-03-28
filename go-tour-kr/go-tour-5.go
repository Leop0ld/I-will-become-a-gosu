package main

/*
이 코드에서는 여러개의 "package" 를 소괄호로 감싸서 import를 표현합니다. 아래와 같이 import 문장을 여러번 사용할 수 도 있습니다.
 */

import (
	"fmt"
	"math"
)
// import "fmt"
// import "main"


func main() {
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}