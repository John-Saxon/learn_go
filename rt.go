package main 

import (
	"fmt"
	"time"
)

func say(s string){
	for i:= range s{
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c ",s[i])
	}
}

func main(){
	go say("world")
	say("hello")
}
