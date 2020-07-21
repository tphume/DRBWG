package notification

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/tphume/DRBWG/internal/reminder"
	"log"
	"time"
)

type Process struct {
	D *discordgo.Session
	R reminder.Pending
}

func (p *Process) AddToCron(c *cron.Cron) (cron.EntryID, error) {
	return c.AddFunc("* * * * *", p.remind)
}

// Job to be run every minute
func (p *Process) remind() {
	now := time.Now()
	start := now.Truncate(time.Minute)
	end := now.Round(time.Minute)

	// Get list of reminders
	gpArgs := reminder.GetPendingArgs{Start: start, End: end}
	gpRes, err := p.R.GetPending(gpArgs)
	if err != nil {
		panic(err)
	}

	// Send reminders to discord channels
	if err := p.send(gpRes.Data); err != nil {
		log.Panic(err)
	}

	// Update state of reminder
	usArgs := reminder.StateArgs{Start: start, End: end}
	if err := p.R.UpdateState(usArgs); err != nil {
		log.Panic(err)
	}
}

// Send notification
func (p *Process) send(reminders []reminder.Reminder) error {
	for _, r := range reminders {
		_, err := p.D.ChannelMessageSendEmbed(r.ChannelId, &discordgo.MessageEmbed{
			URL:    reminder.URL,
			Title:  fmt.Sprintf("REMINDER ALERT for %s :red_circle:", r.Name),
			Color:  reminder.Color,
			Footer: reminder.Footer,
			Author: reminder.Author,
			Fields: []*discordgo.MessageEmbedField{
				{Name: "ID", Value: r.Id},
				{Name: "Name", Value: r.Name},
				{Name: "Timestamp", Value: r.T.Format("Mon Jan 2 15:04:05 MST 2006")},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
