package reminder

import (
	"log"
	"time"
)

type InsertConsole struct{}

func (i InsertConsole) Insert(t time.Time, name string, g string, c string) error {
	log.Printf("New Reminder: {time: %s}, {name: %s}, {guild: %s}, {channel: %s}\n", t, name, g, c)
	return nil
}
