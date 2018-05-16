package main 

import "fmt"

func main(){
	var a int = 4
var b int32
var c float32
var ptr *int

	fmt.Printf("第一行a变量类型为%T\n", a)
	fmt.Printf("第二行b变量类型为%T\n", b)
	fmt.Printf("第三行c变量类型为%T\n", c)

	ptr = &a
	fmt.Printf("a的值为 %d\n", a)
	fmt.Printf("*prt 为 %d\n", *ptr)
}
