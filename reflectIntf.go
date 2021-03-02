package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type UnicomParamValue interface {
	GetParamString(pName string) string
	UnmarshalJSON(data []byte) error
}

type StringParam string

func (s StringParam) GetParamString(pName string) string {
	return fmt.Sprintf("%s=%s", pName, s)
}

func (s StringParam) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &s)
	return err
}

type StringSliceParam []string

func (s StringSliceParam) GetParamString(pName string) string {
	ss := []string{}
	for _, v := range s {
		ss = append(ss, fmt.Sprintf("%s=%s", pName, v))
	}
	return strings.Join(ss, "&")
}

func (s StringSliceParam) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &s)
	return err
}

func main() {
	t := make(map[string]interface{})
	t["a"] = "a"
	t["b"] = []string{"b", "c"}
	for _, v := range t {
		fmt.Println("type:", reflect.TypeOf(v))
		// fmt.Println("value:", reflect.ValueOf(v))
		fmt.Println("type:", reflect.TypeOf(v).String())
		// fmt.Println("value:", reflect.ValueOf(v).String())
		// fmt.Println("kind:", v.Kind())
	}
}
