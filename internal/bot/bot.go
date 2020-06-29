package bot

import "github.com/bwmarrin/discordgo"

type Handle func(cmd string) ([]string, error)

type Bot struct {
	MsgCreateRoutes map[string]Handle
}

func (b *Bot) HandleMsgCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	panic("implement me")
}
