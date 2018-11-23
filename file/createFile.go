package main

import (
	"log"
	"os"
)

func main() {
	// cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)ls
	// }
	// path := path.Join(cwd, "createFile_p.txt")
	// log.Println(path)
	f, err := os.Create("createAndWriteFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
