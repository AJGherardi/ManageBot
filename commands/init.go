package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleInit handles a init command
func HandleInit(i *dgo.InteractionCreate, s *dgo.Session) {
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
		Type: dgo.ChannelType(dgo.ChannelTypeGuildText),
		NSFW: false,
	})
	// Make roles
	moderator, _ := s.GuildRoleCreate(i.GuildID)
	s.GuildRoleEdit(i.GuildID, moderator.ID, "moderator", 50, false, 1543499751, true)
	member, _ := s.GuildRoleCreate(i.GuildID)
	s.GuildRoleEdit(i.GuildID, member.ID, "member", 50, false, 3526209, true)
	// Inform admin
	utils.SendResponse("Server initialized", i, s)
}

// RegesterInit adds the init / command
func RegesterInit(client *dgo.Session, guildID string) {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "init",
			Description: "Adds internal channels and roles to server",
		},
		guildID,
	)
}
