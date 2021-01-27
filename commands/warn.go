package commands

import (
	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleWarn handles a warn command
func HandleWarn(userID, violation string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	utils.SendResponse(user.Mention()+" This is you final warning for "+violation, i, s)
}

// RegesterWarn adds the warn / command
func RegesterWarn(client *dgo.Session, guildID string) types.Handler {
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
	// Return Handler
	return types.Handler{
		Name: "warn", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			HandleWarn(
				i.Interaction.Data.Options[0].Value.(string),
				i.Interaction.Data.Options[1].Value.(string),
				i,
				s,
			)
		},
	}
}
