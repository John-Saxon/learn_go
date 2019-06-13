package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 12.3

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPoiter := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPoiter)
	fmt.Println(*convertPoiter)
	fmt.Println(convertValue)
}
