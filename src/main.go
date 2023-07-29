package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
        log.Fatal(err)
    }

    // Start the bot.
    err = client.Open()
    if err != nil {
        log.Fatal(err)
    }

}
