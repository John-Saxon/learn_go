package main 

import "fmt"

func main(){
	a:=[...]int{1,2,3,4,5}

	var p *[5]int = &a
	
	fmt.Println(*p, a)

	for index,value := range *p{
		fmt.Println(index, value)
	}

	var p2 [5]*int

	i, j := 10,20 

	p2[0] = &i
	p2[1] = &j 
	
	for index,value := range p2{
		if value != nil{
			fmt.Println(index, *value)
		}else{
			fmt.Println(index, value)
		}

	}
}
