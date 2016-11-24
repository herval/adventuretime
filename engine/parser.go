package engine

import "strings"

type CommandParser interface {
	ParseCommand(cmd string) Command
}

// parse command line strings and converts them to valid commands
type StandardParser struct {
}

func (parser *StandardParser) ParseCommand(cmd string) Command {
	return parseCommand(cmd)
}

func parseDirection(dir string) Direction {
	switch dir {
	case "east", "left":
		return EAST
	case "west", "right":
		return WEST
	case "forward", "up", "north":
		return NORTH
	case "back", "down", "south":
		return SOUTH
	default:
		return UNKNOWN
	}
}

func parseCommand(cmd string) Command {
	tokens := strings.Split(cmd, " ")
	if len(tokens) < 1 {
		return &UnknownCommand{}
	}

	invisibles := " \n\r\t"
	op := strings.ToLower(strings.Trim(tokens[0], invisibles))
	switch op {
	case "walk", "move", "go":
		if len(tokens) < 2 {
			return &UnknownCommand{}
		}
		dir := parseDirection(strings.ToLower(strings.Trim(tokens[1], invisibles)))
		if dir != UNKNOWN {
			return &Move{
				Direction: dir,
			}
		}
	case "help", "?":
		return &Help{}
	}
	return &UnknownCommand{}
}

type Help struct{}

func (self *Help) Execute(state *GameState) (*GameState, Result) {
	return state, &Noop{
		Description: self.Describe(),
	}
}

func (self *Help) Describe() string {
	return "Supported commands:\n\n'walk <east|west|north|south>'\n"
}
