package main

import (
	"io"
	"log"
	"os"
)

func copyData(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := copyData(os.Stdout, file); err != nil {
		log.Fatal(err)
	}

	file.Seek(0, 0) // ファイルポインタを先頭に戻す

	file2, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	if err := copyData(file2, file); err != nil {
		log.Fatal(err)
	}
}
