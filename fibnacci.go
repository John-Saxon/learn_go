package main

import "fmt"

func fibnacci(n int) int{
	if n<2{
		return n
	}
	return fibnacci(n-2) + fibnacci(n-1)
}

func main(){
	for i:=0;i<10;i++{
		fmt.Printf("%d\t", fibnacci(i))
}
}
