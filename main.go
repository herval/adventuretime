package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/herval/adventuretime/engine"
	"github.com/herval/adventuretime/twitter"
)

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

// each room has many things
// each thing can have many interactions
// each interaction changes tate from a to a'

func main() {
	mode := os.Getenv("MODE")
	if mode == "twitter" {
		twitterGame()
	} else {
		commandLineGame()
	}
}

func twitterGame() {
	started := time.Now()

	api := twitter.NewApi(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_TOKEN_SECRET"),
	)
	mentions := api.MentionsStream(started)

	parser := twitter.TweetParser{
		Name: os.Getenv("TWITTER_SCREEN_NAME"),
	}
	controller := engine.NewController()

	api.Post(controller.State.Describe())

	for {
		mention := <-mentions

		_, op := controller.Execute(parser.ParseCommand(mention.Text))

		api.Post(fmt.Sprintf("@%s %s", mention.User.ScreenName, op.Describe()))

		switch op.(type) {
		case *engine.Noop: // do nothing
			fmt.Println("Invalid command: ", mention.Text)
		default:
			api.Post(controller.State.Describe())
		}
	}
}

func commandLineGame() {
	reader := bufio.NewReader(os.Stdin)
	parser := engine.StandardParser{}
	controller := engine.NewController()

	for controller.State.Player.Hp > 0 {
		fmt.Println(controller.State.Describe())
		fmt.Print("> ")

		cmd, _ := reader.ReadString('\n')
		_, op := controller.Execute(parser.ParseCommand(cmd))
		fmt.Print(fmt.Sprintf("\n%s\n\n", op.Describe()))
	}
}
