package main

import "fmt"

func main() {
	a := []string{}
	a = nil
	fmt.Println("len a: ", len(a))
	for _, s := range a {
		fmt.Println(s)
	}
}
