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
	if dir == "east" || dir == "left" {
		return EAST
	}
	if dir == "west" || dir == "right" {
		return WEST
	}
	if dir == "forward" || dir == "up" || dir == "north" {
		return NORTH
	}
	if dir == "back" || dir == "down" || dir == "south" {
		return SOUTH
	}
	return UNKNOWN
}

func parseCommand(cmd string) Command {
	tokens := strings.Split(cmd, " ")
	if len(tokens) < 2 {
		return &UnknownCommand{}
	}

	invisibles := " \n\r\t"
	op := strings.ToLower(strings.Trim(tokens[0], invisibles))
	if op == "walk" || op == "move" || op == "go" {
		dir := parseDirection(strings.ToLower(strings.Trim(tokens[1], invisibles)))
		if dir != UNKNOWN {
			return &Move{
				Direction: dir,
			}
		}
	}

	return &UnknownCommand{}
}
