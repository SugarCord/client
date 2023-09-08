package main

import (
	// Config
	"github.com/SugarCord/gateway/src/config"

	// Local Library
	"github.com/SugarCord/gateway/src/pkg/errorHandling"

	// Local Internal
	"github.com/SugarCord/gateway/src/internal/commands"
	"github.com/SugarCord/gateway/src/internal/eventHandler"

	// External Library
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Session Creation
	config.SESSION, config.ERR = discordgo.New("Bot " + config.DISCORD_TOKEN)
	errorHandling.FatalCheck(config.ERR)

	// Interaction Registration
	commands.RegisterBulk()

	// Handlers Instantiation
	eventHandler.Main()

	// WebSocket Connection
	config.ERR = config.SESSION.Open()
	errorHandling.FatalCheck(config.ERR)
	defer config.SESSION.Close()

	// Keep Alive
	<-make(chan struct{})
}
