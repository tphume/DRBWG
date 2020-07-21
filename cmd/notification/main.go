package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/robfig/cron/v3"
	"github.com/tphume/DRBWG/internal/notification"
	"github.com/tphume/DRBWG/internal/reminder"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	token := os.Getenv("DRBWG_TOKEN")
	psqlUri := os.Getenv("PSQL_URI")

	// Create discord session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	// Create psql connection pool
	temp, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pool, err := pgxpool.Connect(temp, psqlUri)

	// Create Repo for Process
	repo := reminder.PsqlPending{Pool: pool}

	// Create cron and process
	c := cron.New()
	defer c.Stop()

	p := notification.Process{D: s, R: repo}
	_, err = p.AddToCron(c)
	if err != nil {
		log.Fatal(err)
	}

	c.Start()

	// Wait for signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
