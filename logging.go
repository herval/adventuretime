package main

var DEBUGGING = false

func Debug(msg string) {
	if DEBUGGING {
		println(msg)
	}
}
