package in

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Handler struct {
	insert insertRepo
}

func (h *Handler) Handle(cmd []string) ([]string, error) {
	if len(cmd) < 2 {
		return nil, INVALID_INPUT
	}

	dur, name := cmd[0], strings.TrimSpace(cmd[1])
	t, err := parse(dur, name)
	if err != nil {
		return nil, err
	}

	if err := h.insert.AddReminder(t, name); err != nil {
		return nil, err
	}

	return []string{
		"---- **Reminder Added** ----",
		fmt.Sprintf("**Name**: %s", name),
		fmt.Sprintf("**Time**: %s", t),
	}, nil
}

// Represent connection to data source
type insertRepo interface {
	AddReminder(t time.Time, name string) error
}

// Helper function to validate input and return timestamp
func parse(dur string, name string) (time.Time, error) {
	// Parses the duration string
	d, err := time.ParseDuration(dur)
	if err != nil || d < 1*time.Minute {
		return time.Time{}, INVALID_DURATION
	}

	// Check length of name
	if len(name) < 3 {
		return time.Time{}, INVALID_NAME
	}

	return time.Now().Add(d).UTC(), nil
}

// List of errors
var (
	INVALID_INPUT    = errors.New("invalid input. not enough arguments")
	INVALID_DURATION = errors.New("invalid duration format")
	INVALID_NAME     = errors.New("invalid reminder name")
)
