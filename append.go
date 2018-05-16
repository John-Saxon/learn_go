package main 

import "fmt"
func main(){
	var n []int
	n = append(n,0)
	
	fmt.Println(n)
	println(n)

	n = append(n,1)
	fmt.Println(n)

	println(n)
	n = append(n, 2,3,4)

	fmt.Println(n)

	println(n)
	n1 := make([]int, len(n), (cap(n))*2)
	fmt.Println(n1)
	println(n1)
	copy(n1,n)
	fmt.Println(n1)
	println(n1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
