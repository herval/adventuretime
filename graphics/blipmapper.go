package graphics

import (
	"github.com/herval/adventuretime/engine"
	"strings"
	"math/rand"
	"github.com/herval/adventuretime/util"
)

// Convert a engine.Dungeon into a Blipmap

func DungeonToBlipmap(dungeon *engine.Dungeon, player *engine.Player) string {
	grid := dungeon.Blueprint.Grid

	str := make([][]string, len(grid))

	playerX, playerY := -1, -1
	if player != nil {
		playerX = player.CurrentLocation.Ref.X + rand.Intn(player.CurrentLocation.Ref.Width) - 3
		playerX = util.Max(player.CurrentLocation.Ref.X+1, playerX)
		playerY = player.CurrentLocation.Ref.Y + rand.Intn(player.CurrentLocation.Ref.Height) - 3
		playerY = util.Max(player.CurrentLocation.Ref.Y+1, playerY)
	}

	// TODO render doors
	// TODO render room details

	for x, row := range grid {
		str[x] = make([]string, len(grid[x]))
		for y, col := range row {
			if x == playerX && y == playerY {
				str[x][y] = Hero
			} else if col == 0 {
				str[x][y] = Nothing
			} else {
				str[x][y] = RoomFloor
			}
		}
	}

	// TODO OMG THIS IS HORRIBLE
	// slap walls around rooms
	for x := range str {
		for y := range str[x] {
			if str[x][y] == RoomFloor {
				if y < len(str[x]) - 1 && str[x][y + 1] == Nothing {
					str[x][y + 1] = Wall
				}
				if y > 0 && str[x][y - 1] == Nothing {
					str[x][y - 1] = Wall
				}
				if x < len(str) - 1 && str[x + 1][y] == Nothing {
					str[x + 1][y] = Wall
				}
				if x > 0 && str[x - 1][y] == Nothing {
					str[x - 1][y] = Wall
				}
				if x < len(str) - 1 && y < len(str[x]) - 1 && str[x + 1][y + 1] == Nothing {
					str[x + 1][y + 1] = Wall
				}
				if x > 0 && y > 0 && str[x - 1][y - 1] == Nothing {
					str[x - 1][y - 1] = Wall
				}
				if x < len(str) - 1 && y > 0 && str[x + 1][y - 1] == Nothing {
					str[x + 1][y - 1] = Wall
				}
				if x > 0 && y < len(str[x]) - 1 && str[x - 1][y + 1] == Nothing {
					str[x - 1][y + 1] = Wall
				}
			}
		}
	}

	lines := ""
	for _, str := range str {
		lines += strings.Join(str, "") + "\n"
	}

	return lines
}
