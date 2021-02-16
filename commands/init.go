package commands

import (
	"github.com/AJGherardi/ManageBot/api"
	dgo "github.com/bwmarrin/discordgo"
)

type InitHandler struct{}

func (h *InitHandler) Name() string {
	return "init"
}

func (h *InitHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	guild := c.GetGuild(i.GetGuildID())
	// Make channels
	guild.CreateChannel("logs", "", int(dgo.ChannelTypeGuildText), false)
	guild.CreateChannel("reports", "", int(dgo.ChannelTypeGuildText), false)
	guild.CreateCategory("tickets")
	guild.CreateCategory("archives")
	// Make roles
	guild.CreateRole("moderator", 50, 1543499751, true)
	guild.CreateRole("member", 50, 3526209, true)
	guild.CreateRole("muted", 50, 1024, true)
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Server initialized")
}

func (h *InitHandler) Regester() api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("init", "Adds internal channels and roles to server")
}
