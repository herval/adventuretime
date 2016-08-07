package main

type GameState struct {
	player  *Player
	dungeon *Dungeon
}

func NewGame() *GameState {
	Debug("Generating dungeon...")
	dungeon := NewDungeon()

	Debug("Configuring character...")
	player := NewPlayer(dungeon.entrance)

	return &GameState{
		player:  player,
		dungeon: dungeon,
	}
}
