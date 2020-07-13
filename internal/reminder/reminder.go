package reminder

import (
	"errors"
	"time"
)

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

type ChannelListArgs struct {
	GuildId   string
	ChannelId string
}

type DelArgs struct {
	Reminder
}

// Response definition for interface
type GuildListRes struct {
	Data []Reminder
}

type ChannelListRes struct {
	Data []Reminder
}

type DelRes struct {
	Data Reminder
}

// Represent connection to data source
type InsertRepo interface {
	Insert(args InsertArgs) error
}

type GuildListRepo interface {
	ListFromGuild(args GuildListArgs) (*GuildListRes, error)
}

type ChannelListRepo interface {
	ListFromChannel(args ChannelListArgs) (*ChannelListRes, error)
}

type DelRepo interface {
	Del(args *DelArgs) error
}

// List of Errors
var (
	ErrNotEnoughArgs = errors.New("invalid input. not enough arguments")
	ErrInvalidId     = errors.New("invalid id format")
	ErrDelNotFound   = errors.New("could not find reminder with that name")
)
