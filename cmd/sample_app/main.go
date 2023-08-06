package main

import (
	"fmt"
	"os"
)

type Warning interface {
	Show(message string)
}

type ConsoleWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
}

type DeskTopWarning struct{}

//func (d DeskTopWarning) Show(message string) {
//	beeep.Alert(os.Args[0], message, "")
//}

func main() {
	var warn Warning
	warn = &ConsoleWarning{}
	warn.Show("console warning")

	warn = &DeskTopWarning{}
	//var _ Warning = &DeskTopWarning{}
	warn.Show("desktop message")
}
