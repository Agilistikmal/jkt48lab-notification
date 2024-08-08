package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	Client *discordgo.Session
}

func NewService() *Service {
	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	return &Service{
		Client: client,
	}
}
