package in

import (
	"errors"
	"time"
)

type Handler struct {
	insert insertRepo
}

func (h *Handler) Handle(cmd []string) ([]string, error) {
	panic("implement me")
}

// Represent connection to data source
type insertRepo interface {
	AddRecord()
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
	INVALID_DURATION = errors.New("invalid duration format")
	INVALID_NAME     = errors.New("invalid reminder name")
)
