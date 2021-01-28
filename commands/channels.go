package commands

import (
	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

var channelSubcommands []types.Subcommand = []types.Subcommand{
	{
		Name: "create",
		Callback: func(parms types.SubcommandParms) {
			handleCreateChannel(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(string),
				parms.Option.Options[2].Value.(float64),
				parms.Option.Options[3].Value.(bool),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "delete",
		Callback: func(parms types.SubcommandParms) {
			handleDeleteChannel(
				parms.Option.Options[0].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "create-group",
		Callback: func(parms types.SubcommandParms) {
			handleCreateChannelGroup(
				parms.Option.Options[0].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},
}

func handleDeleteChannel(channelID string, i *dgo.InteractionCreate, s *dgo.Session) {
	channel, _ := s.Channel(channelID)
	s.ChannelDelete(channelID)
	utils.SendResponse("Deleted channel "+channel.Name, i, s)
}

func handleCreateChannel(name, parentID string, channelType float64, NSFW bool, i *dgo.InteractionCreate, s *dgo.Session) {
	channel, _ := s.GuildChannelCreateComplex(i.GuildID, dgo.GuildChannelCreateData{
		Name:     name,
		Type:     dgo.ChannelType(channelType),
		ParentID: parentID,
		NSFW:     NSFW,
	})
	utils.SendResponse("Added channel "+channel.Mention(), i, s)
}

func handleCreateChannelGroup(name string, i *dgo.InteractionCreate, s *dgo.Session) {
	channel, _ := s.GuildChannelCreateComplex(i.GuildID, dgo.GuildChannelCreateData{
		Name: name,
		Type: dgo.ChannelTypeGuildCategory,
	})
	utils.SendResponse("Added channel group "+channel.Mention(), i, s)
}

// RegesterChannel adds the channel / commands
func RegesterChannel(client *dgo.Session, guildID string) types.Handler {
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
	// Return Handler
	return types.Handler{
		Name: "channel", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			utils.MatchSubcommand(i, s, channelSubcommands)
		},
	}
}
