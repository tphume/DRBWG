package at

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/tphume/DRBWG/internal/reminder"
	"strings"
	"time"
)

type Handler struct {
	Setter reminder.SetRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) ([]string, error) {
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

	return []string{
		"**Reminder Added** :white_check_mark:",
		fmt.Sprintf("**ID**: %s", args.Id),
		"\n**---------------------------------------------------------**",
		fmt.Sprintf("**Name**: %s", name),
		fmt.Sprintf("**Time**: %s", t.Format("Mon Jan 2 15:04:05 MST 2006")),
		fmt.Sprintf("Will remind in **%s**", t.Sub(now)),
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
