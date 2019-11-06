package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		// fmt.Println(rand.Intn(10))
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Float32() * 10)
	}
}
