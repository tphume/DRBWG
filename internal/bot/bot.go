package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tphume/DRBWG/internal/help"
	"github.com/tphume/DRBWG/internal/in"
	"github.com/tphume/DRBWG/internal/lsg"
	"github.com/tphume/DRBWG/internal/reminder"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	DRBWG = "drbwg"

	// List of event types
	MSG_CREATE = "MSG_CREATE"
)

type Handle func(cmd []string, m *discordgo.MessageCreate) ([]string, error)

type Bot struct {
	session         *discordgo.Session
	msgCreateRoutes map[string]Handle
}

// Return a new DRBWG bot with handlers attached
func New(s *discordgo.Session, pool *pgxpool.Pool) *Bot {
	b := &Bot{session: s, msgCreateRoutes: make(map[string]Handle)}

	// Create Repo
	psql := &reminder.Psql{Pool: pool}

	// Create in route
	inHandler := in.Handler{Insert: psql}
	lsgHandler := lsg.Handler{GuildList: psql}

	// Register routes
	if err := b.addMsgRoutes(MSG_CREATE, "help", help.Handle); err != nil {
		log.Fatal(err)
	}

	if err := b.addMsgRoutes(MSG_CREATE, "in", inHandler.Handle); err != nil {
		log.Fatal(err)
	}

	if err := b.addMsgRoutes(MSG_CREATE, "lsg", lsgHandler.Handle); err != nil {
		log.Fatal(err)
	}

	// Add handlers
	b.session.AddHandler(b.handleMsgCreate)

	return b
}

// Return new DRBWG bot with debugging handlers attached
func NewDebug(s *discordgo.Session) *Bot {
	b := &Bot{session: s, msgCreateRoutes: make(map[string]Handle)}

	// Create Repo
	console := &reminder.Console{Data: []reminder.Reminder{}}

	// Create in route
	inHandler := in.Handler{Insert: console}
	lsgHandler := lsg.Handler{GuildList: console}

	// Register handlers to routes
	if err := b.addMsgRoutes(MSG_CREATE, "help", help.Handle); err != nil {
		log.Fatal(err)
	}

	if err := b.addMsgRoutes(MSG_CREATE, "in", inHandler.Handle); err != nil {
		log.Fatal(err)
	}

	if err := b.addMsgRoutes(MSG_CREATE, "lsg", lsgHandler.Handle); err != nil {
		log.Fatal(err)
	}

	// Add handlers
	b.session.AddHandler(b.handleMsgCreate)

	return b
}

// Blocking call that connects the bot to discord
func (b *Bot) Run() error {
	if err := b.session.Open(); err != nil {
		return err
	}

	defer b.session.Close()
	log.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}

func (b *Bot) addMsgRoutes(event string, route string, h Handle) error {
	switch event {
	case MSG_CREATE:
		b.msgCreateRoutes[route] = h
	default:
		return errors.New("event type does not exist")
	}

	return nil
}

func (b *Bot) handleMsgCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore message created by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Trim and split the string by spaces
	// Return if command is not called
	msg := strings.Split(strings.TrimSpace(m.Content), " ")
	if len(msg) < 2 || msg[0] != DRBWG {
		return
	}

	// Route and handle
	route, ok := b.msgCreateRoutes[msg[1]]
	if !ok {
		return
	}

	res, err := route(msg[2:], m)
	if err != nil {
		log.Println(err)
		_, _ = s.ChannelMessageSend(m.ChannelID, "An error occurred")
	}

	// Format response and send message
	fRes := strings.Join(res, "\n")
	_, _ = s.ChannelMessageSend(m.ChannelID, fRes)
}
