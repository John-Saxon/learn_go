package main

import(
"os"
"fmt"
)

func main(){
	f,err := os.Create("/tmp/test.txt")
    defer f.Close()
    if err !=nil {
        fmt.Println(err.Error())
    } else {
        _,err=f.Write([]byte("要写入的文本内容"))
	fmt.Println(err)
    }
}
