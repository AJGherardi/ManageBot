package commands

import (
	"github.com/AJGherardi/ManageBot/api"
)

// InviteHandler handles a invite command
type InviteHandler struct{}

func (h *InviteHandler) Name() string {
	return "invite"
}

func (h *InviteHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	channel := c.GetChannel(i.GetChannelID())
	code := channel.CreateInviteCode(10, false)
	channel.SendEmbedMessage("Invite link https://discord.gg/" + code)
}

func (h *InviteHandler) Regester(c api.Connection) api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("invite", "Generate a invite link")
}
