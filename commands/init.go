package commands

import (
	"github.com/AJGherardi/ManageBot/api"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

type InitHandler struct{}

func (h *InitHandler) Name() string {
	return "init"
}

func (h *InitHandler) Callback(i *dgo.InteractionCreate, s *dgo.Session) {
	// Make channels
	s.GuildChannelCreateComplex(i.GuildID, dgo.GuildChannelCreateData{
		Name: "logs",
		Type: dgo.ChannelType(dgo.ChannelTypeGuildText),
		NSFW: false,
	})
	s.GuildChannelCreateComplex(i.GuildID, dgo.GuildChannelCreateData{
		Name: "reports",
		Type: dgo.ChannelType(dgo.ChannelTypeGuildText),
		NSFW: false,
	})
	s.GuildChannelCreateComplex(i.GuildID, dgo.GuildChannelCreateData{
		Name: "tickets",
		Type: dgo.ChannelTypeGuildCategory,
	})
	s.GuildChannelCreateComplex(i.GuildID, dgo.GuildChannelCreateData{
		Name: "archives",
		Type: dgo.ChannelTypeGuildCategory,
	})
	// Make roles
	moderator, _ := s.GuildRoleCreate(i.GuildID)
	s.GuildRoleEdit(i.GuildID, moderator.ID, "moderator", 50, false, 1543499751, true)
	member, _ := s.GuildRoleCreate(i.GuildID)
	s.GuildRoleEdit(i.GuildID, member.ID, "member", 50, false, 3526209, true)
	muted, _ := s.GuildRoleCreate(i.GuildID)
	s.GuildRoleEdit(i.GuildID, muted.ID, "muted", 50, false, 1024, true)
	// Inform admin
	utils.SendResponse("Server initialized", i, s)
}

func (h *InitHandler) Regester() api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("init", "Adds internal channels and roles to server")
}
