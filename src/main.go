package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//println(filepath.Base("/src/file.txt"))
	//println(filepath.Dir("/src/file.txt"))
	//println(filepath.Clean("/src/../file.txt"))
	//println(filepath.Ext("/src/file.txt"))
	//println(filepath.IsAbs("/src/file.txt"))
	//println(filepath.IsAbs("./src/file.txt"))
	//println(filepath.Join("/src", "/file.txt"))
	//absolute, err := filepath.Abs("../file.txt")
	//if err == nil {
	//	println(absolute)
	//}
	//
	//absolute, err = filepath.Rel("/src", "/src/file.txt")
	//if err == nil {
	//	println(absolute)
	//}

	var files []string
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	fmt.Println(files)
}
