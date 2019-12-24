package main

import (
	"fmt"
)

type MyStruct struct {
	Id   int
	Name string
	Age  int
}

func main() {

	ss := []MyStruct{
		MyStruct{
			Id:   1,
			Name: "1",
			Age:  1,
		},
	}

	StructSlice(ss)

	sps := []*MyStruct{
		&MyStruct{
			Id:   1,
			Name: "1",
			Age:  1,
		},
	}

	StructPointerSlice(sps)
	fmt.Println("StructPointerSlice: ", sps)
	for _, s := range sps {
		fmt.Println(s)
	}
	PointerStructSlice(&sps)
	fmt.Println("PointerStructSlice: ", sps)
	for _, s := range sps {
		fmt.Println(s)
	}

	AppendSlice(sps)
	fmt.Println("PointerStructSlice: ", sps)
	for _, s := range sps {
		fmt.Println(s)
	}
	DeleteSlice(sps)
	fmt.Println("PointerStructSlice: ", sps)
	for _, s := range sps {
		fmt.Println(s)
	}
}

func StructSlice(ms []MyStruct) {
	for _, s := range ms {
		s.Age += 1
	}
	fmt.Println(ms)
}

func StructPointerSlice(ms []*MyStruct) {
	for _, s := range ms {
		s.Age += 1
	}
	fmt.Println(ms)
}

func PointerStructSlice(ms *[]*MyStruct) {
	for _, s := range *ms {
		s.Age += 1
	}
	fmt.Println(ms)
}

func AppendSlice(ms []*MyStruct) {
	ms = append(ms,
		&MyStruct{
			Id:   100,
			Name: "100",
			Age:  100,
		})

	fmt.Println("In AppendSlice: ", ms)
}

func DeleteSlice(ms []*MyStruct) {
	for index, s := range ms {
		if s.Age <= 100 {
			ms = append(ms[:index], ms[index+1:]...)
		}
	}
	fmt.Println("In DeleteSlice: ", ms)
}
