package main 

import "fmt"
	
func main(){
	a:= 100
	b:=200
	var ret int 

	ret = max(a,b)

	fmt.Println(ret)
}

func max(num1, num2 int) int{
	var result int 
	if (num1 > num2){
		result = num1
	}else{
		result = num2
	}
	return result
}

