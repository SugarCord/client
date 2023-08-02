package commandstruct

import (
	"github.com/bwmarrin/discordgo"
)

var COMMANDS = []*discordgo.ApplicationCommand{
		{
			Name: "settings",
			Description: "Display the settings menu",
			Type: 1,
		},
		{
			Name: "help",
			Description: "Display the help menu",
			Type: 1,
		},
}
