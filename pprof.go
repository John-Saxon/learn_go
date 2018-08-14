package main

import (
_ "net/http/pprof"
"net/http"
"time"
"fmt"
)

func main(){

	go func(){
	http.ListenAndServe("0.0.0.0:8888",nil)
	}()

	for{
		fmt.Println("1")
		time.Sleep(time.Duration(2)*time.Second)
	}

	return
}	
