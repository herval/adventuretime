package engine

import (
	"strings"
	"math/rand"
)

// Generate random sentences based on a grammar

type Node struct {
	optional bool
	options  []string
	chance   int
	next     *Node
}

//[The] <location>{1} <of> <the>? [<adjective>? [[<kind>? <subject>?] | <plurals>]{1} ]{1}

func RandomName() string {
	return strings.Join(
		nonBlanks(
			atLeastOne(
				"The",
				of(locationAdjectives, 20),
				of(locations, 100),
			),
			oneOf(
				atLeastOne(
					"of the",
					of(adjectives, 30),
					of(kinds, 40),
					of(titles, 50),
				),
				atLeastOne(
					"of",
					of(adjectives, 30),
					of(plurals, 50),
				),
			),
		),
		" ",
	)
}

func nonBlanks(generators ...func() string) []string {
	res := make([]string, 0)
	for _, r := range generators {
		generated := r()
		if generated != "" {
			res = append(res, generated)
		}
	}
	return res
}

func oneOf(randomizers ...func() string) func() string {
	return func() string {
		res := nonBlanks(randomizers...)
		ret := res[rand.Intn(len(res) - 1)]
		return ret
	}
}

func atLeastOne(prefix string, randomizers ...func() string) func() string {
	maxTries := 100
	return func() string {
		for i := 0; i < maxTries; i++ {
			res := nonBlanks(randomizers...)

			if len(res) > 0 {
				str := strings.Join(res, " ")
				if prefix != "" {
					return prefix + " " + str
				} else {
					return str
				}
			}
		}
		panic("Couldn't possibly generate!")
		return ""
	}
}

// return a random string x% of the time (blank otherwise)
func of(options []string, chance int) func() string {
	return func() string {
		if rand.Intn(100) <= chance {
			return options[rand.Intn(len(options) - 1)]
		}
		return ""
	}
}

var locations = []string{
	"halls",
	"cave",
	"dungeon",
	"prison",
	"labyrinth",
	"maze",
	"pit",
	"crypt",
	"chambers",
	"vault",
	"oubliette",
	"underground city",
}

var locationAdjectives = []string{
	"mourning",
	"eternal",
	"bloody",
	"vanishing",
	"lost",
	"invisible",
	"relentless",
	"bleak",
	"forgotten",
	"great",
	"calamitous",
}

var adjectives = append(locationAdjectives,
	"dying",
	"dead",
	"crying",
	"longing",
	"everlasting",
	"perpetual",
	"murderous",
	"desperate",
	"hopeless",
	"doomed",
	"invincible",
	"unequaled",
)

var kinds = []string{
	"Goblin",
	"Elf",
	"half-blooded",
	"Dwarf",
	"Wizard",
	"Giant",
	"Troll",
}

var titles = []string{
	"king",
	"viscount",
	"prince",
	"duchess",
	"duke",
	"sage",
	"queen",
	"empire",
	"dragon",
	"monarch",
	"monster",
	"titan",
	"colossus",
	"master",
	"emperor",
	"folk",
}

var plurals = []string{
	"goblins",
	"elven",
	"dwarven",
	"wizards",
	"giants",
	"trolls",
	"kings",
	"sages",
	"queens",
	"dragons",
	"monarchs",
	"monsters",
	"titans",
	"bastards",
}
