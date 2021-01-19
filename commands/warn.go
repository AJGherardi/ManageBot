package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleWarn handles a warn command
func HandleWarn(userID, violation string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	utils.SendResponse(user.Mention()+" This is you final warning for "+violation, i, s)
}
