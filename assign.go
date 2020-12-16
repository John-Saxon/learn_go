package main

import(
	"fmt"
)

type t struct {
	A string
	B int
}

func main(){
	var t t 
	t.B, err := getValue()
	fmt.Println(t)
}

func getValue() (int, error) {
	return 0, nil
}
