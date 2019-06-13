package main

import (
	"fmt"
	"regexp"
)

func idCheck(id string) bool {
	if len(id) == 18 {
		return true
	} else if len(id) == 11 {
		match, b := regexp.MatchString(`\d{11}`, id)
		fmt.Println(match)
		fmt.Println(b)
		return true
	} else {
		return false
	}
}

func main() {
	idCheck("11111111111")
	idCheck("1111111 111")
}
