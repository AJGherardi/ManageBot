package main

import dgo "github.com/bwmarrin/discordgo"

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
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "purge",
			Description: "Removes specified number of msgs from current channel",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Number",
					Description: "Number of messages to remove",
					Required:    true,
				},
			},
		},
		guildID,
	)
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "channel",
			Description: "Manage channels",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "create",
					Description: "Adds a channel",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Name",
							Description: "Name to give new channel",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionChannel,
							Name:        "Group",
							Description: "Group to add channel to",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionInteger,
							Name:        "Type",
							Description: "Type of new channel",
							Choices: []*dgo.ApplicationCommandOptionChoice{
								{Name: "Text", Value: dgo.ChannelTypeGuildText},
								{Name: "Voice", Value: dgo.ChannelTypeGuildVoice},
							},
							Required: true,
						},
						{
							Type:        dgo.ApplicationCommandOptionBoolean,
							Name:        "NSFW",
							Description: "Contains explicit material only applys to text channels",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "create-group",
					Description: "Adds a channel group",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Name",
							Description: "Name to give new channel groupo",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "delete",
					Description: "Remove a channel",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionChannel,
							Name:        "Channel",
							Description: "Channel to remove",
							Required:    true,
						},
					},
				},
			},
		},
		guildID,
	)
}
