package commands

import (
	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleKick handles a kick command
func HandleKick(userID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	// Kick user
	s.GuildMemberDelete(i.GuildID, user.ID)
	utils.SendResponse("Kicked "+user.Username, i, s)
}

// RegesterKick adds the kick / command
func RegesterKick(client *dgo.Session, guildID string) types.Handler {
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
	// Return Handler
	return types.Handler{
		Name: "kick", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			HandleKick(
				i.Interaction.Data.Options[0].Value.(string),
				i,
				s,
			)
		},
	}
}
