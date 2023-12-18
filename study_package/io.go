package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("study_package/text.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes:\n", count)
	fmt.Println(string(data[:count]))
}
