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
	Tiles     [][]string // [y][x] aka [row][col]
	floorMap  [][]string
	spriteMap [][]string
	Blipmap   string
}

func NewScene(blipmap string) Scene {
	tiles := breakDown(blipmap)

	// save a mapping only with floor tiles and one with renderable sprites
	floorMap := dup(tiles)
	sprites := dup(tiles)

	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			if isSprite(tiles[row][col]) {
				floorMap[row][col] = RoomFloor
			} else if tiles[row][col] == WallWithDecoration {
				floorMap[row][col] = Wall
			} else if tiles[row][col] == Door {
				floorMap[row][col] = RoomFloor
			}

			if !isSprite(tiles[row][col]) && tiles[row][col] != WallWithDecoration {
				sprites[row][col] = Nothing
			}
		}
	}

	return Scene{
		Tiles:     tiles,
		Blipmap:   blipmap,
		floorMap:  floorMap,
		spriteMap: sprites,
	}
}

// AREUFREAKINKIDDINME.
func dup(array [][]string) [][]string {
	b := make([][]string, len(array))
	for i := range b {
		b[i] = make([]string, len(array[i]))
		for j := range b[i] {
			b[i][j] = array[i][j]
		}
	}

	return b
}

// func (s *Scene) IsIndoors(row int, col int) {
//     return s.IsTile(row, col, )
// }

// func (s *Scene) IsTileOrOutOfBounds(row int, col int, kind string) bool {
// 	return s.IsTile(row, col, kind) || row < 0 || row >= len(s.Tiles) || col < 0 || col >= len(s.Tiles[row])
// }

func (s *Scene) IsTile(row int, col int, kind string) bool {
	if row < 0 || row >= len(s.Tiles) || col < 0 || col >= len(s.Tiles[row]) {
		return false
	}

	if kind == Wall || kind == RoomFloor {
		return s.floorMap[row][col] == kind
	}

	return s.Tiles[row][col] == kind
}

type Surroundings struct {
	top    string
	left   string
	right  string
	bottom string
}

// top/left/right/bottom
func (s *Scene) SurroundingScenario(row int, col int) Surroundings {
	res := Surroundings{
		top:    Nothing,
		left:   Nothing,
		right:  Nothing,
		bottom: Nothing,
	}

	if row > 0 {
		res.top = s.floorMap[row-1][col]
	}
	if row < len(s.floorMap)-1 {
		res.bottom = s.floorMap[row+1][col]
	}
	if col < len(s.floorMap[row])-1 {
		res.right = s.floorMap[row][col+1]
	}
	if col > 0 {
		res.left = s.floorMap[row][col-1]
	}

	return res
}

func isSprite(kind string) bool {
	for _, spriteKind := range sprites {
		if kind == spriteKind {
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
