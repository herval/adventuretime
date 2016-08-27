package util

var DEBUGGING = true

func Debug(msg string) {
	if DEBUGGING {
		println(msg)
	}
}
