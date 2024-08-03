package commands

import "github.com/bwmarrin/discordgo"

func Yee(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: "https://tenor.com/view/yee-gif-8561646",
		},
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Oops! Something went wrong while sending the GIF.")
	}
}
