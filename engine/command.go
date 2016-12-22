package engine

// a command - eg walk, fight, talk, etc. Commands alter the game state and return a result
type Command interface {
	Execute(state *GameState) (*GameState, Result)
}

func CommandDescriptions() string {
	return "walk <east|west|north|south>\n" +
		"rest\n"
}

// the result of a command
type Result struct {
	Invalid     bool   // command is not valid
	Noop        bool   // command didn't do anything
	Description string // a poetic output
}

// Eat/consume/drink
// Pickup/take
// Drop

type UnknownCommand struct{}

func (self *UnknownCommand) Execute(state *GameState) (*GameState, Result) {
	return state, Result{
		Invalid:     true,
		Noop:        true,
		Description: "Unknown command - type 'help' or enter a valid command.",
	}
}

//--------
// Resting
//--------

type Rest struct {
}

func (self *Rest) Execute(state *GameState) (returnedState *GameState, result Result) {
	returnedState = state

	returnedState.Player.Heal(1)

	result = Result{
		Description: "You rest a little bit.",
	}

	return
}

//-------
// Moving
//-------

type Move struct {
	Direction Direction
}

func (self *Move) Execute(state *GameState) (returnedState *GameState, result Result) {
	returnedState = state

	if self.Direction == UNKNOWN {
		result = Result{
			Description: "Can't walk that way.",
		}
		return
	}

	player := state.Player
	currentRoom := state.Player.CurrentLocation

	for _, door := range currentRoom.Doors {
		if self.Direction == door.facing {
			if door.Locked {
				result = Result{
					Description: "The door is Locked.",
				}
			} else {
				player.CurrentLocation = door.Open()
				result = Result{
					Description: "You walked " + DirectionToStr(self.Direction) + ".",
				}
			}
			return
		}
	}

	result = Result{
		Noop:        true,
		Description: "Can't walk that way.",
	}
	return
}
