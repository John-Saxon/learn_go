package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("/tmp/test")
	fmt.Println(err)
}
