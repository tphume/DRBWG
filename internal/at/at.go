package at

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/tphume/DRBWG/internal/reminder"
	"strings"
	"time"
)

type Handler struct {
	Setter reminder.SetRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	if len(cmd) < 2 {
		return nil, reminder.ErrNotEnoughArgs
	}

	now := time.Now()
	dur, name := cmd[0], strings.TrimSpace(cmd[1])

	t, err := parse(dur, name, now)
	if err != nil {
		return nil, err
	}

	args := reminder.SetArgs{
		Reminder: reminder.Reminder{
			Id:        uuid.New().String(),
			GuildId:   m.GuildID,
			ChannelId: m.ChannelID,
			T:         t,
			Name:      name,
		},
	}

	if err := h.Setter.Set(args); err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		URL:         reminder.URL,
		Title:       "Added new reminder :white_check_mark:",
		Description: "A new reminder has successfully been added",
		Color:       reminder.Color,
		Footer:      reminder.Footer,
		Author:      reminder.Author,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "ID", Value: args.Id},
			{Name: "Name", Value: name},
			{Name: "Timestamp", Value: t.Format("Mon Jan 2 15:04:05 MST 2006")},
			{Name: "Remaining Duration", Value: t.Sub(now).String()},
		},
	}, nil
}

// Helper function to validate input and return timestamp
func parse(dur string, name string, now time.Time) (time.Time, error) {
	// Parses the timestamp string
	t, err := time.Parse("Jan-2-15:04-MST-2006", dur)
	if err != nil || t.Before(now) {
		return time.Time{}, reminder.ErrBadArgs
	}

	// Check length of name
	if len(name) < 3 {
		return time.Time{}, reminder.ErrBadArgs
	}

	return t, nil
}
