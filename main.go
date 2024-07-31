package main

import (
	"fmt"
	"log"
	"math/rand"
	"momo/config"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	sess, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world")
		}

		args := strings.Split(m.Content, " ")

		if !strings.HasPrefix(m.Content, cfg.Prefix) {
			return
		} else {
			if args[1] == "hello" {
				s.ChannelMessageSend(m.ChannelID, "world!")

			}

			if args[1] == "proverbs" {
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

				author := discordgo.MessageEmbedAuthor{
					Name: "Rob Pike",
					URL:  "https://go-proverbs.github.io/",
				}
				embed := discordgo.MessageEmbed{
					Title:  proverbs[selection],
					Author: &author,
				}

				s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			}
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
