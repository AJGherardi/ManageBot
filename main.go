package main

import (
	"os"

	"github.com/AJGherardi/ManageBot/api"
)

var (
	botToken, guildID string
)

func main() {
	// Get env vars
	botToken = os.Getenv("TOKEN")
	guildID = os.Getenv("GUILD_ID")
	connection := api.ConnectToDiscord(botToken, guildID)
	connection.StartCommandHandler(getCommands(), guildID)
	// Keep the app runing
	for {
	}
}
