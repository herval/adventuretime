package engine

import (
	"math/rand"

	"fmt"

	"github.com/herval/adventuretime/util"
	"github.com/meshiest/go-dungeon/dungeon"
	"strings"
)

type Dungeon struct {
	Name      string
	Entrance  *Room
	Rooms     []*Room
	Blueprint *dungeon.Dungeon
}

// Generate a new dungeon
func NewDungeon(size int, totalRooms int) *Dungeon {
	//totalRooms := rand.Intn(30) + 10
	//size := 100
	blueprint := dungeon.NewDungeon(size, totalRooms)

	// build a room map based on the blueprint
	rooms := make([]*Room, len(blueprint.Rooms))
	for i := 0; i < len(rooms); i++ {
		rooms[i] = &Room{
			Id:  i,
			Ref: blueprint.Rooms[i],
		}
	}

	// connect rooms and make doors
	for i := 0; i < len(rooms); i++ {
		for j := range AllDirections {
			// no door on that direction yet
			if !rooms[i].ContainsDoor(AllDirections[j]) {
				nextRoom, _ := connectedRoom(blueprint, rooms[i], rooms, AllDirections[j])

				util.DebugFmt("Searching", DirectionToStr(AllDirections[j]), "of", rooms[i].Id)
				// TODO this will generate different doors on both sides
				// TODO use the same door obj/type for opposing doors?

				if nextRoom != nil && nextRoom != rooms[i] {
					util.DebugFmt("Connecting rooms", rooms[i].Id, nextRoom.Id)

					door := NewDoor(
						rooms[i],
						nextRoom,
						AllDirections[j],
						false,
						randomDoorKind(),
					)
					rooms[i].Doors = append(rooms[i].Doors, door)

					//if !nextRoom.ContainsDoor(wall) {
					//	door2 := NewDoor(
					//		nextRoom,
					//		rooms[i],
					//		wall,
					//		false,
					//		randomDoorKind(),
					//	)
					//	nextRoom.Doors = append(nextRoom.Doors, door2)
					//}

					// TODO generate the reverse door here instead of searching backwards
				}
			} else {
				util.DebugFmt("Door already placed on", rooms[i].Id)
			}
		}
	}

	// pick the first room w/ space for an extra door as entrance, and stick an entrance to it
	var entrance *Room
	for i := range rooms {
		if len(rooms[i].Doors) < 4 {
			entrance = rooms[i]
			break
		}
	}

	// the main dungeon door in all it's glory
	for i := range AllDirections {
		if !entrance.ContainsDoor(AllDirections[i]) {
			youShallNotPass := &Door{
				facing: AllDirections[i],
				Locked: true,
			}
			entrance.Doors = append(entrance.Doors, youShallNotPass)
			entrance.details = "This is the entrance of the Dungeon."

		}
	}

	//mainHall := randomRoom(nil, youShallNotPass)
	//mainHall.details = "This is the entrance of the Dungeon."

	//util.DebugFmt("Init dungeon: %s", mainHall)

	return &Dungeon{
		Name:        RandomName(),
		Entrance:    entrance,
		Rooms:       rooms,
		Blueprint:   blueprint,
	}
}

func contains(room *Room, x int, y int) bool {
	return x >= room.Ref.X &&
		y >= room.Ref.Y &&
		x <= room.Ref.X + room.Ref.Width - 1 &&
		y <= room.Ref.Y + room.Ref.Height - 1
}

//recursively move on the same direction until a room is found
func findRoomForward(blueprint *dungeon.Dungeon, currentX int, currentY int, origin *Room, rooms []*Room, ignoreDirection Direction) (*Room, Direction) {

	if blueprint.Grid[currentX][currentY] == 1 {
		printGrid(blueprint.Grid, currentX, currentY, "X")

		// is it a room?
		// TODO optimize w/ a lookup map
		for i := range rooms {
			if contains(rooms[i], currentX, currentY) {
				printGrid(blueprint.Grid, currentX, currentY, "!")
				var directionOnNewRoom Direction
				if !contains(rooms[i], currentX - 1, currentY) {
					directionOnNewRoom = NORTH
				} else if !contains(rooms[i], currentX + 1, currentY) {
					directionOnNewRoom = SOUTH
				} else if !contains(rooms[i], currentX, currentY - 1) {
					directionOnNewRoom = EAST
				} else if !contains(rooms[i], currentX, currentY + 1) {
					directionOnNewRoom = WEST
				}

				util.DebugFmt("Found room #", rooms[i].Id, "at", currentX, currentY, "facing", DirectionToStr(directionOnNewRoom))
				return rooms[i], directionOnNewRoom
			}
		}
	}

	directionsToSearch := DirectionsMinus(ignoreDirection)

	for i := range directionsToSearch {
		var x, y int = currentX, currentY

		switch directionsToSearch[i] {
		case NORTH:
			y = currentY + 1
		case SOUTH:
			y = currentY - 1
		case EAST:
			x = currentX + 1
		case WEST:
			x = currentX - 1
		}

		// don't move back to current position
		back := DirectionOpposite(directionsToSearch[i])

		if blueprint.Grid[x][y] == 1 {
			return findRoomForward(blueprint, x, y, origin, rooms, back)
		}
	}
	return nil, UNKNOWN
}
func printGrid(ints [][]int, x int, y int, current string) {
	if !util.DEBUGGING {
		return
	}

	str := ""
	for i := range ints {
		for j := range ints[i] {
			if i == x && j == y {
				str += current
			} else {
				if ints[i][j] == 1 {
					str += "#"
				} else {
					str += " "
				}
			}
		}
		str += "\n"
	}

	util.Debug(str)
	//fmt.Println(str)
}

