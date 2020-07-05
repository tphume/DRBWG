package reminder

import "log"

type InsertConsole struct{}

func (i InsertConsole) Insert(args InsertArgs) error {
	log.Printf("%+v\n", args)
	return nil
}
