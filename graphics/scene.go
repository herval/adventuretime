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

// a scene is an array of things, all stacked over each other:
// ground, objects, decorations, etc.
type Scene struct {
	//Tiles     [][]string // [y][x] aka [row][col]
	FloorMap  [][]string // a list of actual *sprites* for floors and empty spaces(constants from spritemap)
	WallsMap  [][]string // a list of actual *sprites* for walls (constants from spritemap)
	SpriteMap [][]string // a list of actual *sprites* for things that move + decorations (constants from spritemap)
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
			case WallWithDecoration:
				wallsMap[row][col] = Wall // a decoration must go on a wall
				floorMap[row][col] = Wall
				sprites[row][col] = WallWithDecoration

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
func toSprites(tiles [][]string, floorMap [][]string, wallsMap [][]string) [][]string {
	res := empty(tiles, "")

	for row := 0; row < len(tiles); row++ {
		for col := 0; col < len(tiles[row]); col++ {
			isA := func(kind string) bool {
				return isTile(tiles, row, col, kind)
			}

			sprite := ""

			if isA(Nothing) {
				sprite = TheUnknown
			} else if isA(RoomFloor) {
				surrounds := surroundingScenario(floorMap, row, col)

				// TODO corners on intersections
				switch {
				case surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top != RoomFloor && surrounds.bottom == RoomFloor: // cap up
					sprite = random(FloorLeftRightTops)
				case surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top == RoomFloor && surrounds.bottom != RoomFloor: // cap up
					sprite = random(FloorLeftRightBottoms)
				case surrounds.right != RoomFloor && surrounds.left == RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor: // cap up
					sprite = random(FloorTopBottomRights)
				case surrounds.right == RoomFloor && surrounds.left != RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor: // cap up
					sprite = random(FloorTopBottomLefts)
				case surrounds.right == RoomFloor && surrounds.left == RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor: // corridor sideways
					sprite = random(FloorTopBottoms)
				case surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top == RoomFloor && surrounds.bottom == RoomFloor: // corridor up
					sprite = random(FloorLeftRights)
				case surrounds.top != RoomFloor && surrounds.left != RoomFloor && surrounds.right == RoomFloor && surrounds.bottom == RoomFloor: // top-left
					sprite = random(FloorTopLefts)
				case surrounds.top != RoomFloor && surrounds.right != RoomFloor && surrounds.left == RoomFloor && surrounds.bottom == RoomFloor: // top-right
					sprite = random(FloorTopRights)
				case surrounds.top != RoomFloor && surrounds.bottom == RoomFloor: // top
					sprite = random(FloorTops)
				case surrounds.left != RoomFloor && surrounds.right == RoomFloor: // left walls
					sprite = random(FloorLefts)
				case surrounds.right != RoomFloor && surrounds.left == RoomFloor: // right walls
					sprite = random(FloorRights)
				default: // everything else
					sprite = random(Floors)
				}
			} else if isA(Wall) { // left/right/bottom "walls" are just empty space w/ shadows
				floorSurrounds := surroundingScenario(floorMap, row, col)
				wallSurrounds := surroundingScenario(wallsMap, row, col)

				switch {
				case floorSurrounds.right == RoomFloor && floorSurrounds.left == Nothing:
					sprite = TheUnknown
				case  floorSurrounds.left == Nothing && wallSurrounds.right == Wall:
					sprite = TheUnknown
				case floorSurrounds.left == RoomFloor && floorSurrounds.right == Nothing:
					sprite = TheUnknown
				case floorSurrounds.right == Nothing && wallSurrounds.left == Wall:
					sprite = TheUnknown
				case floorSurrounds.bottom == Nothing && floorSurrounds.top == RoomFloor:
					sprite = TheUnknown
				case floorSurrounds.top == RoomFloor && wallSurrounds.bottom == Wall:
					sprite = TheUnknown
				 default:
					sprite = random(Walls)
				}
			} else if isA(WallWithDecoration) {
				sprite = BannerRed1
			} else if isA(Hero) {
				sprite = HeroArmed2
			} else if isA(BigMonster) {
				sprite = GorgonArmed
			} else if isA(SmallMonster) {
				sprite = GoblinArmed
			}

			if sprite != "" {
				res[row][col] = sprite
			}
		}
	}

	return res
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
