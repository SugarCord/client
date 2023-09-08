package commands

import (
	// Config
	"github.com/SugarCord/gateway/src/config"

	// Local Library
	"github.com/SugarCord/gateway/src/pkg/errorHandling"

	// External Library
	"github.com/bwmarrin/discordgo"
)

var (
	// List of Registered Application Commands
	commandsList = []*discordgo.ApplicationCommand {
		{
			Name: "help",
			Description: "Display the help menu",
			Type: 1,
		},
		{
			Name: "settings",
			Description: "Display the settings menu",
			Type: 1,
		},
	}

	// Map of Command Names to their Handlers
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {},
	}
)

func RegisterBulk() {
	commandsList, config.ERR = config.SESSION.ApplicationCommandBulkOverwrite(config.DISCORD_APPLICATION_ID, "", commandsList)
	errorHandling.LogCheck(config.ERR)
}
