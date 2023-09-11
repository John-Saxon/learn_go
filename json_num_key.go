package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	type ParamOptionReq struct {
		ParamId  int64            `json:"param_id"`
		ParamMap map[int64]string `json:"param_map"`
	}
	// a := make(map[int64]string)
	v := `{"1":"a"}`
	v = `{
		"param_id": 719,
		"param_map": {
		  "719": "无劳动能力','部分丧失劳动能力','有劳动能力','完全丧失劳动能力','"
		}
	  }`
	a := ParamOptionReq{}
	// v = `{
	// 		"719": "无劳动能力','部分丧失劳动能力','有劳动能力','完全丧失劳动能力','"
	// 	  }`
	err := json.Unmarshal([]byte(v), &a)
	fmt.Println(err)
	fmt.Println(a)
}