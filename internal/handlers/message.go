package handlers

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageHandler handles message creation events.
func MessageHandler(prefix string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		content := m.Content
		if !strings.HasPrefix(content, prefix) {
			return
		}

		args := strings.Fields(content[len(prefix):])
		if len(args) == 0 {
			return
		}

		command := strings.ToLower(args[0])
		executeCommand(s, m, command)
	}
}

// executeCommand executes a specific command.
func executeCommand(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	switch command {
	// Add more subcommands here...
	case "hello":
		handleHello(s, m)
	case "proverbs":
		handleProverbs(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command.")
	}
}

func handleHello(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "world!")
}

func handleProverbs(s *discordgo.Session, m *discordgo.MessageCreate) {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}

	selection := rand.Intn(len(proverbs))
	author := &discordgo.MessageEmbedAuthor{
		Name: "Rob Pike",
		URL:  "https://go-proverbs.github.io/",
	}
	embed := &discordgo.MessageEmbed{
		Title:  proverbs[selection],
		Author: author,
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
