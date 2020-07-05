package reminder

import "log"

type InsertConsole struct{}

func (i InsertConsole) Insert(args InsertArgs) error {
	log.Println(args)
}
