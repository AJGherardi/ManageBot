package commands

import (
	"github.com/AJGherardi/ManageBot/api"
)

// NicknameHandler handles a nickname command
type NicknameHandler struct{}

func (h *NicknameHandler) Name() string {
	return "nickname"
}

func (h *NicknameHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	guild := c.GetGuild(i.GetGuildID())
	guild.SetNickname(i.GetStringParm(0), i.GetStringParm(1))
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Changed Nickname")
}

func (h *NicknameHandler) Regester() api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("nickname", "Changes a server members nickname",
		api.MakeUserParmSinginture("User", "User that will have nickname changed", true),
		api.MakeStringParmSinginture("Nickname", "New Nickname", true),
	)
}
