package main

import (
	"fmt"

	"github.com/oushu/gocommon/other"
	"github.com/shirou/gopsutil/process"
)

func main() {
	pss, err := process.Processes()
	// pss, err := process.Pids()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(other.EncodeJson(pss))
	i := 0
	for _, ps := range pss {
		// if i > 10 {
		// 	break
		// }
		if ps.Pid == 5984 || ps.Pid == 5979 {
			fmt.Println(ps.Name())
			fmt.Println(ps.Ppid())
			i++
		}
	}
}

func getProcByName(name string, cache []*process.Process) (*process.Process, error) {

}

func getProcByGrandPid(pid int32, cache []*process.Process)
