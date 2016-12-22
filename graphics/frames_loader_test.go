package graphics_test

import (
	"fmt"
	"testing"
	"github.com/herval/adventuretime/graphics"
)

func TestSitemapLoader(t *testing.T) {
	loader := graphics.FramesLoader{}

	data := loader.Parse("../resources/sprites.json")

	if data == nil {
		t.Failed()
	}

	fmt.Printf("Loaded succesfully: %+v sprites\n", len(data)) //reflect.ValueOf(data).MapKeys())
}
