package main

import (
	"fmt"
	"time"
)

/*
select 의 default 케이스는 현재 수행 준비가 완료된 케이스가 없을 때 수행됩니다.
블로킹 없이(비동기적인) 송/수신을 하고자 할 때 default 케이스를 사용하세요.
*/

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)

	for {
		select {
		case <- tick:
			fmt.Println("Tick.")
		case <- boom:
			fmt.Println("Boom!")
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}