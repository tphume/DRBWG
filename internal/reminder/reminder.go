package reminder

import "time"

// Argument definition for interface
type InsertArgs struct {
	Id        string
	GuildId   string
	ChannelId string
	T         time.Time
	Name      string
}

// Represent connection to data source
type InsertRepo interface {
	Insert(args InsertArgs) error
}
