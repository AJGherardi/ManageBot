package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleInvite handles a invite command
func HandleInvite(i *dgo.InteractionCreate, s *dgo.Session) {
	// Kick user
	invite, _ := s.ChannelInviteCreate(i.ChannelID, dgo.Invite{
		MaxAge:    100,
		MaxUses:   10,
		Temporary: false,
	})
	utils.SendResponse("Invite link https://discord.gg/"+invite.Code, i, s)
}
