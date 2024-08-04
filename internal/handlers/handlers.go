package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// RegisterHandlers registers all event handlers for the Discord session.
func RegisterHandlers(s *discordgo.Session, prefix string) {
	// Add more handlers here...
	s.AddHandler(ReadyHandler)
	s.AddHandler(MessageHandlerWithPrefix(prefix))
	s.AddHandler(MessageHandlerWithoutPrefix())
}
