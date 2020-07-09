package lsg

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Handler struct {
	GuildList reminder.GuildListRepo
}

func (h *Handler) Handle(_ []string, m *discordgo.MessageCreate) ([]string, error) {
	panic("implement me")
}
