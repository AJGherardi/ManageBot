package api

import (
	"time"

	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// Connection Wraps the discord api for easy use
type Connection struct {
	client *dgo.Session
}

// ConnectToDiscord opens a new connection to discord
func ConnectToDiscord(botToken, guildID string) Connection {
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
	deleteAllCommands(client, guildID)
	return Connection{client: client}
}

// StartCommandHandler Regesters all commands and begines command routing
func (c *Connection) StartCommandHandler(commands []types.StandaloneCommand, guildID string) {
	// Regester all commands
	for _, command := range commands {
		command.Regester(c.client, guildID)
	}
	handle := func(s *dgo.Session, i *dgo.InteractionCreate) {
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
		// Route to appliction command handler
		for _, handler := range commands {
			if handler.Name() == i.Interaction.Data.Name {
				handler.Callback(i, s)
			}
		}
	}
	// Regester handler
	c.client.AddHandler(handle)
}

func deleteAllCommands(client *dgo.Session, guildID string) {
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
