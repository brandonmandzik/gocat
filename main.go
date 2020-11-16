// +build darwin

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("gocat expectes at least one more file path")
		os.Exit(1)
	}
	for _, filePath := range os.Args[1:] {
		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("Could not open: %s \n %s", filePath, err)
			file.Close()
			continue
		}
		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			log.Printf("Could not print file ... \n%s", err)
		}
		file.Close()
	}
}
