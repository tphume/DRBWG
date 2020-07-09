package reminder

import "time"

type Reminder struct {
	Id        string
	GuildId   string
	ChannelId string
	T         time.Time
	Name      string
}

// Argument definition for interface
type InsertArgs struct {
	Reminder
}

type GuildListArgs struct {
	GuildId string
}

// Response definition for interface
type GuildListRes struct {
	Data []Reminder
}

// Represent connection to data source
type InsertRepo interface {
	Insert(args InsertArgs) error
}

type GuildListRepo interface {
	GuildList(args GuildListArgs) (*GuildListRes, error)
}
