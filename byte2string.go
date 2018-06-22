package main 

import "fmt"

func main(){

	var str string = "This is a string."

	var data []byte = []byte(str)

	fmt.Println("str:", str)

	fmt.Println("[]byte:", data)

	fmt.Println("byte to string:", string(data))

	return 
}
