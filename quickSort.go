package main

import "fmt"

func main(){
	// a := []int{1, 2, 3, 4, 5}
	// a := []int{5, 4, 3, 2, 1}
	a := []int{4, 5, 1, 3, 2}
	QuickSort_V2(a)
	fmt.Println("a:\n", a)
}

func QuickSort_V2(values []int){
    if len(values)<=1{
        return
    }
    mid, i := values[0], 1
    head, tail := 0, len(values)-1

    for head<tail{
        fmt.Println(values)
        if values[i]>mid{
            values[i], values[tail] = values[tail], values[i]
            tail--
        }else{
            values[i], values[head]=values[head], values[i]
            head++
            i++
        }
    }
	fmt.Println("head: ", head)
    values[head] = mid
    QuickSort_V2(values[:head])
    QuickSort_V2(values[head+1:])
}