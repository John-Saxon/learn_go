package main

import "fmt"

func main() {
	ids := []int{1, 2, 3}
	fmt.Println(ids[:len(ids)])
	// fmt.Println(ids[:-1])
	// fmt.Println(ids[-1])
}
