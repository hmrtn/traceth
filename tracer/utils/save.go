package utils

import (
	"fmt"
	"log"
	"os"
)

func Save(filename string, trace []byte) {
	f, err := os.Create(filename)
	defer f.Close()
	ok, err := f.Write(trace)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nSaved %d bytes to %s", ok, filename)
}
