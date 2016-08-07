package main

import (
	"fmt"
	"math/rand"
)

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
		Debug("No 'to' set! Generating...")
		self.to = RandomRoom(self.from, self)
		Debug(fmt.Sprintf("New to: %s", self.to))
	} else {
		Debug(fmt.Sprintf("Moving to existing room %s", self.to))
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

	Debug(fmt.Sprintf("Gen Doors: %s", doors))

	return doors
}

// func DoorReversed(door *Door) *Door {
// 	Debug(fmt.Sprint("Reversing door %s", door))
// 	return &Door{
// 		// facing: DirectionOpposite(door.facing),
// 		locked: door.locked,
// 		to:     door.from,
// 		from:   door.to,
// 	}
// }
