package main

import "fmt"

func main(){
	var x interface{}

	switch i:=x.(type){
		case nil:
			fmt.Printf(i)
		case int:
			fmt.Printf("int")
		case bool,string:
			fmt.Printf(i)
		default:
			fmt.Printf("未知类型")
	}
}
