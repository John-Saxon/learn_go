package main

import (
	"fmt"
	"os/user"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("~/a")
	fmt.Println("path: ", path)

	user, err := user.Current()
	if nil == err {
		fmt.Println("home: ", user.HomeDir)
	}

}
