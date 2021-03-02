package main

import (
	"fmt"

	"github.com/oushu/gocommon/oscommand"
)

func main() {
	fmt.Println(CreateTaskDir("~/justatest/ha"))
}

func CreateTaskDir(dir string) (res bool, report string) {
	command := "mkdir -p " + dir
	_, stderr, err := oscommand.Run(command)
	if err != nil {
		return false, "[FAILED]can not create dir: " + stderr
	}
	return true, "[OK] dir:" + dir
}
