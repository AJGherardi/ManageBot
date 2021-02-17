package commands

import (
	"github.com/AJGherardi/ManageBot/api"
)

// KickHandler handles a kick command
type KickHandler struct{}

func (h *KickHandler) Name() string {
	return "kick"
}

func (h *KickHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	guild := c.GetGuild(i.GetGuildID())
	guild.KickUser(i.GetStringParm(0), "None specified")
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("User kicked")
}

func (h *KickHandler) Regester() api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("kick", "Kicks a user",
		api.MakeUserParmSinginture("User", "User to kick", true),
	)
}
