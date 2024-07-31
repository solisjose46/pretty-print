package main

import (
	"github.com/solisjose46/pretty-print/debug"
	"errors"
)

func main() {
	debug.PrintInfo(main, "Print testing", "newline confirmed")
	err := errors.New("This is test err")
	debug.PrintError(main, err)
	debug.PrintSucc(main, "everything goo", "everything gaga")
}