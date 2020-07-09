package reminder

import "log"

type Console struct {
	Data []Reminder
}

func (i *Console) Insert(args InsertArgs) error {
	log.Printf("%+v\n", args)
	i.Data = append(i.Data, args.Reminder)
	return nil
}
