package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/bot"
	"log"
	"os"
)

func main() {
	var drbwg *bot.Bot

	token := os.Getenv("DRBWG_TOKEN")
	debug := os.Getenv("DEBUG")

	// Create discord session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	// Create and Run DRBWG Bot
	if debug != "" {
		drbwg = bot.NewDebug(s)
	} else {
		drbwg = bot.New(s)
	}

	log.Fatal(drbwg.Run())
}
