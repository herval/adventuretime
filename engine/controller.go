package engine

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/herval/adventuretime/util"
)

type Controller struct {
	State *GameState
}

func NewController() *Controller {
	rand.Seed(time.Now().Unix())

	return &Controller{
		State: NewGame(),
	}
}

func (self *Controller) Execute(op Command) (*GameState, Result) {
	util.Debug(fmt.Sprintf("%+v", op))
	state, executed := op.Execute(self.State)
	self.State = state
	return self.State, executed
}
