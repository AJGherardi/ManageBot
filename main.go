package main

import (
	"time"

	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// botToken and guildID must be added to consts.go

func main() {
	// Creates a new client object
	client, _ := dgo.New("Bot " + botToken)
	// Set intents
	client.Identify.Intents = dgo.MakeIntent(
		dgo.IntentsAllWithoutPrivileged |
			dgo.IntentsGuildPresences |
			dgo.IntentsGuildMembers,
	)
	// Opens the connection
	client.Open()
	// Remove all commands
	deleteAllCommands(client)
	// Regesters the commands
	handlers := regesterCommands(client, guildID)
	// Regesters a event handeler for when the command is called
	client.AddHandler(commandHandler(client, handlers))
	// Keep the app runing
	for {
	}
}

func commandHandler(client *dgo.Session, handlers []types.Handler) func(s *dgo.Session, i *dgo.InteractionCreate) {
	return func(s *dgo.Session, i *dgo.InteractionCreate) {
		// Makes a reaponse
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: "Please wait",
		}
		// Sends the inital response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
		// Wait a half sec
		time.Sleep(500 * time.Millisecond)
		// Chack perms
		var authorized bool
		for _, roleID := range i.Interaction.Member.Roles {
			role, _ := s.State.Role(i.GuildID, roleID)
			permited := (role.Permissions & dgo.PermissionAdministrator) == dgo.PermissionAdministrator
			if permited {
				authorized = true
				break
			}
		}
		// Remove initial reaponse
		s.InteractionResponseDelete("", i.Interaction)
		// Check if authorized
		if authorized == false {
			utils.SendResponse("Not authorized", i, s)
			return
		}
		// Match command to handler function
		for _, handler := range handlers {
			if handler.Name == i.Interaction.Data.Name {
				handler.Callback(i, s)
			}
		}
	}
}

func deleteAllCommands(client *dgo.Session) {
	apps, _ := client.Applications()
	for _, app := range apps {
		cmds, _ := client.ApplicationCommands(app.ID, guildID)
		for _, cmd := range cmds {
			client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
		}
	}
	cmds, _ := client.ApplicationCommands("", guildID)
	for _, cmd := range cmds {
		client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
	}
}
