package commands

import "github.com/bwmarrin/discordgo"

func Hello(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "world!")
}