// find a room that connects to the origin room via a corridor
// returns the room + which of its walls was reached
func connectedRoom(blueprint *dungeon.Dungeon, origin *Room, rooms []*Room, direction Direction) (*Room, Direction) {
	var startX, endX, startY, endY int = 0, 0, 0, 0

	// search along one of the room walls (luckily they're all rectangular!)
	switch direction {
	case NORTH:
		startX = origin.Ref.X
		endX = origin.Ref.X + origin.Ref.Width - 1
		startY = origin.Ref.Y + origin.Ref.Height
		endY = startY
	case SOUTH:
		startX = origin.Ref.X
		endX = origin.Ref.X + origin.Ref.Width - 1
		startY = origin.Ref.Y - 1
		endY = startY
	case EAST:
		startX = origin.Ref.X + origin.Ref.Width
		endX = startX
		startY = origin.Ref.Y
		endY = origin.Ref.Y + origin.Ref.Height - 1
	case WEST:
		startX = origin.Ref.X - 1
		endX = startX
		startY = origin.Ref.Y
		endY = origin.Ref.Y + origin.Ref.Height - 1
	}

	util.DebugFmt("Searching around room", origin.Id, startX, endX, startY, endY, DirectionToStr(direction))

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			printGrid(blueprint.Grid, x, y, "?")

			if x < len(blueprint.Grid) && y < len(blueprint.Grid[x]) && blueprint.Grid[x][y] == 1 { // a passage!
				util.DebugFmt("Found", x, y)
				// follow the path until reaching a new room
				return findRoomForward(blueprint, x, y, origin, rooms, DirectionOpposite(direction))
			}
		}
	}
	return nil, UNKNOWN
}

// =====

// Rooms are a linked list pointing to up to 4 other rooms through Doors
type Room struct {
	Id      int
	Doors   []*Door           // Doors, duh
	props   []*Prop           // stuff you don't interact with
	npcs    []*Npc            // things that can kill you
	details string            // mood and ambiance
	Ref     dungeon.Rectangle // room location on blueprint
}

func (r *Room) Describe() string {
	var str = []string{
		"You are in a room",
	}

	if len(r.Doors) > 0 {
		doors := len(r.Doors)
		s := make([]string, doors - 1)
		for i := 0; i < doors - 1; i++ {
			s[i] = DirectionToStr(r.Doors[i].facing)
		}
		last := DirectionToStr(r.Doors[doors - 1].facing)
		isOrAre := ""
		if doors == 1 {
			isOrAre = "There is a door"
		} else {
			isOrAre = "There are Doors"
		}
		str = append(
			str,
			fmt.Sprintf(" %s facing %s and %s.", isOrAre, strings.Join(s, ", "), last),
		)
	}

	if len(r.npcs) > 0 {
		descriptions := make([]string, len(r.npcs))
		for i, npc := range r.npcs {
			descriptions[i] = npc.Description
		}
		str = append(
			str,
			fmt.Sprintf("You see %s.", strings.Join(descriptions, ", ")),
		)
	}

	if len(r.props) > 0 {
		descriptions := make([]string, len(r.props))
		for i, prop := range r.props {
			descriptions[i] = prop.Description
		}
		str = append(
			str,
			fmt.Sprintf("There's %s", strings.Join(descriptions, ", ")),
		)
	}

	if r.details != "" {
		str = append(
			str,
			r.details,
		)
	}

	return strings.Join(str, ". ")
}

func (room *Room) ContainsDoor(facing Direction) bool {
	for i := range room.Doors {
		if room.Doors[i].facing == facing {
			return true
		}
	}
	return false
}

// =====

// "Door" kinds
type DoorKind int

const (
	Wooden     DoorKind = iota
	Gated               = iota
	Passageway          = iota
)

// The Doors
type Door struct {
	facing Direction // direction
	to     *Room     // these might be figured out only after the door is open
	from   *Room     // as the labirinth is generated on the fly
	Locked bool
	kind   DoorKind
}

func NewDoor(from *Room, to *Room, facing Direction, locked bool, kind DoorKind) *Door {
	return &Door{
		facing: facing,
		to:     to,
		from:   from,
		Locked: locked,
		kind:   kind,
	}
}

func randomDoorKind() DoorKind {
	switch rand.Intn(3) {
	case 1:
		return Gated
	case 2:
		return Passageway
	}
	return Wooden
}

// -----
// Props (decorative stuff you can't interact with)
// -----

type Prop struct {
	Description string
}

func generateProps() []*Prop {
	props := make([]*Prop, rand.Intn(4))
	for i := 0; i < len(props); i++ {
		props[i] = randomProp()
	}
	return props
}

var propTypes = []string{
	"a red banner decorates the wall",
	"a yellow banner decorates the wall",
	"a torch flickers on the wall",
	"the rotting remains of an adventurer lay on the floor",
	"a pool of blood on the floor",
	"a wooden table",
}

func randomProp() *Prop {
	return &Prop{
		Description: propTypes[rand.Intn(len(propTypes) - 1)],
	}
}

// =====

// non-playable stuff
type Npc struct {
	Health      int
	Hostile     bool
	Description string
}

func generateNpcs() []*Npc {
	res := make([]*Npc, 0)

	// a creature
	if rand.Intn(100) <= 20 {

	}

	// a gang
	if rand.Intn(100) <= 10 {

	}

	// a boss
	if rand.Intn(100) <= 5 {

	}

	// the princess!
	if rand.Intn(100) == 1 {

	}

	return res
}
