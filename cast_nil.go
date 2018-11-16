package main

import (
	"fmt"
)

func main() {
	// l := interface{0, 0, 0, 0, 0}
	var l interface{} = []int32{0, 0, 0, 0, 0, 0}
	fmt.Println(l)

	cl, test := l.([]int64)
	fmt.Println(cl)
	fmt.Println(test)
	return

}
