package listener

import (
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/agilistikmal/jkt48lab-notification/internal/app/discord"
	"github.com/agilistikmal/jkt48lab-notification/internal/app/idnlive"
	"github.com/agilistikmal/jkt48lab-notification/internal/app/live"
	"github.com/agilistikmal/jkt48lab-notification/internal/app/showroom"
	"github.com/bwmarrin/discordgo"
)

type Service struct {
	IDNLiveService  *idnlive.Service
	ShowroomService *showroom.Service
	DiscordService  *discord.Service
}

func NewService(idnLiveService *idnlive.Service, showroomService *showroom.Service, discordService *discord.Service) *Service {
	return &Service{
		IDNLiveService:  idnLiveService,
		ShowroomService: showroomService,
		DiscordService:  discordService,
	}
}

var onLiveUsernames []string
var notifiedLiveUsernames []string

func (s *Service) Listen() {
	go func() {
		for {
			var lives []*live.Live

			// IDN Live
			idnLives, err := s.IDNLiveService.GetLives()
			if err != nil {
				log.Fatal(err)
			}
			lives = append(lives, idnLives...)

			// Showroom
			showroomLives, err := s.ShowroomService.GetLives()
			if err != nil {
				log.Fatal(err)
			}
			lives = append(lives, showroomLives...)

			err = s.FilterAndSendNotification(lives)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(10 * time.Second)
		}
	}()
}

func (s *Service) FilterAndSendNotification(lives []*live.Live) error {
	var liveUsernames []string

	for _, live := range lives {
		liveUsernames = append(liveUsernames, live.Member.Username)

		containInOnLive := slices.Contains(onLiveUsernames, live.Member.Username)
		if !containInOnLive {
			// onLives = append(onLives[:i], onLives[i+1:]...)
			onLiveUsernames = append(onLiveUsernames, live.Member.Username)
		}

		containInNotified := slices.Contains(notifiedLiveUsernames, live.Member.Username)
		if !containInNotified {
			err := s.SendChannelNotification(live)
			if err != nil {
				log.Printf("failed to send channel notification %v", live.Member.Username)
				log.Println(err)
			}
			notifiedLiveUsernames = append(notifiedLiveUsernames, live.Member.Username)
		}
	}

	for i, onLiveUsername := range onLiveUsernames {
		containInLiveUsernames := slices.Contains(liveUsernames, onLiveUsername)
		if !containInLiveUsernames {
			log.Println("Removed onLive:", onLiveUsername)
			onLiveUsernames = append(onLiveUsernames[:i], onLiveUsernames[i+1:]...)
			notifiedLiveUsernames = append(notifiedLiveUsernames[:i], notifiedLiveUsernames[i+1:]...)
		}
	}

	for i, notifiedLiveUsername := range notifiedLiveUsernames {
		containInLiveUsernames := slices.Contains(liveUsernames, notifiedLiveUsername)
		if !containInLiveUsernames {
			log.Println("Removed notified:", notifiedLiveUsername)
			notifiedLiveUsernames = append(notifiedLiveUsernames[:i], notifiedLiveUsernames[i+1:]...)
		}
	}

	return nil
}

func (s *Service) SendChannelNotification(live *live.Live) error {
	embed := &discordgo.MessageEmbed{
		Title:       "Notifikasi Live",
		Description: fmt.Sprintf("**%s** sedang live di **[%s](%s)**", live.Member.Name, live.Platform, live.OriginalUrl),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "JKT48Lab",
				Value:  "Soon",
				Inline: true,
			},
			{
				Name:   fmt.Sprintf("Tonton di %s", live.Platform),
				Value:  fmt.Sprintf("[klik disini](%s)", live.OriginalUrl),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "JKT48Lab",
		},
		Timestamp: live.StartedAt.Format(time.RFC3339),
		Color:     0xccec1c,
		Image: &discordgo.MessageEmbedImage{
			URL: live.ImageUrl,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: live.Member.ImageUrl,
		},
	}
	_, err := s.DiscordService.Client.ChannelMessageSendEmbed(os.Getenv("DISCORD_NOTIFICATION_CHANNEL_ID"), embed)
	return err
}
