package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	in := bytes.NewBuffer(nil)
	var out bytes.Buffer
	sql := `psql -h 40.73.33.130 -U oushu -d postgres -c "select * from creditcard_fraud_train limit 1"`

	cmd := exec.Command("/bin/bash")
	inPipe, err := cmd.StdinPipe()
	stdout, err := cmd.StdoutPipe()
	in.WriteString(sql)
	cmd.Stdout = &out

	fmt.Println("I am here4")
	err = cmd.Start()
	reader := bufio.NewReader(stdout)

	fmt.Println("I am here2")
	go func() {
		// in.WriteString("test\n")
		fmt.Println("I am here1")
		inPipe.Write([]byte("test"))
		fmt.Println("I am here")
		inPipe.Write([]byte("\n"))
		fmt.Println("I am here5")

	}()
	go func() {
		for {
			line, err2 := reader.ReadString('a')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(line)
		}
	}()
	fmt.Println("I am here3")
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	fmt.Println("Out :", out.String())
}

// func main() {
// 	in := bytes.NewBuffer(nil)
// 	var out bytes.Buffer
// 	sql := `psql -h 40.73.33.130 -U oushu -d postgres -c "select * from creditcard_fraud_train limit 1"`

// 	cmd := exec.Command("/bin/bash")
// 	cmd.Stdin = in
// 	in.WriteString(sql)
// 	cmd.Stdout = &out
// 	err := cmd.Start()
// 	go func() {
// 		in.WriteString("test")
// 	}()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = cmd.Wait()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Out :", out.String())
// }
