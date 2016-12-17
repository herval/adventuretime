package engine

import (
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
	util.DebugFmt("%+v", op)
	state, executed := op.Execute(self.State)

	if !executed.Noop {
		state.CurrentTurn += 1
	}
	// TODO update the rest of the world

	self.State = state
	return self.State, executed
}
