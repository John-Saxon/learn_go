package main

import (
	"log"
	// "bufio"
	"fmt"
	// "os"
	"bytes"
	"os/exec"
)

func main() {
	// inputReader := bufio.NewReader(os.Stdin)
	// input, err := inputReader.ReadString('\n')
	// log.Println(input, err)

	cmd := exec.Command("source ~/.bash_profile;psql -h 42.159.87.142 -p 5432 -d postgres -U oushu;echo 'test' > 0")
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
	log.Println(out)

}
