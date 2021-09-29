package main

import (
	"fmt"
	"os"

	"github.com/AJGherardi/ManageBot/api"
	"github.com/AJGherardi/ManageBot/commands"
)

var (
	botToken, guildID string
)

func main() {
	fmt.Println("hello")
	// Get env vars
	botToken = os.Getenv("TOKEN")
	guildID = os.Getenv("GUILD_ID")
	// Get connection
	connection := api.ConnectToDiscord(botToken, guildID)
	// Register and handle commands
	standaloneCommands, parentCommands := getCommands()
	connection.StartCommandHandler(standaloneCommands, parentCommands, guildID)
	fmt.Println("running")
	// Keep the app running
	for {
	}
}

// Returns a list of top lever commands standalone and parent
func getCommands() ([]api.StandaloneCommand, []api.ParentCommand) {
	return []api.StandaloneCommand{
			&commands.InitHandler{},
			&commands.InviteHandler{},
			&commands.KickHandler{},
			&commands.NicknameHandler{},
			&commands.PurgeHandler{},
			&commands.WarnHandler{},
			&commands.StatsHandler{},
			&commands.VoteHandler{},
		}, []api.ParentCommand{
			&commands.ChannelHandler{},
			&commands.SayHandler{},
			&commands.RemindHandler{},
			&commands.RoleHandler{},
		}
}
