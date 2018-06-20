package main 

import "fmt"

func main(){
	
	ch := make(chan int,5)

	a := [...]int{1,2,3,4,5}
	for i:= range a{
		go func(){
			ch <- i
		}()
	}
	sum := 0
	for{
		fmt.Println(<-ch)
		sum++
		if sum>4{
			break
		}
	}
	
	return
}
