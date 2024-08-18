package commands

import (
	"log"
	"momo/config"
	"momo/internal/integration/google"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ImageResponses(cmd string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// Load env.
	strg, err := config.LoadStorageFromEnv()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the storage client
	storageClient, err := google.NewStorageClient(strg.ServiceAccount)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Oops! Something went wrong with Google Cloud Storage.")
		return
	}
	defer storageClient.Close()

	// Generate a signed URL for the object
	bucketName := strg.Bucket

	// Get object name from param.
	var objectName string
	switch cmd {
	case "yee":
		objectName = "yee"
	case "oh shit":
		objectName = "here_count_dat_boi"
	case "eternal":
		objectName = "eternal"
	case "donald":
		objectName = "donald"
	}

	expiry := 15 * time.Minute
	url, err := storageClient.GenerateSignedURL(bucketName, objectName, expiry)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Oops! Something went wrong while generating URL.")
		return
	}

	// Send the embed with the image
	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
	}

	// Send a message for any error encountered.
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Oops! Something went wrong.")
	}
}
