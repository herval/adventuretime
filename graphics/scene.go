package graphics

import (
	"math/rand"
	"strings"
)

const (
	Nothing      = "."
	RoomFloor    = "_"
	Wall         = "|"
	Door         = "D"
	WallBanner   = "^"
	Hero         = "H"
	Placeholder  = "X" // occupied space for big objects
	Chair        = "C"
	Chest        = "Ü"
	ChestOpen    = "U"
	Table        = "T"
	BigMonster   = "M"
	SmallMonster = "m"
)

// a scene is an array of things, all stacked over each other:
// ground, objects, decorations, etc.
// data is in [y][x] aka [row][col]
type Scene struct {
	FloorMap  [][]*Sprite // a list of actual *sprites* for floors and empty spaces(constants from spritemap)
	WallsMap  [][]*Sprite // a list of actual *sprites* for walls (constants from spritemap)
	SpriteMap [][]*Sprite // a list of actual *sprites* for things that move + decorations (constants from spritemap)
}

func NewScene(blipmap string) Scene {
	tiles := breakDown(blipmap)

	// save a mapping only with floor tiles and one with renderable sprites
	floorMap := empty(tiles, Nothing)
	sprites := empty(tiles, "")
	wallsMap := empty(tiles, "")

	// convert the scene markers to actual sprites
	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			switch tiles[row][col] {
			case WallBanner:
				wallsMap[row][col] = Wall // a decoration must go on a wall
				floorMap[row][col] = Wall
				sprites[row][col] = WallBanner

			case Wall:
				wallsMap[row][col] = Wall
				floorMap[row][col] = Wall

			case RoomFloor:
				floorMap[row][col] = RoomFloor

			default:
				if tiles[row][col] != Nothing && tiles[row][col] != "" {
					floorMap[row][col] = RoomFloor // no sprites live in the void
					sprites[row][col] = tiles[row][col]
				}
			}
		}
	}

	return Scene{
		FloorMap:  toSprites(floorMap, floorMap, wallsMap),
		WallsMap:  toSprites(wallsMap, floorMap, wallsMap),
		SpriteMap: toSprites(sprites, floorMap, wallsMap),
	}
}

// convert the scene "markers" to actual sprites
// Incredible typing & crazy complex, I know. :-|
func toSprites(tiles [][]string, floorMap [][]string, wallsMap [][]string) [][]*Sprite {
	res := emptySprites(tiles)

	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			isA := func(kind string) bool {
				return isTile(tiles, row, col, kind)
			}

			var sprite *Sprite

			switch {
			case isA(Nothing):
				sprite = &TheUnknown

			case isA(RoomFloor):
				sprite = floor(floorMap, row, col)

			case isA(Wall):
				sprite = wall(floorMap, wallsMap, row, col)

			case isA(WallBanner):
				sprite = random(Banners)

			case isA(Table):
				if tiles[row][col+1] == Placeholder {
					sprite = &TableHorizontal
				} else {
					sprite = &TableVertical
				}
			case isA(Hero):
				sprite = &HeroArmed2

			case isA(Door):
				floorSurrounds := surroundingScenario(floorMap, row, col)
				switch {
				case floorSurrounds.top == Wall && floorSurrounds.bottom == Wall: // door sideways
				//sprite =random(DoorsSideways)
				case floorSurrounds.left == Wall && floorSurrounds.right == Wall: // door up/down
					sprite = random(DoorsFrontFacing)
				}

			case isA(BigMonster):
				sprite = random(BigMonsters)

			case isA(SmallMonster):
				sprite = random(SmallMonsters)
			}

			if sprite != nil {
				res[row][col] = sprite
			}
		}
	}

	return res
}

