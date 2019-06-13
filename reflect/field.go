package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Class Class
}
type Class struct {
	ClassId   int
	ClassName string
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func main() {
	user := User{1, "Allen.Wu", 25, Class{2, "class1"}}
	DoFiledAndMethod(user)

	num := 1
	DoFiledAndMethod(num)
}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fileds is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}
