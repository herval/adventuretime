package graphics

import (
	"fmt"
	"testing"
)

func TestSitemapLoader(t *testing.T) {
	loader := FramesLoader{}

	data := loader.Parse("../resources/sprites.json")

	if data == nil {
		t.Failed()
	}

	fmt.Printf("Loaded succesfully: %+v sprites\n", len(data)) //reflect.ValueOf(data).MapKeys())
}
