package main

import (
	// Local Library
	"github.com/SugarCord/gateway/src/pkg/errorHandling"

	// Local Internal

	commandstruct "github.com/SugarCord/gateway/src/internal/commandStruct"
	"github.com/SugarCord/gateway/src/internal/handler"
	periodicCheck "github.com/SugarCord/gateway/src/internal/periodicChecks"

	// Standard Library
	"os"

	// External Library
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	ERR error

	DISCORD_TOKEN, DISCORD_APPLICATION_ID, DISCORD_PUBLIC_KEY, DISCORD_CLIENT_SECRET string
	SESSION *discordgo.Session
)

func init() {

	// Load env from .env file
	godotenv.Load("../.env")
	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	DISCORD_APPLICATION_ID = os.Getenv("DISCORD_APPLICATION_ID")
	DISCORD_PUBLIC_KEY = os.Getenv("DISCORD_APPLICATION_ID")
	DISCORD_CLIENT_SECRET = os.Getenv("DISCORD_CLIENT_SECRET")

	// Session creation
	SESSION, ERR = discordgo.New("Bot " + DISCORD_TOKEN)
	errorHandling.FatalCheck(ERR)

	// WebSocket creation
	ERR = SESSION.Open()
	errorHandling.FatalCheck(ERR)
	defer SESSION.Close()

	// Interaction registration
	commandstruct.COMMANDS, ERR = SESSION.ApplicationCommandBulkOverwrite(DISCORD_APPLICATION_ID, "", commandstruct.COMMANDS)
	errorHandling.LogCheck(ERR)

}

func main() {
	go handler.ApplicationCommand()
	go handler.GatewayEvent()
	go periodicCheck.Main()
}
