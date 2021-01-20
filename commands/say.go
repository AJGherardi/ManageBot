package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleSay handles a say command
func HandleSay(message string, number float64, i *dgo.InteractionCreate, s *dgo.Session) {

	for r := 0; r < int(number); r++ {
		utils.SendResponse(message, i, s)
	}

}
