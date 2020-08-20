package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {

	r, _ := regexp.Compile(`\$\{[\S]*?\}`)
	src := "我于${time}在${loc}购买了${good}, 然后把${good}丢了,${}"
	values := map[string]string{
		"time": "2020年2月22日",
		"loc":  "北京",
		"good": "烤鸭",
	}

	matches := r.FindAllString(src, -1)
	fmt.Println("matches: ", matches)

	for _, match := range matches {
		if len(match) < 4 {
			log.Fatal("Invalid match: ", match)
		}
		key := match[2 : len(match)-1]
		fmt.Println("key: ", key)
		value, ok := values[key]
		if !ok {
			log.Fatal("Invalid key: ", key)
		}
		src = strings.Replace(src, match, value, -1)
		// src := r.ReplaceAllString(src)
	}

	fmt.Println(src)
}
