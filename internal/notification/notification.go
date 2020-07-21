package notification

import (
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/tphume/DRBWG/internal/reminder"
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
	panic("implement me")
}

// Send notification
func (p *Process) send() error {
	panic("implement me")
}
