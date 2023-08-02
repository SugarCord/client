package commandStruct

import "github.com/bwmarrin/discordgo"

var COMMANDS = []*discordgo.ApplicationCommand {
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
