package commands

import (
	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

type reactionRole struct {
	EmojiID   string
	MessageID string
	RoleID    string
	GuildID   string
}

var reactionRoles []reactionRole

var saySubcommands []types.Subcommand = []types.Subcommand{
	{
		Name: "message",
		Callback: func(parms types.SubcommandParms) {
			handleMessage(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(float64),
				parms.Option.Options[2].Value.(bool),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "dm",
		Callback: func(parms types.SubcommandParms) {
			handleDM(
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
		Name: "reaction-role",
		Callback: func(parms types.SubcommandParms) {
			handleReactionRole(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(string),
				parms.Option.Options[2].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},
}

func reactionHandler(s *dgo.Session, reaction *dgo.MessageReactionAdd) {
	for _, rr := range reactionRoles {
		if rr.EmojiID == reaction.Emoji.Name {
			if reaction.MessageID == rr.MessageID {
				s.GuildMemberRoleAdd(rr.GuildID, reaction.UserID, rr.RoleID)
			}
		}
	}
}

// handleDM handles a say dm command
func handleDM(userID, message string, number float64, embed bool, i *dgo.InteractionCreate, s *dgo.Session) {
	channel, _ := s.UserChannelCreate(userID)
	for r := 0; r < int(number); r++ {
		if embed {
			utils.SendDM(message, channel.ID, s)
		} else {
			s.ChannelMessageSend(channel.ID, message)
		}
	}
}

// handleMessage handles a say message command
func handleMessage(message string, number float64, embed bool, i *dgo.InteractionCreate, s *dgo.Session) {
	for r := 0; r < int(number); r++ {
		if embed {
			utils.SendResponse(message, i, s)
		} else {
			s.ChannelMessageSend(i.ChannelID, message)
		}
	}
}

// handleReactionRole handles a say reaction-role command
func handleReactionRole(message, emoji, roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	m := utils.SendResponse(message, i, s)
	// Add initial reaction
	s.MessageReactionAdd(i.ChannelID, m.ID, emoji)
	// Add reaction role
	reactionRoles = append(reactionRoles, reactionRole{
		EmojiID:   emoji,
		MessageID: m.ID,
		RoleID:    roleID,
		GuildID:   i.GuildID,
	})
}

// RegesterSay adds the say / command and its subcommands
func RegesterSay(client *dgo.Session, guildID string) types.Handler {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "say",
			Description: "More powerful messages",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "message",
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
						{
							Type:        dgo.ApplicationCommandOptionBoolean,
							Name:        "Embed",
							Description: "Sends message in a embed",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "dm",
					Description: "Sends a dm",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionUser,
							Name:        "User",
							Description: "User to dm",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Message",
							Description: "Message to send",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionInteger,
							Name:        "Repeat",
							Description: "Number of times to send",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionBoolean,
							Name:        "Embed",
							Description: "Sends message in a embed",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "reaction-role",
					Description: "Grants role to whoever reacts using given emoji",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "message",
							Description: "Message content",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Emoji",
							Description: "Emoji for reaction",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "Role",
							Description: "Role assigned",
							Required:    true,
						},
					},
				},
			},
		},
		guildID,
	)
	// Add reaction handler
	client.AddHandler(reactionHandler)
	// Return Handler
	return types.Handler{
		Name: "say", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			utils.MatchSubcommand(i, s, saySubcommands)
		},
	}
}
