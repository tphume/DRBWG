package set

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Handler struct {
	Setter reminder.SetRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) ([]string, error) {
	panic("implement me")
}
