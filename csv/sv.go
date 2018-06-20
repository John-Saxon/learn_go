package main 

import (
	"log"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"os"
)

func main(){
	file, err := os.Open("/Users/johnsaxon/tmp/test.csv")
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    reader := csv.NewReader(file)
    reader.Comment = '#' //可以设置读入文件中的注释符
	reader.Comma = ',' //默认是逗号，也可以自己设置
	k:=0
	for{
		record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
            return
		}
		log.Println("record:", record)
		if k > 0 { //record是[]strings， 怎样直接获得域值
            for _, v := range record {
				tmp := strings.Split(v, "|")
				log.Println(tmp)
            }
        }
        k = k + 1

	}


}