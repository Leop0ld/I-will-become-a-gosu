package main

import (
	"fmt"
	"net/http"
)

/*
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

Hello 라는 타입은 http.Handler 를 구현함
*/

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
