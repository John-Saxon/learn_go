package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	in := bytes.NewBuffer(nil)
	var out bytes.Buffer
	sql := `psql -h 40.73.33.130 -U oushu -d postgres -c "select * from creditcard_fraud_train limit 1"`

	cmd := exec.Command("/bin/bash")
	inPipe, err := cmd.StdinPipe()
	in.WriteString(sql)
	cmd.Stdout = &out

	fmt.Println("I am here4")
	err = cmd.Run()

	fmt.Println("I am here2")
	go func() {
		in.WriteString("lavaadmin")
		fmt.Println("I am here1")
		inPipe.Write([]byte("lavaadmin"))
		fmt.Println("I am here")
		inPipe.Write([]byte("\n"))

	}()
	fmt.Println("I am here3")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Out :", out.String())
}
