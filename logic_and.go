package main 

import "fmt"

func main(){
	
	a := 0
	b := 1
	var flag bool

	flag = a&b
	fmt.Println(flag)
	flag = a|b
	fmt.Println(flag)

	return
}
