package main

// This story told us to be careful when using 'make'
import (
	"fmt"
)

func main() {
	urls := make(map[string]string, 3)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	fmt.Println(urls)
	fmt.Println(len(urls))
	names := make([]string, len(urls))
	for key, _ := range urls {
		fmt.Println(key, len(key))
		names = append(names, key)
	}
	fmt.Println(names, len(names))
	for i, v := range names {
		fmt.Println(i, v)
	}

	var names1 []string
	for key, _ := range urls {
		fmt.Println(key, len(key))
		names1 = append(names1, key)
	}
	fmt.Println(names1, len(names1))

	names2 := []string{}
	for key, _ := range urls {
		fmt.Println(key, len(key))
		names2 = append(names2, key)
	}
	fmt.Println(names2, len(names2))

	return
}
