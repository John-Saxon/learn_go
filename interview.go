package main

import "fmt"
import "strings"
import "strconv"

type person struct {
	name string
	age  int
	sex  string
}

type Group struct {
	Id     int
	Name   string
	Member map[int]person
}

func (g *Group) AddPerson(id int, p *person) {
	g.Member[id] = *p
}

func NewGroup(nameAndId string) *Group {
	// splits := strings.Split(nameAndId, '_')
	splits := strings.Split(nameAndId, "_")
	// Id := string(splits[1])
	Id, _ := strconv.Atoi(splits[1])
	return &Group{
		Id:     Id,
		Name:   splits[0],
		Member: make(map[int]person),
	}
}

func main() {
	//初始化组织
	var nameAndId = "oushu_1"
	group := NewGroup(nameAndId)

	//初始化成员信息
	p1 := person{}
	p1.name = "p1"
	p1.age = 25

	p2 := &person{}
	p2.name = "p2"
	p2.sex = "male"

	//添加成员
	group.AddPerson(1, &p1)
	go group.AddPerson(2, p2)

	fmt.Println(group)

	// 信息补录
	// group.Member[1].sex = "female"
	// group.Member[2].age = 30

	fmt.Println(group)
}
