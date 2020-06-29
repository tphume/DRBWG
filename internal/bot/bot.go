package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const DRBWG = "drbwg"

type Handle func(cmd []string) ([]string, error)

type Bot struct {
	MsgCreateRoutes map[string]Handle
}

func (b *Bot) HandleMsgCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	route, ok := b.MsgCreateRoutes[msg[1]]
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
