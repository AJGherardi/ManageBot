package main

import (
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

// botToken and guildID must be added to consts.go

func main() {
	// Creates a new client object
	client, _ := dgo.New("Bot " + botToken)
	// Regesters a event handeler for when the command is called
	client.AddHandler(commandHandler(client))
	// Opens the connection
	client.Open()
	// Remove all commands
	deleteAllCommands(client)
	// Regesters the commands
	regesterCommands(client)
	// Keep the app runing
	for {
	}
}

func commandHandler(client *dgo.Session) func(s *dgo.Session, i *dgo.InteractionCreate) {
	return func(s *dgo.Session, i *dgo.InteractionCreate) {
		// Makes a reaponse
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: "Pls wait",
		}
		// Sends the inital response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
		// Wait a half sec
		time.Sleep(500 * time.Millisecond)
		// Match command to handler function
		switch i.Interaction.Data.Name {
		case "warn":
			handleWarn(
				i.Interaction.Data.Options[0].Value.(string),
				i.Interaction.Data.Options[1].Value.(string),
				i,
				s,
			)
		case "role":
			handleRole(
				i,
				s,
			)
		}
	}
}

func handleRole(i *dgo.InteractionCreate, s *dgo.Session) {
	for _, option := range i.Interaction.Data.Options {
		switch option.Name {
		case "assign":
			handleAssignRole(
				option.Options[0].Value.(string),
				option.Options[1].Value.(string),
				i,
				s,
			)
		case "revoke":
			handleRevokeRole(
				option.Options[0].Value.(string),
				option.Options[1].Value.(string),
				i,
				s,
			)
		case "create":
			handleCreateRole(
				option.Options[0].Value.(string),
				i,
				s,
			)
		case "delete":
			handleDeleteRole(
				option.Options[0].Value.(string),
				i,
				s,
			)
		}
	}
}

func handleAssignRole(userID, roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Assign role to user
	s.GuildMemberRoleAdd(i.GuildID, user.ID, role.ID)
	s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
		Content: "Added role " + role.Mention() + " to " + user.Mention(),
	})
}

func handleRevokeRole(userID, roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Removed role from user
	s.GuildMemberRoleRemove(i.GuildID, user.ID, role.ID)
	s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
		Content: "Revoked role " + role.Mention() + " from " + user.Mention(),
	})
}

func handleCreateRole(name string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Make a new role
	role, _ := s.GuildRoleCreate(i.GuildID)
	// Set new role info
	s.GuildRoleEdit(i.GuildID, role.ID, name, 10, false, 0, true)
	s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
		Content: "Added role " + role.Mention(),
	})
}

func handleDeleteRole(roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	s.GuildRoleDelete(i.GuildID, role.ID)
	s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
		Content: "Role Removed  " + role.Name,
	})
}

func handleWarn(userID, violation string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
		Content: user.Mention() + " This is you final warning for " + violation,
	})
}

func regesterCommands(client *dgo.Session) {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "warn",
			Description: "Warn for user rule violation",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionUser,
					Name:        "User",
					Description: "User to warn",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Violation",
					Description: "Rules violated",
					Choices: []*dgo.ApplicationCommandOptionChoice{
						{
							Name:  "Gore",
							Value: "Gore",
						},
						{
							Name:  "Harassment",
							Value: "Harassment",
						},
						{
							Name:  "Disrespecting staff",
							Value: "Disrespecting staff",
						},
						{
							Name:  "Sexually explicit content",
							Value: "Sexually explicit content",
						},
						{
							Name:  "Advertizing",
							Value: "Advertizing",
						},
						{
							Name:  "Spam",
							Value: "Spam",
						},
						{
							Name:  "Obsessive pinging",
							Value: "Obsessive Pinging",
						},
						{
							Name:  "Hate Speech",
							Value: "Hate Speech",
						},
						{
							Name:  "Threatening People",
							Value: "Threatening People",
						},
						{
							Name:  "Sending Dangerous Links",
							Value: "Sending Dangerous Links",
						},
					},
					Required: true,
				},
			},
		},
		guildID,
	)
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "role",
			Description: "Manage user roles",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "assign",
					Description: "Adds a role to a user",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionUser,
							Name:        "User",
							Description: "User to add role to",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "Role",
							Description: "Role to add",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "revoke",
					Description: "Revokes a current role form a user",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionUser,
							Name:        "User",
							Description: "User to remove role from",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "Role",
							Description: "Role to remove",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "create",
					Description: "Makes a new role",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Name",
							Description: "Role name",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "delete",
					Description: "Removes a role",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "role",
							Description: "Role to remove",
							Required:    true,
						},
					},
				},
			},
		},
		guildID,
	)
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "kick",
			Description: "Kicks a user",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionUser,
					Name:        "User",
					Description: "User to kick",
					Required:    true,
				},
			},
		},
		guildID,
	)
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
