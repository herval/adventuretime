package main

// Rooms are a linked list pointing to up to 4 other rooms through doors
type Room struct {
	doors   []*Door
	details string
}

// func (self *Room) SetDoors(doors []*Door) {
// 	for _, door := range doors {
// 		door.from = self
// 	}

// 	self.doors = doors
// }

func (self *Room) Describe() string {
	var str = "You are in a room."

	doors := len(self.doors)

	if doors == 1 {
		str += " There's a door facing " + DirectionToStr(self.doors[0].facing) + "."
	} else if doors > 1 {
		str += " There are doors facing "
		for i := 0; i < doors-1; i++ {
			str += DirectionToStr(self.doors[i].facing)
			if i < doors-2 {
				str += ", "
			}
		}
		str += " and " + DirectionToStr(self.doors[doors-1].facing) + "."
	}

	if len(self.details) > 0 {
		str += "\n" + self.details
	}

	return str
}

// entryDoor is the door from the *previous room* that leads to the new room
func RandomRoom(comingFromRoom *Room, entryDoor *Door) *Room {
	Debug("Building new room...")

	room := &Room{} // TODO add stuff

	room.doors = generateDoors(comingFromRoom, room, entryDoor)
	entryDoor.to = room // god bless side effects :-|

	return room
}
