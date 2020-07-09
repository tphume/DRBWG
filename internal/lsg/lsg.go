package lsg

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Handler struct {
	GuildList reminder.GuildListRepo
}

func (h *Handler) Handle(_ []string, m *discordgo.MessageCreate) ([]string, error) {
	res, err := h.GuildList.ListFromGuild(reminder.GuildListArgs{GuildId: m.GuildID})
	if err != nil {
		return nil, err
	}

	return format(res.Data), nil
}

func format(data []reminder.Reminder) []string {
	res := make([]string, len(data)+1)

	res[0] = fmt.Sprintf("**Retrieved a total of %d reminders for this guild**\n", len(data))
	for i := 1; i < len(data)+1; i++ {
		res[i] = fmt.Sprintf("**Reminder Added** :white_check_mark:\n**ID**: %s\n**Name**: %s\n**Time**: %s\n",
			data[i].Id, data[i].Name, data[i].T)
	}

	return res
}
