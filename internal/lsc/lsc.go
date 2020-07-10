package lsc

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Handler struct {
	ChannelList reminder.ChannelListRepo
}

func (h *Handler) Handle(_ []string, m *discordgo.MessageCreate) ([]string, error) {
	res, err := h.ChannelList.ListFromChannel(reminder.ChannelListArgs{GuildId: m.GuildID, ChannelId: m.ChannelID})
	if err != nil {
		return nil, err
	}

	return format(res.Data), nil
}

func format(data []reminder.Reminder) []string {
	res := make([]string, len(data)+1)

	res[0] = fmt.Sprintf("**This Channel has a total of %d pending reminders** :face_with_monocle:\n", len(data))
	for i := 0; i < len(data); i++ {
		res[i+1] = fmt.Sprintf("**---------------------------------------------------------**"+
			"\n**ID**: %s\n**Name**: %s\n**Time**: %s\n",
			data[i].Id, data[i].Name, data[i].T)
	}

	return res
}
