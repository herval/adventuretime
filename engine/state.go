package engine

import "github.com/herval/adventuretime/util"

type GameState struct {
	Player  *Player
	Dungeon *Dungeon
}

func NewGame() *GameState {
	util.Debug("Generating dungeon...")
	dungeon := NewDungeon()

	util.Debug("Configuring character...")
	player := NewPlayer(dungeon.Entrance)

	return &GameState{
		Player:  player,
		Dungeon: dungeon,
	}
}
