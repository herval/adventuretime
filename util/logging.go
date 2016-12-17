package util

import "fmt"

var DEBUGGING = false

func Debug(msg string) {
	if DEBUGGING {
		println(msg)
	}
}

func DebugFmt(msg string, params ...interface{}) {
	if DEBUGGING {
		println(fmt.Sprintf(msg, params))
	}
}