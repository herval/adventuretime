package graphics

import (
	"github.com/herval/adventuretime/engine"
	"strings"
)

// Convert a engine.Dungeon into a Blipmap

func DungeonToBlipmap(dungeon *engine.Dungeon) string {
	grid := dungeon.Blueprint.Grid

	str := make([][]string, len(grid))

	// TODO render doors

	for x, row := range grid {
		str[x] = make([]string, len(grid[x]))
		for y, col := range row {
			if col == 0 {
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