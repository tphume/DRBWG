package help

// Expects an empty string or a sub-command
func Handle(cmd []string) ([]string, error) {
	if len(cmd) == 0 {
		return CMD[all], nil
	}

	res, ok := CMD[cmd[0]]
	if !ok {
		return CMD[all], nil
	}

	return res, nil
}

// List of sub-commands
const (
	at     = "at"
	in     = "in"
	lsg    = "lsg"
	lsc    = "lsc"
	update = "update"
	del    = "del"
	all    = ""
)

// List of brief description for sub-commands
const (
	atBrief     = ":alarm_clock: **at** - Set a reminder at specified time"
	inBrief     = ":timer: **in** - Set a reminder for the current time + some given duration"
	lsgBrief    = ":books: **lsg** - List all the pending reminders for the current guild"
	lscBrief    = ":bookmark: **lsc** - List all the pending reminders for the current channel"
	updateBrief = ":point_up: **updated** - Update a pending reminder with a new specified time"
	delBrief    = ":red_circle: **del** - Delete a pending reminder"
)

// List of description for specific sub-commands
const (
	atDesc = "at sub-command requires 2 argument.\nthe first is a date and time in the following format - Jan-2-15:04-MST-2006.\n" +
		"the second is a name to identify the reminder in the following format - reminderName.\nFull example - **drbwg at " +
		"Jan-2-15:04-MST-2006 reminderName**"
	inDesc = "in sub-command requires 2 argument.\nthe first is a duration in the following format (where the first number " +
		"denotes the number of hours and the second denotes the minutes) - 1-20.\n" +
		"the second is a name to identify the reminder in the following format - reminderName.\nFull example - **drbwg at " +
		"1-20 reminderName**"
	lsgDesc    = "lsg sub-command requires no argument"
	lscDesc    = "lsc sub-command requires no argument"
	updateDesc = "update sub-command requires 2 argument.\nthe first is the id of the reminder.\nthe second is either the " +
		"duration or date and time.\nsee the at and in sub-command for the format. Full example - **drbwg 432904329430 1-4**"
	delDesc = "del sub-command requires one argument.\nthe id of the reminder to be deleted.\nFull example - **drbwg delete " +
		"32432432**"
)

// List of sub-commands and its description
var CMD = map[string][]string{
	at:     {atBrief, atDesc},
	in:     {inBrief, inDesc},
	lsg:    {lsgBrief, lsgDesc},
	lsc:    {lscBrief, lscDesc},
	update: {updateBrief, updateDesc},
	del:    {delBrief, delDesc},
	all:    {atBrief, inBrief, lsgBrief, lscBrief, updateBrief, delBrief},
}
