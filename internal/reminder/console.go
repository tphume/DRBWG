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

func (i *Console) ListFromGuild(args GuildListArgs) (*GuildListRes, error) {
	res := &GuildListRes{[]Reminder{}}
	for _, d := range i.Data {
		if d.GuildId == args.GuildId {
			res.Data = append(res.Data, d)
		}
	}

	return res, nil
}
