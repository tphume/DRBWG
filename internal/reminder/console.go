package reminder

import "log"

type InsertConsole struct {
	Data []Reminder
}

func (i *InsertConsole) Insert(args InsertArgs) error {
	log.Printf("%+v\n", args)
	i.Data = append(i.Data, args.Reminder)
	return nil
}
