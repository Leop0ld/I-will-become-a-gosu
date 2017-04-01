package main

/*
모든 Go 프로그램은 패키지로 구성되어 있습니다.

프로그램은 main 패키지에서부터 실행을 시작합니다.

이 프로그램은 `"fmt"`와 "math" 패키지를 import 해서 사용하고 있습니다.

패키지 이름은 디렉토리 경로의 마지막 이름을 사용하는 것이 규칙입니다.

예를 들어 "path/filepath" 를 사용한다면 패키지명은 filepath 입니다.
 */

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Happy ", math.Pi, "Day!")
}