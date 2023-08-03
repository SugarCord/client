package main

import (
	// Config
	"github.com/SugarCord/gateway/src/config"

	// Local Library
	"github.com/SugarCord/gateway/src/pkg/errorHandling"

	// Local Internal
	"github.com/SugarCord/gateway/src/internal/commandStruct"
	"github.com/SugarCord/gateway/src/internal/handler"
	"github.com/SugarCord/gateway/src/internal/periodicCheck"

	// External Library
	"github.com/bwmarrin/discordgo"
)

func init() {
	// Session creation
	config.SESSION, config.ERR = discordgo.New("Bot " + config.DISCORD_TOKEN)
	errorHandling.FatalCheck(config.ERR)

	// WebSocket connection
	config.ERR = config.SESSION.Open()
	errorHandling.FatalCheck(config.ERR)
	defer config.SESSION.Close()

	// Interaction registration
	commandStruct.COMMANDS, config.ERR = config.SESSION.ApplicationCommandBulkOverwrite(config.DISCORD_APPLICATION_ID, "", commandStruct.COMMANDS)
	errorHandling.LogCheck(config.ERR)
}

func main() {
	go handler.ApplicationCommand()
	go handler.GatewayEvent()
	go periodicCheck.Main()
}
