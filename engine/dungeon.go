package engine

import (
	"fmt"

	"github.com/herval/adventuretime/util"
	"math/rand"
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

	mainHall := RandomRoom(nil, youShallNotPass)
	mainHall.details = "This is the entrance of the Dungeon."

	util.Debug(fmt.Sprintf("Init dungeon: %s", mainHall))

	return &Dungeon{
		Entrance: mainHall,
	}
}


// =====

// Rooms are a linked list pointing to up to 4 other rooms through doors
type Room struct {
	doors   []*Door
	props []*Prop
	npcs []*Npc
	details string
}

func (r *Room) Describe() string {
	var str = "You are in a room."

	doors := len(r.doors)

	if doors == 1 {
		str += " There's a door facing " + directionToStr(r.doors[0].facing) + "."
	} else if doors > 1 {
		str += " There are doors facing "
		for i := 0; i < doors-1; i++ {
			str += directionToStr(r.doors[i].facing)
			if i < doors-2 {
				str += ", "
			}
		}
		str += " and " + directionToStr(r.doors[doors-1].facing) + ". "
	}

	if len(r.npcs) > 0 {
		descriptions := ""
		for i, npc := range r.npcs {
			descriptions += npc.Describe()
			if i < len(r.npcs)-1 {
				descriptions += ", "
			}
		}
		str += "You see " + descriptions
	}

	if len(r.props) > 0 {
		descriptions := ""
		for i, prop := range r.props {
			descriptions += prop.Describe()
			if i < len(r.npcs)-1 {
				descriptions += ", "
			}
		}
		str += "There's " + descriptions
	}

	if r.details != "" {
		str += "\n" + r.details
	}

	return str
}

// entryDoor is the door from the *previous room* that leads to the new room
func RandomRoom(comingFromRoom *Room, entryDoor *Door) *Room {
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
		self.to = RandomRoom(self.from, self)
		util.Debug(fmt.Sprintf("New to: %s", self.to))
	} else {
		util.Debug(fmt.Sprintf("Moving to existing room %s", self.to))
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
		if rand.Int()%2 == 0 {
			doors = append(
				doors,
				&Door{
					facing: facing,
					from:   currentRoom,
					locked: false, // TODO lock some for fun
				})
		}
	}

	util.Debug(fmt.Sprintf("Gen Doors: %s", doors))

	return doors
}


// =====

type Prop struct {

}

func generateProps() []*Prop {
	return  make([]*Prop, 0)
}

func (n *Prop) Describe() string {
	return "a foo"
}

// =====

// non-playable stuff
type Npc struct {
	health int
}

func (n *Npc) Describe() string {
	return "a foo"
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