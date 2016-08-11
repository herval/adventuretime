package main

import (
	"fmt"
	"testing"
)

func TestSitemapLoader(t *testing.T) {
	loader := FramesLoader{}

	data := loader.Parse("./sprites.json")

	if data == nil {
		t.Failed()
	}

	fmt.Printf("Loaded succesfully: ")
	fmt.Printf("%+v", data) //reflect.ValueOf(data).MapKeys())
}
