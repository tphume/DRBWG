package del

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Handler struct {
	DelRepo reminder.DelRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	if len(cmd) == 0 {
		return nil, reminder.ErrNotEnoughArgs
	}

	if _, err := uuid.Parse(cmd[0]); err != nil {
		return nil, reminder.ErrInvalidId
	}

	args := &reminder.DelArgs{
		Reminder: reminder.Reminder{
			Id:        cmd[0],
			GuildId:   m.GuildID,
			ChannelId: m.ChannelID,
		},
	}

	if err := h.DelRepo.Del(args); err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		URL:         reminder.URL,
		Title:       "Delete reminder :exclamation:",
		Description: "A reminder has successfully been deleted",
		Color:       reminder.Color,
		Footer:      reminder.Footer,
		Author:      reminder.Author,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "ID", Value: args.Id},
			{Name: "Name", Value: args.Name},
			{Name: "Timestamp", Value: args.T.Format("Mon Jan 2 15:04:05 MST 2006")},
		},
	}, nil
}
