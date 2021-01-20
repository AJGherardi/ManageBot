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
		case "general-permissions-set":
			handleRoleGeneralPermissionsSet(
				option.Options[0].Value.(string),
				option.Options[1].Value.(bool),
				option.Options[2].Value.(bool),
				option.Options[3].Value.(bool),
				option.Options[4].Value.(bool),
				option.Options[5].Value.(bool),
				option.Options[6].Value.(bool),
				option.Options[7].Value.(bool),
				option.Options[8].Value.(bool),
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

func handleRoleGeneralPermissionsSet(
	roleID string,
	viewChannels bool,
	manageChannels bool,
	manageRoles bool,
	manageEmojis bool,
	viewAuditLog bool,
	manageWebhooks bool,
	manageServer bool,
	administrator bool,
	i *dgo.InteractionCreate,
	s *dgo.Session,
) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Copy perm int
	permissions := role.Permissions
	// Set permissions
	permissions = setPermission(viewChannels, permissions, dgo.PermissionViewChannel)
	permissions = setPermission(manageChannels, permissions, dgo.PermissionManageChannels)
	permissions = setPermission(manageRoles, permissions, dgo.PermissionManageRoles)
	permissions = setPermission(manageEmojis, permissions, dgo.PermissionManageEmojis)
	permissions = setPermission(viewAuditLog, permissions, dgo.PermissionViewAuditLogs)
	permissions = setPermission(manageWebhooks, permissions, dgo.PermissionManageWebhooks)
	permissions = setPermission(manageServer, permissions, dgo.PermissionManageServer)
	permissions = setPermission(administrator, permissions, dgo.PermissionAdministrator)
	// Save perm int
	s.GuildRoleEdit(
		i.GuildID,
		role.ID,
		role.Name,
		role.Color,
		role.Hoist,
		permissions,
		role.Mentionable,
	)
}

func setPermission(value bool, permissions int, permission int) int {
	if value {
		permissions |= permission
	} else {
		permissions &= ^permission
	}
	return permissions
}

// RegesterRoles adds all role related / commands
func RegesterRoles(client *dgo.Session, guildID string) {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "role",
			Description: "Manage user roles",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "assign",
					Description: "Adds a role to a user",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionUser,
							Name:        "User",
							Description: "User to add role to",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "Role",
							Description: "Role to add",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "revoke",
					Description: "Revokes a current role form a user",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionUser,
							Name:        "User",
							Description: "User to remove role from",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "Role",
							Description: "Role to remove",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "create",
					Description: "Makes a new role",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Name",
							Description: "Role name",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "delete",
					Description: "Removes a role",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "role",
							Description: "Role to remove",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "general-permissions-set",
					Description: "Removes a role",
					Options: append(
						[]*dgo.ApplicationCommandOption{
							{
								Type:        dgo.ApplicationCommandOptionRole,
								Name:        "role",
								Description: "Role to remove",
								Required:    true,
							},
						},
						generalPermissionsList()...,
					),
				},
			},
		},
		guildID,
	)
}

func generalPermissionsList() []*dgo.ApplicationCommandOption {
	return []*dgo.ApplicationCommandOption{
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ViewChannels",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageChannels",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageRoles",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageEmojis",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ViewAuditLog",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageWebhooks",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageServer",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "Administrator",
			Description: "Check discord permissions list",
			Required:    true,
		},
	}
}

func membershipPermissionsList() []*dgo.ApplicationCommandOption {
	return []*dgo.ApplicationCommandOption{
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "CreateInvite",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ChangeNicknames",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageNicknames",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "KickMembers",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "BanMembers",
			Description: "Check discord permissions list",
			Required:    true,
		},
	}
}

func textPermissionsList() []*dgo.ApplicationCommandOption {
	return []*dgo.ApplicationCommandOption{
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "SendMessages",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "EmbedLinks",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "AttachFiles",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "AddReactions",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "UseExternalEmoji",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "MentionAllRoles",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ManageMessages",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "ReadMessageHistory",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "SendTTSMessages",
			Description: "Check discord permissions list",
			Required:    true,
		},
	}
}

func voicePermissionsList() []*dgo.ApplicationCommandOption {
	return []*dgo.ApplicationCommandOption{
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "Connect",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "Speek",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "Video",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "UseVoiceActivity",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "PrioritySpeeker",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "MuteMembers",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "DeafenMembers",
			Description: "Check discord permissions list",
			Required:    true,
		},
		{
			Type:        dgo.ApplicationCommandOptionBoolean,
			Name:        "MoveMembers",
			Description: "Check discord permissions list",
			Required:    true,
		},
	}
}
