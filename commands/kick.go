package commands

import (
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
