package commands

import (
	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleNickname handles a nickname command
func HandleNickname(userID, newNickname string, i *dgo.InteractionCreate, s *dgo.Session) {
	s.GuildMemberNickname(i.GuildID, userID, newNickname)
	utils.SendResponse("Changed Nickname", i, s)
}

// RegesterNickname adds the nickname / command
func RegesterNickname(client *dgo.Session, guildID string) types.Handler {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "nickname",
			Description: "Changes a server members nickname",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionUser,
					Name:        "User",
					Description: "User that will have nickname changed",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "nickname",
					Description: "New Nickname",
					Required:    true,
				},
			},
		},
		guildID,
	)
	// Return Handler
	return types.Handler{
		Name: "nickname", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			HandleNickname(
				i.Interaction.Data.Options[0].Value.(string),
				i.Interaction.Data.Options[1].Value.(string),
				i,
				s,
			)
		},
	}
}
