package debug

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func PrintError(i interface{}, err error) {
	methodName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fmt.Printf("%sThere was an error in %s\nERROR: {%s}%s\n", red, methodName, err, reset)
}

func PrintInfo(i interface{}, strs ...string) {
	methodName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	message := strings.Join(strs, ", ")
	fmt.Printf("%s%s INFO: {%s}%s\n", yellow, methodName, message, reset)
}

func PrintSucc(i interface{}, strs ...string) {
	methodName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	message := strings.Join(strs, " ")
	fmt.Printf("%s%s SUCC: {%s}%s\n", green, methodName, message, reset)
}