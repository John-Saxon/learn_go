package main

import (
	"log"
	"os"
)

func main() {
	dir := "/Users/johnsaxon/test/github.com/learn_go/file/test/test"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal(err)
	}

}
