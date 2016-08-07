package main

// a command - eg walk, fight, talk, etc. Commands alter the game state and return a result
type Command interface {
	Execute(state *GameState) (*GameState, Result)
}

// the result of a command
type Result interface {
	Describe() string
}

type Ok struct {
	description string
}

func (self *Ok) Describe() string {
	return self.description
}

type Noop struct {
	description string
}

func (self *Noop) Describe() string {
	return self.description
}

// Eat/consume/drink
// Walk/move/go
// Pickup/take
// Drop

type UnknownCommand struct{}

func (self *UnknownCommand) Execute(state *GameState) (*GameState, Result) {
	return state, &Noop{
		description: "Unknown command",
	}
}

type Move struct {
	direction Direction
}

func (self *Move) Execute(state *GameState) (returnedState *GameState, result Result) {
	returnedState = state

	if self.direction == UNKNOWN {
		result = &Noop{
			description: "Can't walk that way.",
		}
		return
	}

	player := state.player
	currentRoom := state.player.currentLocation

	for _, door := range currentRoom.doors {
		if self.direction == door.facing {
			if door.locked {
				result = &Noop{
					description: "The door is locked.",
				}
			} else {
				player.currentLocation = door.Open()
				result = &Ok{
					description: "You walked " + DirectionToStr(self.direction) + ".",
				}
			}
			return
		}
	}

	result = &Noop{
		description: "Can't walk that way.",
	}
	return
}
