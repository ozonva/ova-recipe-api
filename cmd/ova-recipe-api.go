package main

import (
	"fmt"
	"io"
	"os"
)
func ReadFileLoop(loopCount uint64, filePath string) {
	readFileFn := func(filePath string) ([]byte, error) {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				fmt.Printf("Can not close file '%s', error: '%s'", filePath, err)
			}
		}(file)
		return io.ReadAll(file)
	}
	for idx := uint64(0); idx < loopCount; idx++ {
		if fileContent, err := readFileFn(filePath); err != nil {
			fmt.Printf("Can not read file, error: '%s'\n", err)
		} else {
			fmt.Printf("File content: '%s'\n", fileContent)
		}
	}
}

func main() {
	fmt.Println("Hi, i am ova-recipe-api!")
	ReadFileLoop(1, "./Makefile")
}
