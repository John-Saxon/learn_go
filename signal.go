package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println("All args:", os.Args)
	var pid int
	pid_ptr := flag.Int("pid", -1, "pid")
	fmt.Println("Got pid ", *pid_ptr)
	flag.Parse()
	pid = *pid_ptr
	fmt.Println("Got pid ", pid)
	if pid == -1 {
		fmt.Println("Wrong pid", pid)
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Process Finding Error: %s\n", err)
		return
	}
	sig := syscall.SIGINT
	fmt.Printf("Send signal '%s' to the process (pid=%d)...\n", sig, pid)
	err = proc.Signal(sig)
	if err != nil {
		fmt.Printf("Signal Sending Error:%s\n", err)
		return
	}

}
