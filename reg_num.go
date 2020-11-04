package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	out := "Stdout: COPY 361"
	re := regexp.MustCompile(`\d+`)
	res := re.FindString(out)
	num := int64(0)
	num, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
