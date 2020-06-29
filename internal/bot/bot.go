package bot

import "github.com/bwmarrin/discordgo"

type Handle func(cmd string) ([]string, error)

type Bot struct {
	Session *discordgo.Session
	routes  map[string]Handle
}
