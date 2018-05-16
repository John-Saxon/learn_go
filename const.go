package main 

import "fmt"

func main(){
	const LENGTH = 7
	const WIDTH int = 5
	var area int
	const a,b,c = 1, false, "str"
//	LENGTH = 8
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
}
