package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "12345"
	i := strings.Index(s, "5")
	fmt.Println(s[:i])
}
