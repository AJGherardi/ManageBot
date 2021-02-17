package main

import (
	"os"

	"github.com/AJGherardi/ManageBot/api"
	"github.com/AJGherardi/ManageBot/commands"
)

var (
	botToken, guildID string
)

func main() {
	// Get env vars
	botToken = os.Getenv("TOKEN")
	guildID = os.Getenv("GUILD_ID")
	// Get connection
	connection := api.ConnectToDiscord(botToken, guildID)
	// Regester and handle commands
	standaloneCommands, parentCommands := getCommands()
	connection.StartCommandHandler(standaloneCommands, parentCommands, guildID)
	// Keep the app runing
	for {
	}
}

// Returns a list of top lever commands standalone and parrent
func getCommands() ([]api.StandaloneCommand, []api.ParentCommand) {
	return []api.StandaloneCommand{
			&commands.InitHandler{},
			&commands.InviteHandler{},
			&commands.KickHandler{},
			&commands.NicknameHandler{},
			&commands.PurgeHandler{},
			&commands.WarnHandler{},
			&commands.StatsHandler{},
		}, []api.ParentCommand{
			&commands.ChannelHandler{},
		}
}
