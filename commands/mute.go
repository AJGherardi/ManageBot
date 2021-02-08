package commands

import (
	"fmt"
	"time"

	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleMute handles a mute command
func HandleMute(userID string, roleID string, duration float64, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Assign role to user
	s.GuildMemberRoleAdd(i.GuildID, user.ID, role.ID)
	utils.SendResponse("Muted "+user.Mention()+" for "+fmt.Sprint(duration)+" min", i, s)
	// Add a timer for the mute
	time.AfterFunc(
		(time.Duration(duration) * time.Minute), func() {
			// Removes mute from user
			s.GuildMemberRoleRemove(i.GuildID, user.ID, role.ID)
		},
	

}

// RegesterMute adds the /mute command
func RegesterMute(client *dgo.Session, guildID string) types.Handler {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "mute",
			Description: "Mutes a user",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionUser,
					Name:        "User",
					Description: "User that will be muted",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Time",
					Description: "Amount of time to mute user",
					Required:    true,
				},
			},
		},
		guildID,
	)
	// Return Handler
	return types.Handler{
		Name: "mute", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			HandleMute(
				i.Interaction.Data.Options[0].Value.(string),
				i.Interaction.Data.Options[2].Value.(string),
				i.Interaction.Data.Options[3].Value.(float64),
				i,
				s,
			)
		},
	}
}
