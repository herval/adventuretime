package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
}

// Follow commands via replies
// Search for the princess
// Monsters

// https://muddesigner.codeplex.com

// Monster
// Attack
// Defend
// Flee

// Npc
// Interact
// Attack

// Food
// Water
// Potion
// Weapon - knife/sword
// Key
// Torch

// Cues
// Hear noises from direction (easier with low fear)
// See light/something moving (easier with torch)

// Each turn = half hour
// Day passed

// Exit door - closed initially

// Princess is a cat

// Dungeon - build in memory

func main() {
	reader := bufio.NewReader(os.Stdin)

	controller := NewController()

	for controller.state.player.hp > 0 {
		fmt.Println(controller.state.player.currentLocation.Describe())

		cmd, _ := reader.ReadString('\n')
		_, op := controller.Execute(cmd)
		fmt.Println(op.Describe())
	}
}
