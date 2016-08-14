package main

import (
	"strings"

	"github.com/herval/adventuretime/engine"
)

type CommandParser interface {
	ParseCommand(cmd string) engine.Command
}

// parse command line strings and converts them to valid commands
type StandardParser struct {
}

func (parser *StandardParser) ParseCommand(cmd string) engine.Command {
	return parseCommand(cmd)
}

func parseDirection(dir string) engine.Direction {
	if dir == "east" || dir == "left" {
		return engine.EAST
	}
	if dir == "west" || dir == "right" {
		return engine.WEST
	}
	if dir == "forward" || dir == "up" || dir == "north" {
		return engine.NORTH
	}
	if dir == "back" || dir == "down" || dir == "south" {
		return engine.SOUTH
	}
	return engine.UNKNOWN
}

func parseCommand(cmd string) engine.Command {
	tokens := strings.Split(cmd, " ")
	if len(tokens) < 2 {
		return &engine.UnknownCommand{}
	}

	invisibles := " \n\r\t"
	op := strings.ToLower(strings.Trim(tokens[0], invisibles))
	if op == "walk" || op == "move" || op == "go" {
		dir := parseDirection(strings.ToLower(strings.Trim(tokens[1], invisibles)))
		if dir != engine.UNKNOWN {
			return &engine.Move{
				Direction: dir,
			}
		}
	}

	return &engine.UnknownCommand{}
}
