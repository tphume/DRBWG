package reminder

import (
	"errors"
	"github.com/bwmarrin/discordgo"
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
type SetArgs struct {
	Reminder
}

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

type UpdateArgs struct {
	Reminder
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

// Represent connection to data source
type SetRepo interface {
	Set(args SetArgs) error
}

type InsertRepo interface {
	Insert(args InsertArgs) error
}

type GuildListRepo interface {
	ListFromGuild(args GuildListArgs) (*GuildListRes, error)
}

type ChannelListRepo interface {
	ListFromChannel(args ChannelListArgs) (*ChannelListRes, error)
}

type UpdateRepo interface {
	Update(args UpdateArgs) error
}

type DelRepo interface {
	Del(args *DelArgs) error
}

// List of Errors
var (
	ErrNotEnoughArgs = errors.New("invalid input. not enough arguments")
	ErrBadArgs       = errors.New("bad argument for command")
	ErrInvalidId     = errors.New("invalid id format")
	ErrDelNotFound   = errors.New("could not find reminder with that name")
)

// List of things for embed messages
var (
	Color = 15224923
	URL   = "https://github.com/tphume/DRBWG"

	Footer = &discordgo.MessageEmbedFooter{
		Text:    "Made by tphume",
		IconURL: "https://cdn.discordapp.com/embed/avatars/4.png",
	}

	Author = &discordgo.MessageEmbedAuthor{
		URL:     "https://github.com/tPhume",
		Name:    "tphume",
		IconURL: "https://avatars1.githubusercontent.com/u/41682682?s=460&u=ef6f7c71c9bfd5ae6c8de6299a96cc0075e34767&v=4",
	}
)
