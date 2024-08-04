package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// ReadyHandler handles the bot's state and status.
func ReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Bot is ready as %s#%s", s.State.User.Username, s.State.User.Discriminator)

	// Set the bot's status or activity.
	err := s.UpdateCustomStatus("Peaching it up! 🍑")
	if err != nil {
		log.Printf("Error setting status: %v", err)
	}
}
