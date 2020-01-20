package main

import (
	"encoding/json"
	"fmt"
)

type MindMap struct {
	Name     string    `json:"name"`
	Value    int       `json:"value"`
	Children []MindMap `json:"children,omitempty"`
}

type MindMap struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime string    `json:"create_time"`
	TaskId     int       `json:"task_id"`
	Detail     string    `json:"Detail"`
	Children   []MindMap `json:"children,omitempty"`
}

func main() {
	mmStr := `{"name":"中国","children":[{"name":"浙江","children":[{"name":"杭州","value":100},{"name":"宁波","value":100},{"name":"温州","value":100},{"name":"绍兴","value":100}]},{"name":"广西","children":[{"name":"桂林","children":[{"name":"秀峰区","value":100},{"name":"叠彩区","value":100}]},{"name":"南宁","value":100},{"name":"柳州","value":100},{"name":"防城港","value":100}]},{"name":"黑龙江","children":[{"name":"哈尔滨","value":100},{"name":"齐齐哈尔","value":100},{"name":"牡丹江","value":100},{"name":"大庆","value":100}]},{"name":"新疆","children":[{"name":"乌鲁木齐"},{"name":"克拉玛依"},{"name":"吐鲁番"},{"name":"哈密"}]}]}`
	mm := MindMap{}
	err := json.Unmarshal([]byte(mmStr), &mm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mm)

	newStr, err := json.Marshal(mm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(newStr))
}
