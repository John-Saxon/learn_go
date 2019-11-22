package main

import (
	"fmt"
)

type ttt struct {
	A int
	B []string
}

func main() {
	var a []string
	fmt.Println(append(a, "aaa"))

	var b ttt
	b.B = append(b.B, "bbb")
	fmt.Println(b)
}
