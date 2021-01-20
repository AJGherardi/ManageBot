package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleSay handles a say command
func HandleSay(message string, i *dgo.InteractionCreate, s *dgo.Session) {
	utils.SendResponse(message, i, s)

}
