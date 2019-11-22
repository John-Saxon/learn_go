package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(3 * time.Second)
	stop := time.Now()
	fmt.Println(stop.Sub(start).Seconds())
}
