package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	retVal := make(map[string]int, 0)

	for _, value := range ss {
		retVal[value] += 1
	}

	return retVal
}

func main() {
	wc.Test(WordCount)
}
