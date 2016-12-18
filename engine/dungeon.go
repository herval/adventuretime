package engine

import (
	"math/rand"

	"fmt"

	"github.com/herval/adventuretime/util"
	"strings"
)

type Dungeon struct {
	Entrance *Room
}

// Generate dungeon... lazily
func NewDungeon() *Dungeon {
	// the main dungeon door in all it's glory
	youShallNotPass := &Door{
		facing: NORTH,
		locked: true,
	}

	mainHall := randomRoom(nil, youShallNotPass)
	mainHall.details = "This is the entrance of the Dungeon."

	util.DebugFmt("Init dungeon: %s", mainHall)

	return &Dungeon{
		Entrance: mainHall,
	}
}

// =====

// Rooms are a linked list pointing to up to 4 other rooms through doors
type Room struct {
	doors   []*Door // doors, duh
	props   []*Prop // stuff you don't interact with
	npcs    []*Npc  // things that can kill you
	details string  // mood and ambiance
}

func (r *Room) Describe() string {
	var str = []string{
		"You are in a room",
	}

	if len(r.doors) > 0 {
		doors := len(r.doors)
		s := make([]string, doors - 1)
		for i := 0; i < doors - 1; i++ {
			s[i] = directionToStr(r.doors[i].facing)
		}
		last := directionToStr(r.doors[doors - 1].facing)
		isOrAre := ""
		if doors == 1 {
			isOrAre = "There is a door"
		} else {
			isOrAre = "There are doors"
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

// entryDoor is the door from the *previous room* that leads to the new room
func randomRoom(comingFromRoom *Room, entryDoor *Door) *Room {
	util.Debug("Building new room...")

	room := &Room{} // TODO add stuff

	room.doors = generateDoors(comingFromRoom, room, entryDoor)
	room.props = generateProps()
	room.npcs = generateNpcs()

	entryDoor.to = room // god bless side effects :-|

	return room
}

// =====

// The Doors
type Door struct {
	facing Direction // direction
	to     *Room     // these might be figured out only after the door is open
	from   *Room     // as the labirinth is generated on the fly
	locked bool
}

// open a door and move unto the unkown
func (self *Door) Open() *Room {
	if self.to == nil {
		util.Debug("No 'to' set! Generating...")
		self.to = randomRoom(self.from, self)
		util.DebugFmt("New to: %s", self.to)
	} else {
		util.DebugFmt("Moving to existing room %s", self.to)
	}

	return self.to
}

func generateDoors(previousRoom *Room, currentRoom *Room, enteringFrom *Door) []*Door {
	doors := []*Door{}

	// var potentialDirections []int

	// door "from" is always present, of course
	// if entryDoor != nil {
	entrance := &Door{
		facing: DirectionOpposite(enteringFrom.facing),
		to:     previousRoom,
		from:   currentRoom,
		locked: enteringFrom.locked,
	}
	doors = append(doors, entrance)
	// 1 to 3 doors on a room (plus the door where you came from, of course)
	potentialDirections := DirectionsMinus(entrance.facing)
	// } else {
	// 	potentialDirections = AllDirections
	// }

	for _, facing := range potentialDirections {
		if rand.Int() % 2 == 0 {
			doors = append(
				doors,
				&Door{
					facing: facing,
					from:   currentRoom,
					locked: false, // TODO lock some for fun
				})
		}
	}

	util.DebugFmt("Gen Doors: %s", doors)

	return doors
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
