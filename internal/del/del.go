package del

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Handler struct {
	DelRepo reminder.DelRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) ([]string, error) {
	panic("implement me")
}
