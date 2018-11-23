package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("createAndWriteFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := f.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
