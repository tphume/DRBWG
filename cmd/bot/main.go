package main

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tphume/DRBWG/internal/bot"
	"log"
	"os"
	"time"
)

func main() {
	var drbwg *bot.Bot

	token := os.Getenv("DRBWG_TOKEN")
	psqlUri := os.Getenv("PSQL_URI")
	debug := os.Getenv("DEBUG")

	// Create discord session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	// Create psql connection pool
	temp, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pool, err := pgxpool.Connect(temp, psqlUri)

	// Create and Run DRBWG Bot
	if debug != "" {
		drbwg = bot.NewDebug(s)
	} else {
		drbwg = bot.New(s, pool)
	}

	log.Fatal(drbwg.Run())
}
