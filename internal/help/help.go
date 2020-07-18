package help

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tphume/DRBWG/internal/reminder"
)

// Expects an empty string or a sub-command
func Handle(cmd []string, _ *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	if len(cmd) == 0 {
		return newEmbedAll(), nil
	}

	_, ok := CMD[cmd[0]]
	if !ok {
		return newEmbedAll(), nil
	}

	return newEmbed(cmd[0]), nil
}

// Helper func to return MessageEmbed
func newEmbed(c string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       CMD[c].title,
		Description: CMD[c].brief,
		Color:       reminder.Color,
		Footer:      reminder.Footer,
		Author:      reminder.Author,
		Fields: []*discordgo.MessageEmbedField{
			{Name: CMD[c].key, Value: CMD[c].desc},
		},
	}
}

func newEmbedAll() *discordgo.MessageEmbed {
	f := make([]*discordgo.MessageEmbedField, 0)
	for _, d := range CMD {
		f = append(f, &discordgo.MessageEmbedField{
			Name:  d.title,
			Value: d.brief,
		})
	}

	return &discordgo.MessageEmbed{
		URL:         reminder.URL,
		Title:       "All Commands",
		Description: allBrief,
		Color:       reminder.Color,
		Footer:      reminder.Footer,
		Author:      reminder.Author,
		Fields:      f,
	}
}

// List of sub-commands
const (
	at     = "at"
	in     = "in"
	lsg    = "lsg"
	lsc    = "lsc"
	update = "update"
	del    = "del"
	all    = "all"
)

// List of brief description for sub-commands
const (
	allBrief    = "List of all sub-commands. Use ```drbwg help [sub-command]``` to view details of the sub-command."
	atBrief     = "Set a reminder at specified time."
	inBrief     = "Set a reminder for the current time + some given duration."
	lsgBrief    = "List all the pending reminders for the current guild."
	lscBrief    = "List all the pending reminders for the current channel."
	updateBrief = "Update a pending reminder with a new specified time."
	delBrief    = "Delete a pending reminder"
)

// List of description for specific sub-commands
const (
	atDesc = "at sub-command requires 2 argument.\nThe first is a date and time in the following format - ```Jan-2-15:04-MST-2006```\n" +
		"The second is a name to identify the reminder in the following format - reminderName.\n\n**Full example** ```drbwg at " +
		"Jan-2-15:04-MST-2006 reminderName```"
	inDesc = "in sub-command requires 2 argument.\nThe first is a duration in the following format (where the first number " +
		"denotes the number of hours and the second denotes the minutes) - ```1h20m```\n" +
		"The second is a name to identify the reminder in the following format - reminderName.\n\n**Full example** ```drbwg at " +
		"1h20m reminderName```"
	lsgDesc    = "lsg sub-command requires no argument"
	lscDesc    = "lsc sub-command requires no argument"
	updateDesc = "update sub-command requires 2 argument.\nThe first is the id of the reminder.\nThe second is either the " +
		"duration or date and time.\nSee the at and in sub-command for the format. \n\n**Full example** ```drbwg 432904329430 1h4m```"
	delDesc = "del sub-command requires one argument.\nThe id of the reminder to be deleted.\n\n**Full example** ```drbwg delete " +
		"32432432```"
)

// List of sub-commands and its description
var CMD = map[string]struct {
	title string
	brief string
	key   string
	desc  string
}{
	at:     {title: at, brief: atBrief, key: "How does __at__ command work?", desc: atDesc},
	in:     {title: in, brief: inBrief, key: "How does __in__ command work?", desc: inDesc},
	lsg:    {title: lsg, brief: lsgBrief, key: "How does __lsg__ command work?", desc: lsgDesc},
	lsc:    {title: lsc, brief: lscBrief, key: "How does __lsc__ command work?", desc: lscDesc},
	update: {title: update, brief: updateBrief, key: "How does __update__ command work?", desc: updateDesc},
	del:    {title: del, brief: delBrief, key: "How does __del__ command work?", desc: delDesc},
}
