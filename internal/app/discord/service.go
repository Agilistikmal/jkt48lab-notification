package discord

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	Client *discordgo.Session
}

func NewService() *Service {
	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	client.Identify.Intents = discordgo.IntentsGuildMessages

	err = client.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	client.UpdateListeningStatus("JKT48")

	return &Service{
		Client: client,
	}
}
