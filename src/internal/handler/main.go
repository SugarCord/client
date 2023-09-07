package handler

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

func Main() {
	interaction()
}

func interaction() {

	// interaction registration with a switch


	config.SESSION.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	})
}