package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//now := time.Now()
	//fmt.Println(now.Format("2006/01/02 15:04:05"))
	//fmt.Println(now.Format("2006/1/02 03:04:05"))
	//fmt.Println(now.Format(time.RFC3339))

	//var s = "2022/12/25 07:42:38"
	//d, err := time.Parse("2006/01/02 15:04:05", s)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(d)

	d, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

	d, err = time.ParseDuration("4m")
	if err != nil {
		log.Fatal(d)
	}
	fmt.Println(d / 2)
}
