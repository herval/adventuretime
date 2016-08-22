package graphics

import "strings"

const (
	Nothing            = "."
	RoomFloor          = "_"
	Wall               = "|"
	Door               = "D"
	WallWithDecoration = "^"
	Hero               = "H"
	Placeholder        = "X" // occupied space for big objects
	Chair              = "C"
	Chest              = "Ü"
	ChestOpen          = "U"
	Table              = "T"
	BigMonster         = "M"
	SmallMonster       = "m"
)

var sprites = []string{
	Hero,
	Chair,
	Chest,
	ChestOpen,
	Placeholder,
	TableHorizontal,
	Table,
	BigMonster,
	SmallMonster,
}

// a scene is an array of things, all stacked over each other:
// ground, objects, decorations, etc
type Scene struct {
	Tiles   [][]string // [y][x] aka [row][col]
	Blipmap string
}

func NewScene(blipmap string) Scene {
	return Scene{
		Tiles:   breakDown(blipmap),
		Blipmap: blipmap,
	}
}

// func (s *Scene) IsIndoors(row int, col int) {
//     return s.IsTile(row, col, )
// }

func (s *Scene) IsTileOrOutOfBounds(row int, col int, kind string) bool {
	return s.IsTile(row, col, kind) || row < 0 || row >= len(s.Tiles) || col < 0 || col >= len(s.Tiles[row])
}

func (s *Scene) IsTile(row int, col int, kind string) bool {
	if row < 0 || row >= len(s.Tiles) || col < 0 || col >= len(s.Tiles[row]) {
		return false
	}
	return s.Tiles[row][col] == kind
}

func (s *Scene) IsSprite(row int, col int) bool {
	for _, kind := range sprites {
		if s.IsTile(row, col, kind) {
			return true
		}
	}
	return false
}

func breakDown(blipmap string) [][]string {
	lines := strings.Split(blipmap, "\n")

	res := make([][]string, len(lines))

	for i, l := range lines {
		res[i] = strings.Split(l, "")
	}

	return res
}
