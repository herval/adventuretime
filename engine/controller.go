package engine

import (
	"math/rand"
	"time"
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
	state, executed := op.Execute(self.State)
	self.State = state
	return self.State, executed
}