func floor(floorMap [][]string, row int, col int) *Sprite {
	surrounds := surroundingScenario(floorMap, row, col)

	// TODO corners on intersections
	switch {
	case surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top != RoomFloor && surrounds.bottom == RoomFloor: // cap up
		return random(FloorLeftRightTops)
	case surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top == RoomFloor && surrounds.bottom != RoomFloor: // cap up
		return random(FloorLeftRightBottoms)
	case surrounds.right != RoomFloor && surrounds.left == RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor: // cap up
		return random(FloorTopBottomRights)
	case surrounds.right == RoomFloor && surrounds.left != RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor: // cap up
		return random(FloorTopBottomLefts)
	case surrounds.right == RoomFloor && surrounds.left == RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor: // corridor sideways
		return random(FloorTopBottoms)
	case surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top == RoomFloor && surrounds.bottom == RoomFloor: // corridor up
		return random(FloorLeftRights)
	case surrounds.top != RoomFloor && surrounds.left != RoomFloor && surrounds.right == RoomFloor && surrounds.bottom == RoomFloor: // top-left
		return random(FloorTopLefts)
	case surrounds.top != RoomFloor && surrounds.right != RoomFloor && surrounds.left == RoomFloor && surrounds.bottom == RoomFloor: // top-right
		return random(FloorTopRights)
	case surrounds.top != RoomFloor && surrounds.bottom == RoomFloor: // top
		return random(FloorTops)
	case surrounds.left != RoomFloor && surrounds.right == RoomFloor: // left walls
		return random(FloorLefts)
	case surrounds.right != RoomFloor && surrounds.left == RoomFloor: // right walls
		return random(FloorRights)
	default: // everything else
		return random(Floors)
	}
}

func wall(floorMap [][]string, wallsMap [][]string, row int, col int) *Sprite {
	// left/right/bottom "walls" are just empty space w/ shadows
	floorSurrounds := surroundingScenario(floorMap, row, col)
	wallSurrounds := surroundingScenario(wallsMap, row, col)

	switch {
	case floorSurrounds.right == RoomFloor && floorSurrounds.left == Nothing:
		return &TheUnknown
	case floorSurrounds.left == Nothing && wallSurrounds.right == Wall:
		return &TheUnknown
	case floorSurrounds.left == RoomFloor && floorSurrounds.right == Nothing:
		return &TheUnknown
	case floorSurrounds.right == Nothing && wallSurrounds.left == Wall:
		return &TheUnknown
	case floorSurrounds.bottom == Nothing && floorSurrounds.top == RoomFloor:
		return &TheUnknown
	case floorSurrounds.top == RoomFloor && wallSurrounds.bottom == Wall:
		return &TheUnknown
	default:
		return random(Walls)
	}
}

// empty array, same size
func empty(array [][]string, defaultStr string) [][]string {
	res := make([][]string, len(array))
	for i, _ := range res {
		res[i] = make([]string, len(array[i]))
		for j, _ := range res[i] {
			res[i][j] = defaultStr
		}
	}
	return res
}

func emptySprites(array [][]string) [][]*Sprite {
	res := make([][]*Sprite, len(array))
	for i, _ := range res {
		res[i] = make([]*Sprite, len(array[i]))
	}
	return res
}

// AREUFREAKINKIDDINME.
//func dup(array [][]string) [][]string {
//	b := make([][]string, len(array))
//	for i := range b {
//		b[i] = make([]string, len(array[i]))
//		for j := range b[i] {
//			b[i][j] = array[i][j]
//		}
//	}
//
//	return b
//}

// is the given tile of a given kind
func isTile(tiles [][]string, row int, col int, kind string) bool {
	if row < 0 || row >= len(tiles) || col < 0 || col >= len(tiles[row]) {
		return false
	}

	return tiles[row][col] == kind
}

type Surroundings struct {
	top    string
	left   string
	right  string
	bottom string
}

func random(source []Sprite) *Sprite {
	// if len(source) > 1 {
	return &source[rand.Intn(len(source))]
	// } else {
	// return source[0]
	// }
}

// top/left/right/bottom
func surroundingScenario(floorMap [][]string, row int, col int) Surroundings {
	res := Surroundings{
		top:    Nothing,
		left:   Nothing,
		right:  Nothing,
		bottom: Nothing,
	}

	if row > 0 {
		res.top = floorMap[row-1][col]
	}
	if row < len(floorMap)-1 {
		res.bottom = floorMap[row+1][col]
	}
	if col < len(floorMap[row])-1 {
		res.right = floorMap[row][col+1]
	}
	if col > 0 {
		res.left = floorMap[row][col-1]
	}

	return res
}

func breakDown(blipmap string) [][]string {
	lines := strings.Split(blipmap, "\n")

	res := make([][]string, len(lines))

	for i, l := range lines {
		res[i] = strings.Split(l, "")
	}

	return res
}
