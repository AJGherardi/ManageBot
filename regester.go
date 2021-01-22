package main

import (
	"github.com/AJGherardi/ManageBot/commands"
	dgo "github.com/bwmarrin/discordgo"
)

func regesterCommands(client *dgo.Session, guildID string) {
	commands.RegesterRoles(client, guildID)
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
			Name:        "vote",
			Description: "Make a vote",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Title",
					Description: "Title of vote message",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Caption",
					Description: "Caption for vote message",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Time",
					Description: "Time till end of vote in min",
					Required:    true,
				},
			},
		},
		guildID,
	)
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "say",
			Description: "Repeats a message",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Message",
					Description: "Message to repeat",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Repeat",
					Description: "Number of times to repeat",
					Required:    true,
				},
			},
		},
		guildID,
	)
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "invite",
			Description: "Generate a invite link",
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
			Name:        "stats",
			Description: "Shows stats of a server or channel",
		},
		guildID,
	)
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "say",
			Description: "Repeats a message",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Message",
					Description: "Message to repeat",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Repeat",
					Description: "Number of times to repeat",
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
							Name:        "Category",
							Description: "Category to add channel to",
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
