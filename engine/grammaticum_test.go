package engine_test

import (
	"testing"
	"github.com/herval/adventuretime/engine"
	"fmt"
	"time"
	"math/rand"
	"github.com/herval/adventuretime/util"
)

func TestNames(t *testing.T) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	util.DebugFmt("Seed:", seed)

	for i := 0; i < 50; i++ {
		name := engine.RandomName()
		if name == "" {
			t.Fail()
		}
		fmt.Println(engine.RandomName())
	}
}
