package engine

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
	UNKNOWN
)

type Direction int

func DirectionOpposite(direction Direction) Direction {
	switch direction {
	case NORTH:
		return SOUTH
	case SOUTH:
		return NORTH
	case EAST:
		return WEST
	case WEST:
		return EAST
	}
	return UNKNOWN
}

var AllDirections = []Direction{NORTH, SOUTH, EAST, WEST}

func DirectionsMinus(direction Direction) []Direction {
	all := []Direction{}
	for i := 0; i < UNKNOWN; i++ {
		if Direction(i) != direction {
			all = append(all, Direction(i))
		}
	}
	return all
}

func directionToStr(direction Direction) string {
	switch direction {
	case NORTH:
		return "north"
	case EAST:
		return "east"
	case SOUTH:
		return "south"
	case WEST:
		return "west"
	}
	return "unknown"
}
