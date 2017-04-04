package main

import (
	"fmt"
	"runtime"
)

/*
switch case 문
알아서 break함
*/

func main() {
	fmt.Println("Go runs on! ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
