package notification

import (
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/tphume/DRBWG/internal/reminder"
)

type Process struct {
	d *discordgo.Session
	c *cron.Cron
	r reminder.Pending
}

// Return new cron process with notification jobs attached
func NewProcess(c *cron.Cron) {
	panic("implement me")
}

// Job to be run every minute
func (p *Process) remind() {
	panic("implement me")
}

// Send notification
func (p *Process) send() error {
	panic("implement me")
}
