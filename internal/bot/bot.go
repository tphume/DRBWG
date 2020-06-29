package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
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

type Handle func(cmd []string) ([]string, error)

type Bot struct {
	session         *discordgo.Session
	msgCreateRoutes map[string]Handle
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

	return nil
}

func (b *Bot) addMsgCreateRoutes(event string, route string, h Handle) error {
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

	res, err := route(msg[2:])
	if err != nil {
		log.Println(err)
		_, _ = s.ChannelMessageSend(m.ChannelID, "An error occurred")
	}

	// Format response and send message
	fRes := strings.Join(res, "\n")
	_, _ = s.ChannelMessageSend(m.ChannelID, fRes)
}
