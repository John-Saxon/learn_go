package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("UnExist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
