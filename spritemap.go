package main


type Spritemap struct {
	SmallMonsters []*Frame
}

func NewSpritemap(frames map[string]Frame) Spritemap {
	

	return Spritemap{
		SmallMonsters: []*Frame{},
	}
}