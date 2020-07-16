package reminder

type Console struct {
	Data []Reminder
}

func (i *Console) Set(args SetArgs) error {
	i.Data = append(i.Data, args.Reminder)
	return nil
}

func (i *Console) Insert(args InsertArgs) error {
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

func (i *Console) ListFromChannel(args ChannelListArgs) (*ChannelListRes, error) {
	res := &ChannelListRes{Data: []Reminder{}}
	for _, d := range i.Data {
		if d.GuildId == args.GuildId && d.ChannelId == args.ChannelId {
			res.Data = append(res.Data, d)
		}
	}

	return res, nil
}

func (i *Console) Del(args *DelArgs) error {
	for index, d := range i.Data {
		if d.Id == args.Id && d.GuildId == args.GuildId {
			args.Reminder = d

			if index == 0 {
				i.Data = i.Data[1:]
			} else if index+1 == len(i.Data) {
				i.Data = i.Data[:len(i.Data)-1]
			} else {
				i.Data = append(i.Data[:index], i.Data[index+1:]...)
			}

			return nil
		}
	}

	return ErrDelNotFound
}
