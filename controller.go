package main

import (
	"math/rand"
	"time"
)

type Controller struct {
	state  *GameState
	parser CommandParser
}

func NewController() *Controller {
	rand.Seed(time.Now().Unix())

	return &Controller{
		state:  NewGame(),
		parser: &StandardParser{},
	}
}

func (self *Controller) Execute(cmd string) (*GameState, Result) {
	op := self.parser.ParseCommand(cmd)
	state, executed := op.Execute(self.state)
	self.state = state
	return self.state, executed
}
