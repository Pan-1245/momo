package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageResponses(cmd string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var msg string
	switch cmd {
	case "hi":
		msg = fmt.Sprintf("Hello <@%s>!", m.Author.ID)
	case "fuck you":
		msg = fmt.Sprintf("No, fuck you <@%s>!", m.Author.ID)
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}
