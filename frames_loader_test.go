package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestSitemapLoader(t *testing.T) {
	loader := FramesLoader{}

	data := loader.Parse("./sprites.json")

	if data == nil {
		t.Failed()
	}

	fmt.Printf("Loaded succesfully: ")
	fmt.Printf("%+v", reflect.ValueOf(data).MapKeys())
}
