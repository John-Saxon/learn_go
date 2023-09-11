package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	cwd = "~/test/github.com/learn_go/file"
	path := path.Join(cwd, "new_dir", "createFile_p.txt")
	log.Println(path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
