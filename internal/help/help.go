package help

// List of sub-commands
const (
	at     = "at"
	in     = "in"
	lsg    = "lsg"
	lsc    = "lsc"
	update = "update"
	del    = "del"
)

// List of description for specific sub-commands
const (
	atDesc = "at sub-command requires 2 argument.\nthe first is a date and time in the following format - Jan,2,15:04,MST,2006.\n" +
		"the second is a name to identify the reminder in the following format - reminderName.\nFull example - drbwg at " +
		"Jan,2,15:04,MST,2006 reminderName"
	inDesc = "in sub-command requires 2 argument.\nthe first is a duration in the following format - 1h20m.\n" +
		"the second is a name to identify the reminder in the following format - reminderName.\nFull example - drbwg at " +
		"1h20m reminderName"
	lsgDesc    = "lsg sub-command requires no argument"
	lscDesc    = "lsc sub-command requires no argument"
	updateDesc = "update sub-command requires 2 argument.\nthe first is the id of the reminder.\nthe second is either the " +
		"duration or date and time.\nsee the at and in sub-command for the format. Full example - drbwg 432904329430 1h4m "
	delDesc = "del sub-command requires one argument.\nthe id of the reminder to be deleted.\nFull example - drbwg delete " +
		"32432432"
)

// List of sub-commands and its description
var CMD = map[string][]string{
	at:     {"set a reminder at specified time", atDesc},
	in:     {"set a reminder for the current time + some given duration", inDesc},
	lsg:    {"list all the pending reminders for the current guild", lsgDesc},
	lsc:    {"list all the pending reminders for the current channel", lscDesc},
	update: {"update a pending reminder with a new specified time", updateDesc},
	del:    {"del a pending reminder", delDesc},
}

// Expects an empty string or a sub-command
func Handle(cmd string) (string, error) {
	panic("implement me")
}
