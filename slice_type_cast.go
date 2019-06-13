package main

import "fmt"

func main() {
	ints := []interface{}{1, 2, 3}
	fs := float32(ints[0])
	fmt.Println(fs)
}
