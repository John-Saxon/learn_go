package main

import (
"fmt"
"errors"
)

func main(){
	var ret error = nil
	fmt.Println(ret)
	if ret == nil{
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}
}
