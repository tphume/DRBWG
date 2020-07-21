package notification

import (
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
func (p *Process) send(r []reminder.Reminder) error {
	panic("implement me")
}
