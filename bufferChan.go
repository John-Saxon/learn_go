package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	// i := 1
	/*1.
	// fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan receive]:
	// main.main()
	//         /Users/johnsaxon/test/github.com/learn_go/bufferChan.go:15 +0x19c
	// exit status 2
	*/

	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	val, ok := <-ch
	// 	if ok == false {
	// 		fmt.Println("chan is closed")
	// 		break
	// 	}
	// 	fmt.Println(val)
	// }

	select {
	// fmt.Println(i)
	// i++
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println(-1)
	}

	close(ch)
}
