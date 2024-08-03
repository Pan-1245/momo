package commands

import "github.com/bwmarrin/discordgo"

func Hi(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "hi")
}
