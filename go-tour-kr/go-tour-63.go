package main

import (
	"time"
	"fmt"
)

/* 드디어 고루틴 
고루틴은 Go 런타임에 의해 관리되는 경량 쓰레드

go f(x, y, z)
위의 코드는 새로운 고루틴을 시작시킵니다.

f(x, y, z)
현재의 고루틴에서 f , x , y , z 가 evaluation되고, 새로운 고루틴에서 f 가 수행됨

고루틴은 동일한 주소 공간에서 실행되므로, 공유되는 자원으로의 접근은 반드시 동기화 되어야 함
[[http://golang.org/pkg/sync/](sync)] 패키지가 이에 대한 유용한 기본 기능을 제공함
Go 에서는 그외에도 다양한 기본 기능을 제공하니 크게 필요하지 않을 수도 ㅇㅇ
*/

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
	go say("world")
    say("hello")
}