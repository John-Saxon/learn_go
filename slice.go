package main 

import "fmt"

func main(){
	var numbers = make([]int,3,5)
	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("%d, %d, %v\n",len(x), cap(x),x)
}
