package in

import (
	"log"
	"time"
)

type InsertConsole struct{}

func (i InsertConsole) AddReminder(t time.Time, name string) error {
	log.Printf("New Reminder: {time: %s}, {name: %s}", t, name)
	return nil
}
