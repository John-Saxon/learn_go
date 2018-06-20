package main 

import (
	"log"
	// "bufio"
	"fmt"
	// "os"
	"os/exec"
	"bytes"
)

func main(){
	// inputReader := bufio.NewReader(os.Stdin)
	// input, err := inputReader.ReadString('\n')
	// log.Println(input, err)

	cmd := exec.Command("source /Users/johnsaxon/.bash_profile;psql -h 42.159.87.142 -p 5432 -d postgres -U oushu;echo 'lavaadmin' > 0")
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%s", out.String())
	log.Println(out)
	
}