package update

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/tphume/DRBWG/internal/reminder"
	"time"
)

type Handler struct {
	Up reminder.UpdateRepo
}

func (h *Handler) Handle(cmd []string, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	if len(cmd) < 2 {
		return nil, reminder.ErrNotEnoughArgs
	}

	now := time.Now()
	req, id := cmd[1], cmd[0]

	t, err := parse(req, id, now)
	if err != nil {
		return nil, err
	}

	args := reminder.UpdateArgs{
		Reminder: reminder.Reminder{
			Id:        id,
			GuildId:   m.GuildID,
			ChannelId: m.ChannelID,
			T:         t,
		},
	}

	if err := h.Up.Update(args); err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		URL:         reminder.URL,
		Title:       "Updated reminder :ballot_box_with_check: ",
		Description: "A reminder has successfully been updated",
		Color:       reminder.Color,
		Footer:      reminder.Footer,
		Author:      reminder.Author,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "ID", Value: args.Id},
			{Name: "Name", Value: args.Name},
			{Name: "Timestamp", Value: t.Format("Mon Jan 2 15:04:05 MST 2006")},
			{Name: "Remaining Duration", Value: t.Sub(now).String()},
		},
	}, nil
}

// Helper function to validate input and return timestamp
func parse(dur string, id string, now time.Time) (time.Time, error) {
	// Check uuid
	if _, err := uuid.Parse(id); err != nil {
		return time.Time{}, reminder.ErrBadArgs
	}

	// Parses the timestamp string
	t, err1 := time.Parse("Jan-2-15:04-MST-2006", dur)
	d, err2 := time.ParseDuration(dur)

	if err1 != nil && err2 != nil {
		return time.Time{}, reminder.ErrBadArgs
	} else if err1 != nil && d > 1*time.Minute {
		return now.Add(d).UTC(), nil
	} else if err2 != nil && !t.Before(now) {
		return t, reminder.ErrBadArgs
	}

	return time.Time{}, reminder.ErrBadArgs
}
