package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigHandler := SignalSetNew()
	sigHandler.Register(syscall.SIGINT, sigHandlerFunc)
	sigHandler.Register(syscall.SIGUSR1, sigHandlerFunc)
	sigHandler.Register(syscall.SIGUSR2, sigHandlerFunc)
	sigHandler.Register(syscall.SIGQUIT, sigHandlerFunc)
	sigHandler.Register(syscall.SIGKILL, sigHandlerFunc)

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan)

	for true {
		select {
		case sig := <-sigChan:
			err := sigHandler.Handle(sig, nil)
			if err != nil {
				fmt.Println("[ERROR] unknown signal received:", sig)
				os.Exit(1)
			}
		default:
			time.Sleep(time.Duration(3) * time.Second)
		}
	}

}

func sigHandlerFunc(s os.Signal, arg interface{}) {
	switch s {
	case syscall.SIGINT:
		fmt.Println("Receive sigint, stopping...")
	default:
		fmt.Println("Receive ", s, " stopping...")
	}
}

type SignalHandler func(s os.Signal, arg interface{})

type SignalSet struct {
	m map[os.Signal]SignalHandler
}

func SignalSetNew() *SignalSet {
	ss := new(SignalSet)
	ss.m = make(map[os.Signal]SignalHandler)
	return ss
}

func (set *SignalSet) Register(s os.Signal, handler SignalHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *SignalSet) Handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}
	panic("won't reach here")
}
