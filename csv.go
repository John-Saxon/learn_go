package main  
  
import (  
    "encoding/csv"  
    "fmt"  
    "os"  
)  
  
func main() {  
    file, _ := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)  
    w := csv.NewWriter(file)  
    w.Write([]string{"123", "234234", "345345", "234234"})  
    w.Flush()  
  
    file.Close()  
  
    rfile, _ := os.Open("test.csv")  
    r := csv.NewReader(rfile)  
    strs, _ := r.Read()  
    for _, str := range strs {  
        fmt.Print(str, "\t")  
    }  
}
