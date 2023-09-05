package main

import (
	// Config
	"github.com/SugarCord/gateway/src/config"

	// Local Library
	"github.com/SugarCord/gateway/src/pkg/errorHandling"

	// Local Internal
	"github.com/SugarCord/gateway/src/internal/commandStruct"

	// External Library
	"github.com/bwmarrin/discordgo"
)

func init() {
	// Keep Alive
	<-make(chan struct{})

	// Session Creation
	config.SESSION, config.ERR = discordgo.New("Bot " + config.DISCORD_TOKEN)
	errorHandling.FatalCheck(config.ERR)

	// WebSocket Connection
	config.ERR = config.SESSION.Open()
	errorHandling.FatalCheck(config.ERR)
	defer config.SESSION.Close()

	// Create Handlers

	// create an handler that will be called every time a new interaction is created, answering by an "hello world" message
	config.SESSION.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Interaction Response
		config.ERR = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello World!",
			},
		})
		errorHandling.LogCheck(config.ERR)
	})

	// Interaction Registration
	commandStruct.COMMANDS, config.ERR = config.SESSION.ApplicationCommandBulkOverwrite(config.DISCORD_APPLICATION_ID, "", commandStruct.COMMANDS)
	errorHandling.LogCheck(config.ERR)
}

func main() {
}
