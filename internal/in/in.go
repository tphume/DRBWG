package in

import "time"

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
	panic("implement me")
}
