package config

import (
	// Standard Library
	"os"

	// External Library
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	// Re-use
	ERR error

	// Environment
	DISCORD_TOKEN, DISCORD_APPLICATION_ID, DISCORD_PUBLIC_KEY, DISCORD_CLIENT_SECRET string

	// Session
	SESSION *discordgo.Session
)

func init() {
	godotenv.Load("../.env")
	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	DISCORD_APPLICATION_ID = os.Getenv("DISCORD_APPLICATION_ID")
	DISCORD_PUBLIC_KEY = os.Getenv("DISCORD_PUBLIC_KEY")
	DISCORD_CLIENT_SECRET = os.Getenv("DISCORD_CLIENT_SECRET")
}
