package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/bot"
	"log"
	"os"
)

func main() {
	token := os.Getenv("DRBWG_TOKEN")

	// Create discord session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	// Create and Run DRBWG Bot
	drbwg := bot.New(s)
	log.Fatal(drbwg.Run())
}
