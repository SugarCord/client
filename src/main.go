package main

import (
	"github.com/joho/godotenv"
	"os"
	"github.com/bwmarrin/discordgo"
)

func main() {
	godotenv.Load()
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
}
