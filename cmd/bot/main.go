package main

import (
	"log"
	"momo/config"
	"momo/internal/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Load env.
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Create a new Discord session.
	sess, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Handle all handlers.
	handlers.RegisterHandlers(sess, cfg.Prefix)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open a the Discord's session.
	err = sess.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer sess.Close()

	// WaitForShutdown waits for a termination signal to gracefully shut down the bot.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
