package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/agilistikmal/jkt48lab-notification/internal/app/discord"
	"github.com/agilistikmal/jkt48lab-notification/internal/app/idnlive"
	"github.com/agilistikmal/jkt48lab-notification/internal/app/listener"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	idnLiveService := idnlive.NewService()
	discordService := discord.NewService()
	listenerService := listener.NewService(idnLiveService, discordService)

	listenerService.Listen()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discordService.Client.Close()
}
