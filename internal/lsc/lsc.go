package lsc

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
	"time"
)

type Handler struct {
	ChannelList reminder.ChannelListRepo
}

func (h *Handler) Handle(_ []string, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	res, err := h.ChannelList.ListFromChannel(reminder.ChannelListArgs{GuildId: m.GuildID, ChannelId: m.ChannelID})
	if err != nil {
		return nil, err
	}

	return format(res.Data), nil
}

func format(data []reminder.Reminder) *discordgo.MessageEmbed {
	now := time.Now()
	res := &discordgo.MessageEmbed{
		URL:         reminder.URL,
		Title:       "Channel pending reminders :face_with_monocle:",
		Description: fmt.Sprintf("This Channel has a total of %d pending reminders", len(data)),
		Color:       reminder.Color,
		Footer:      reminder.Footer,
		Author:      reminder.Author,
	}

	f := make([]*discordgo.MessageEmbedField, len(data))
	for i := 0; i < len(data); i++ {
		f[i] = &discordgo.MessageEmbedField{
			Name: data[i].Id,
			Value: fmt.Sprintf("**Name**: %s\n**Time**: %s\n**Remaining Duration**: %s\n",
				data[i].Name, data[i].T.Format("Mon Jan 2 15:04:05 MST 2006"), data[i].T.Sub(now)),
		}
	}

	res.Fields = f
	return res
}
