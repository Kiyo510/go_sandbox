package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type record struct {
	Number  int    `csv:"number"`
	Message string `csv:"message"`
}

func main() {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for i := 0; i < 1000*1000; i++ {
			c <- record{
				Number:  i + 1,
				Message: "Hello",
			}
		}
		return
	}()

	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gocsv.MarshalChan(c, gocsv.DefaultCSVWriter(f)); err != nil {
		log.Fatal(err)
	}
}
