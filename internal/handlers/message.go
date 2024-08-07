package handlers

import (
	"momo/internal/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageHandler handles message creation events (with prefix).
func MessageHandlerWithPrefix(prefix string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages from the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Ignore messages without the prefix
		content := m.Content
		if !strings.HasPrefix(content, prefix) {
			return
		}

		// Split the content into command and arguments
		args := strings.Fields(content[len(prefix):])
		if len(args) == 0 {
			return
		}

		command := strings.ToLower(args[0])
		executeCommandWithPrefix(s, m, command)
	}
}

// MessageHandler handles message creation events (without prefix).
func MessageHandlerWithoutPrefix() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages from the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		command := strings.ToLower(m.Content)
		executeCommandWithoutPrefix(s, m, command)
	}
}

// executeCommand executes a specific command (with prefix).
func executeCommandWithPrefix(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	switch command {
	// Add more subcommands here...
	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command.")
	}
}

// executeCommand executes a specific command (without prefix).
func executeCommandWithoutPrefix(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	switch command {
	// Add more subcommands here...
	case "hi":
		commands.MessageResponses("hi", s, m)
	case "fuck you", "fuck u", "fk you", "fk u", "f u", "fu":
		commands.MessageResponses("fuck you", s, m)
	case "yee":
		commands.ImageResponses("yee", s, m)
	case "oh shit":
		commands.ImageResponses("oh shit", s, m)
	case "eternal":
		commands.ImageResponses("eternal", s, m)
	case "tris", "tristan":
		commands.ImageResponses("donald", s, m)
	default:
		return
	}
}
