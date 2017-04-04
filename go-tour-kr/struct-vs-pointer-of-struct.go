package main

import "fmt"

type TestStruct struct {
	name string
}

func (t TestStruct) test_func(name string) {
	fmt.Println("t Change name -> ", name)
	t.name = name
}

func (t *TestStruct) test_func_p(name string) {
	fmt.Println("tp Change name -> ", name)
	t.name = name
}

func main() {
	t := TestStruct{"myungseo"}
	t.test_func("Myungseo")
	fmt.Println(t.name)

	tp := &TestStruct{"kang"}
	tp.test_func_p("Kang")
	fmt.Println(tp.name)
}
