package main

import (
	// Config
	"github.com/SugarCord/gateway/src/config"

	// Local Library
	"github.com/SugarCord/gateway/src/pkg/errorHandling"

	// Local Internal
	"github.com/SugarCord/gateway/src/internal/commandStruct"
	"github.com/SugarCord/gateway/src/internal/handler"

	// External Library
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Session Creation
	config.SESSION, config.ERR = discordgo.New("Bot " + config.DISCORD_TOKEN)
	errorHandling.FatalCheck(config.ERR)

	// Instantiate Handlers
	handler.Main()

	// Interaction Registration
	commandStruct.COMMANDS, config.ERR = config.SESSION.ApplicationCommandBulkOverwrite(config.DISCORD_APPLICATION_ID, "", commandStruct.COMMANDS)
	errorHandling.LogCheck(config.ERR)

	// WebSocket Connection
	config.ERR = config.SESSION.Open()
	errorHandling.FatalCheck(config.ERR)
	defer config.SESSION.Close()

	// Keep Alive
	<-make(chan struct{})
}
