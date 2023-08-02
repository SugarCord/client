package main

import (
	"os"
	"src/errorHandling"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	ERR error
	DISCORD_TOKEN, DISCORD_APPLICATION_ID, DISCORD_PUBLIC_KEY, DISCORD_CLIENT_SECRET string
	SESSION *discordgo.Session
	COMMANDS []*discordgo.ApplicationCommand = []*discordgo.ApplicationCommand{
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
	COMMANDS, ERR = SESSION.ApplicationCommandBulkOverwrite(DISCORD_APPLICATION_ID, "", COMMANDS)
	errorHandling.LogCheck(ERR)

}

// func main() {
// 	go ApplicationCommandHandler()
// 	go GatewayEventsHandler()
// 	go PeriodicCheck()
// }
