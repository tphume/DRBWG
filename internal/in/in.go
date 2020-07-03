package in

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
	"strings"
	"time"
)

type Handler struct {
	insert reminder.InsertRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) ([]string, error) {
	if len(cmd) < 2 {
		return nil, INVALID_INPUT
	}

	dur, name := cmd[0], strings.TrimSpace(cmd[1])
	t, err := parse(dur, name, time.Now())
	if err != nil {
		return nil, err
	}

	if err := h.insert.Insert(t, name, m.GuildID, m.ChannelID); err != nil {
		return nil, err
	}

	return []string{
		"---- **Reminder Added** ----",
		fmt.Sprintf("**Name**: %s", name),
		fmt.Sprintf("**Time**: %s", t),
	}, nil
}

// Helper function to validate input and return timestamp
func parse(dur string, name string, now time.Time) (time.Time, error) {
	// Parses the duration string
	d, err := time.ParseDuration(dur)
	if err != nil || d < 1*time.Minute {
		return time.Time{}, INVALID_DURATION
	}

	// Check length of name
	if len(name) < 3 {
		return time.Time{}, INVALID_NAME
	}

	return now.Add(d).UTC(), nil
}

// List of errors
var (
	INVALID_INPUT    = errors.New("invalid input. not enough arguments")
	INVALID_DURATION = errors.New("invalid duration format")
	INVALID_NAME     = errors.New("invalid reminder name")
)
