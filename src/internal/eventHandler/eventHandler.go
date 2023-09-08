package eventHandler

import (
	// Config
	"github.com/SugarCord/gateway/src/config"

	// External Library
	"github.com/bwmarrin/discordgo"
)

func Main() {
	interactionCreate()
}

func interactionCreate() {
	// Exhaustive Interaction Handler Instantiation
	config.SESSION.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	})
}
