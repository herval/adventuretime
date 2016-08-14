package engine

// a command - eg walk, fight, talk, etc. Commands alter the game state and return a result
type Command interface {
	Execute(state *GameState) (*GameState, Result)
}

// the result of a command
type Result interface {
	Describe() string
}

type Ok struct {
	Description string
}

func (self *Ok) Describe() string {
	return self.Description
}

type Noop struct {
	Description string
}

func (self *Noop) Describe() string {
	return self.Description
}

// Eat/consume/drink
// Walk/move/go
// Pickup/take
// Drop

type UnknownCommand struct{}

func (self *UnknownCommand) Execute(state *GameState) (*GameState, Result) {
	return state, &Noop{
		Description: "Unknown command",
	}
}

type Move struct {
	Direction Direction
}

func (self *Move) Execute(state *GameState) (returnedState *GameState, result Result) {
	returnedState = state

	if self.Direction == UNKNOWN {
		result = &Noop{
			Description: "Can't walk that way.",
		}
		return
	}

	player := state.Player
	currentRoom := state.Player.CurrentLocation

	for _, door := range currentRoom.doors {
		if self.Direction == door.facing {
			if door.locked {
				result = &Noop{
					Description: "The door is locked.",
				}
			} else {
				player.CurrentLocation = door.Open()
				result = &Ok{
					Description: "You walked " + directionToStr(self.Direction) + ".",
				}
			}
			return
		}
	}

	result = &Noop{
		Description: "Can't walk that way.",
	}
	return
}
