package main 

import "fmt"

func main(){
	i:=15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64)(result uint64){
	if(n>0){
		ret := n*Factorial(n-1)
		return ret
	}
	return 1
}
