package main 

import "fmt"

func main(){
	numbers1 := make([]int, 0, 5)

	printSlice(numbers1)

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
