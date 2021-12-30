package util

import (
	"fmt"
	"reflect"
	"runtime"
)

func GetFunctionName(f interface{}) string {
	path := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return path
}

func Logger(msg interface{}) {
	fmt.Println("------------------------------")
	var calldepth = 1
	_, f, l, ok := runtime.Caller(calldepth)
	if ok {
		fmt.Println(f, l)
	}
	fmt.Println(msg)
	fmt.Println("------------------------------")
}
