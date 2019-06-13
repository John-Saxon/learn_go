package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1.23
	fmt.Println("type:", reflect.TypeOf(num))
	fmt.Println("value:", reflect.ValueOf(num))

	fmt.Println("type of type:", reflect.TypeOf(reflect.TypeOf(num)))
	fmt.Println("type of value:", reflect.TypeOf(reflect.ValueOf(num)))

	num = reflect.ValueOf(num)
}
