package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleRole handles a top level role command
func HandleRole(i *dgo.InteractionCreate, s *dgo.Session) {
	for _, option := range i.Interaction.Data.Options {
		switch option.Name {
		case "assign":
			handleAssignRole(
				option.Options[0].Value.(string),
				option.Options[1].Value.(string),
				i,
				s,
			)
		case "revoke":
			handleRevokeRole(
				option.Options[0].Value.(string),
				option.Options[1].Value.(string),
				i,
				s,
			)
		case "create":
			handleCreateRole(
				option.Options[0].Value.(string),
				i,
				s,
			)
		case "delete":
			handleDeleteRole(
				option.Options[0].Value.(string),
				i,
				s,
			)
		}
	}
}

func handleAssignRole(userID, roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Assign role to user
	s.GuildMemberRoleAdd(i.GuildID, user.ID, role.ID)
	utils.SendResponse("Added role "+role.Mention()+" to "+user.Mention(), i, s)
}

func handleRevokeRole(userID, roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Removed role from user
	s.GuildMemberRoleRemove(i.GuildID, user.ID, role.ID)
	utils.SendResponse("Revoked role "+role.Mention()+" from "+user.Mention(), i, s)
}

func handleCreateRole(name string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Make a new role
	role, _ := s.GuildRoleCreate(i.GuildID)
	// Set new role info
	s.GuildRoleEdit(i.GuildID, role.ID, name, 50, false, 0, true)
	utils.SendResponse("Added role "+role.Mention(), i, s)
}

func handleDeleteRole(roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	s.GuildRoleDelete(i.GuildID, role.ID)
	utils.SendResponse("Role Removed  "+role.Name, i, s)
}
