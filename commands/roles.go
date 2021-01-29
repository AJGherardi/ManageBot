package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

var roleSubcommands []types.Subcommand = []types.Subcommand{
	{
		Name: "assign",
		Callback: func(parms types.SubcommandParms) {
			handleAssignRole(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "revoke",
		Callback: func(parms types.SubcommandParms) {
			handleRevokeRole(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "create",
		Callback: func(parms types.SubcommandParms) {
			handleCreateRole(
				parms.Option.Options[0].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "delete",
		Callback: func(parms types.SubcommandParms) {
			handleDeleteRole(
				parms.Option.Options[0].Value.(string),
				parms.Interaction,
				parms.Session,
			)

		},
	},
	{
		Name: "view-permissions",
		Callback: func(parms types.SubcommandParms) {
			handleRoleViewPermissions(
				parms.Option.Options[0].Value.(string),
				parms.Interaction,
				parms.Session,
			)
		},
	},

	{
		Name: "general-permissions-set",
		Callback: func(parms types.SubcommandParms) {
			handleRoleGeneralPermissionsSet(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(bool),
				parms.Option.Options[2].Value.(bool),
				parms.Option.Options[3].Value.(bool),
				parms.Option.Options[4].Value.(bool),
				parms.Option.Options[5].Value.(bool),
				parms.Option.Options[6].Value.(bool),
				parms.Option.Options[7].Value.(bool),
				parms.Option.Options[8].Value.(bool),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "membership-permissions-set",
		Callback: func(parms types.SubcommandParms) {
			handleRoleMembershipPermissionsSet(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(bool),
				parms.Option.Options[2].Value.(bool),
				parms.Option.Options[3].Value.(bool),
				parms.Option.Options[4].Value.(bool),
				parms.Option.Options[5].Value.(bool),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "text-permissions-set",
		Callback: func(parms types.SubcommandParms) {
			handleRoleTextPermissionsSet(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(bool),
				parms.Option.Options[2].Value.(bool),
				parms.Option.Options[3].Value.(bool),
				parms.Option.Options[4].Value.(bool),
				parms.Option.Options[5].Value.(bool),
				parms.Option.Options[6].Value.(bool),
				parms.Option.Options[7].Value.(bool),
				parms.Option.Options[8].Value.(bool),
				parms.Option.Options[9].Value.(bool),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "voice-permissions-set",
		Callback: func(parms types.SubcommandParms) {
			handleRoleVoicePermissionsSet(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(bool),
				parms.Option.Options[2].Value.(bool),
				parms.Option.Options[3].Value.(bool),
				parms.Option.Options[4].Value.(bool),
				parms.Option.Options[5].Value.(bool),
				parms.Option.Options[6].Value.(bool),
				parms.Option.Options[7].Value.(bool),
				parms.Interaction,
				parms.Session,
			)
		},
	},
}

func checkPermission(permissions int, permission int) bool {
	permited := (permissions & permission) == permission
	return permited
}

func handleRoleViewPermissions(roleID string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Get permisisons
	permissionViewChannel := checkPermission(role.Permissions, dgo.PermissionViewChannel)
	permissionManageChannels := checkPermission(role.Permissions, dgo.PermissionManageChannels)
	permissionManageRoles := checkPermission(role.Permissions, dgo.PermissionManageRoles)
	permissionManageEmojis := checkPermission(role.Permissions, dgo.PermissionManageEmojis)
	permissionViewAuditLogs := checkPermission(role.Permissions, dgo.PermissionViewAuditLogs)
	permissionManageWebhooks := checkPermission(role.Permissions, dgo.PermissionManageWebhooks)
	permissionManageServer := checkPermission(role.Permissions, dgo.PermissionManageServer)
	permissionAdministrator := checkPermission(role.Permissions, dgo.PermissionAdministrator)
	permissionCreateInstantInvite := checkPermission(role.Permissions, dgo.PermissionCreateInstantInvite)
	permissionChangeNickname := checkPermission(role.Permissions, dgo.PermissionChangeNickname)
	permissionManageNicknames := checkPermission(role.Permissions, dgo.PermissionManageNicknames)
	permissionKickMembers := checkPermission(role.Permissions, dgo.PermissionKickMembers)
	permissionBanMembers := checkPermission(role.Permissions, dgo.PermissionBanMembers)
	permissionSendMessages := checkPermission(role.Permissions, dgo.PermissionSendMessages)
	permissionEmbedLinks := checkPermission(role.Permissions, dgo.PermissionEmbedLinks)
	permissionAttachFiles := checkPermission(role.Permissions, dgo.PermissionAttachFiles)
	permissionAddReactions := checkPermission(role.Permissions, dgo.PermissionAddReactions)
	permissionUseExternalEmojis := checkPermission(role.Permissions, dgo.PermissionUseExternalEmojis)
	permissionMentionEveryone := checkPermission(role.Permissions, dgo.PermissionMentionEveryone)
	permissionManageMessages := checkPermission(role.Permissions, dgo.PermissionManageMessages)
	permissionReadMessageHistory := checkPermission(role.Permissions, dgo.PermissionReadMessageHistory)
	permissionSendTTSMessages := checkPermission(role.Permissions, dgo.PermissionSendTTSMessages)
	permissionVoiceConnect := checkPermission(role.Permissions, dgo.PermissionVoiceConnect)
	permissionVoiceSpeak := checkPermission(role.Permissions, dgo.PermissionVoiceSpeak)
	permissionVoiceUseVAD := checkPermission(role.Permissions, dgo.PermissionVoiceUseVAD)
	permissionVoicePrioritySpeaker := checkPermission(role.Permissions, dgo.PermissionVoicePrioritySpeaker)
	permissionVoiceMuteMembers := checkPermission(role.Permissions, dgo.PermissionVoiceMuteMembers)
	permissionVoiceDeafenMembers := checkPermission(role.Permissions, dgo.PermissionVoiceDeafenMembers)
	permissionVoiceMoveMembers := checkPermission(role.Permissions, dgo.PermissionVoiceMoveMembers)
	// Send perms
	s.ChannelMessageSendEmbed(i.ChannelID, embed.NewGenericEmbed(
		"",
		"Permissions for role "+role.Mention()+
			"\n \n %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s ",
		"permissionViewChannel: "+fmt.Sprint(permissionViewChannel)+"\n",
		"permissionManageChannels: "+fmt.Sprint(permissionManageChannels)+"\n",
		"permissionManageRoles: "+fmt.Sprint(permissionManageRoles)+"\n",
		"permissionManageEmojis: "+fmt.Sprint(permissionManageEmojis)+"\n",
		"permissionViewAuditLogs: "+fmt.Sprint(permissionViewAuditLogs)+"\n",
		"permissionManageWebhooks: "+fmt.Sprint(permissionManageWebhooks)+"\n",
		"permissionManageServer: "+fmt.Sprint(permissionManageServer)+"\n",
		"permissionAdministrator: "+fmt.Sprint(permissionAdministrator)+"\n",
		"permissionCreateInstantInvite: "+fmt.Sprint(permissionCreateInstantInvite)+"\n",
		"permissionChangeNickname: "+fmt.Sprint(permissionChangeNickname)+"\n",
		"permissionManageNicknames: "+fmt.Sprint(permissionManageNicknames)+"\n",
		"permissionKickMembers: "+fmt.Sprint(permissionKickMembers)+"\n",
		"permissionBanMembers: "+fmt.Sprint(permissionBanMembers)+"\n",
		"permissionSendMessages: "+fmt.Sprint(permissionSendMessages)+"\n",
		"permissionEmbedLinks: "+fmt.Sprint(permissionEmbedLinks)+"\n",
		"permissionAttachFiles: "+fmt.Sprint(permissionAttachFiles)+"\n",
		"permissionAddReactions: "+fmt.Sprint(permissionAddReactions)+"\n",
		"permissionUseExternalEmojis: "+fmt.Sprint(permissionUseExternalEmojis)+"\n",
		"permissionMentionEveryone: "+fmt.Sprint(permissionMentionEveryone)+"\n",
		"permissionManageMessages: "+fmt.Sprint(permissionManageMessages)+"\n",
		"permissionReadMessageHistory: "+fmt.Sprint(permissionReadMessageHistory)+"\n",
		"permissionSendTTSMessages: "+fmt.Sprint(permissionSendTTSMessages)+"\n",
		"permissionVoiceConnect: "+fmt.Sprint(permissionVoiceConnect)+"\n",
		"permissionVoiceSpeak: "+fmt.Sprint(permissionVoiceSpeak)+"\n",
		"permissionVoiceUseVAD: "+fmt.Sprint(permissionVoiceUseVAD)+"\n",
		"permissionVoicePrioritySpeaker: "+fmt.Sprint(permissionVoicePrioritySpeaker)+"\n",
		"permissionVoiceMuteMembers: "+fmt.Sprint(permissionVoiceMuteMembers)+"\n",
		"permissionVoiceDeafenMembers: "+fmt.Sprint(permissionVoiceDeafenMembers)+"\n",
		"permissionVoiceMoveMembers: "+fmt.Sprint(permissionVoiceMoveMembers)+"\n",
	))

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

func handleRoleVoicePermissionsSet(
	roleID string,
	connect bool,
	speek bool,
	useVoiceActivity bool,
	prioritySpeeker bool,
	muteMembers bool,
	deafenMembers bool,
	moveMembers bool,
	i *dgo.InteractionCreate,
	s *dgo.Session,
) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Copy perm int
	permissions := role.Permissions
	// Set permissions
	permissions = setPermission(connect, permissions, dgo.PermissionVoiceConnect)
	permissions = setPermission(speek, permissions, dgo.PermissionVoiceSpeak)
	permissions = setPermission(useVoiceActivity, permissions, dgo.PermissionVoiceUseVAD)
	permissions = setPermission(prioritySpeeker, permissions, dgo.PermissionVoicePrioritySpeaker)
	permissions = setPermission(muteMembers, permissions, dgo.PermissionVoiceMuteMembers)
	permissions = setPermission(deafenMembers, permissions, dgo.PermissionVoiceDeafenMembers)
	permissions = setPermission(moveMembers, permissions, dgo.PermissionVoiceMoveMembers)
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

func handleRoleTextPermissionsSet(
	roleID string,
	sendMessages bool,
	embedLinks bool,
	attachFiles bool,
	addReactions bool,
	useExternalEmoji bool,
	mentionAllRoles bool,
	manageMessages bool,
	readMessageHistory bool,
	sendTTSMessages bool,
	i *dgo.InteractionCreate,
	s *dgo.Session,
) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Copy perm int
	permissions := role.Permissions
	// Set permissions
	permissions = setPermission(sendMessages, permissions, dgo.PermissionSendMessages)
	permissions = setPermission(embedLinks, permissions, dgo.PermissionEmbedLinks)
	permissions = setPermission(attachFiles, permissions, dgo.PermissionAttachFiles)
	permissions = setPermission(addReactions, permissions, dgo.PermissionAddReactions)
	permissions = setPermission(useExternalEmoji, permissions, dgo.PermissionUseExternalEmojis)
	permissions = setPermission(mentionAllRoles, permissions, dgo.PermissionMentionEveryone)
	permissions = setPermission(manageMessages, permissions, dgo.PermissionManageMessages)
	permissions = setPermission(readMessageHistory, permissions, dgo.PermissionReadMessageHistory)
	permissions = setPermission(sendTTSMessages, permissions, dgo.PermissionSendTTSMessages)
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

func handleRoleMembershipPermissionsSet(
	roleID string,
	createInvite bool,
	changeNicknames bool,
	manageNicknames bool,
	kickMembers bool,
	banMembers bool,
	i *dgo.InteractionCreate,
	s *dgo.Session,
) {
	// Get role from parms
	role, _ := s.State.Role(i.GuildID, roleID)
	// Copy perm int
	permissions := role.Permissions
	// Set permissions
	permissions = setPermission(createInvite, permissions, dgo.PermissionCreateInstantInvite)
	permissions = setPermission(changeNicknames, permissions, dgo.PermissionChangeNickname)
	permissions = setPermission(manageNicknames, permissions, dgo.PermissionManageNicknames)
	permissions = setPermission(kickMembers, permissions, dgo.PermissionKickMembers)
	permissions = setPermission(banMembers, permissions, dgo.PermissionBanMembers)
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
func RegesterRoles(client *dgo.Session, guildID string) types.Handler {
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
					Name:        "view-permissions",
					Description: "Displays a roles permissions",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionRole,
							Name:        "role",
							Description: "Role to check",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "general-permissions-set",
					Description: "Sets general permissions for a role",
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
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "membership-permissions-set",
					Description: "Sets membership permissions for a role",
					Options: append(
						[]*dgo.ApplicationCommandOption{
							{
								Type:        dgo.ApplicationCommandOptionRole,
								Name:        "role",
								Description: "Role to remove",
								Required:    true,
							},
						},
						membershipPermissionsList()...,
					),
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "text-permissions-set",
					Description: "Sets text permissions for a role",
					Options: append(
						[]*dgo.ApplicationCommandOption{
							{
								Type:        dgo.ApplicationCommandOptionRole,
								Name:        "role",
								Description: "Role to remove",
								Required:    true,
							},
						},
						textPermissionsList()...,
					),
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "voice-permissions-set",
					Description: "Sets voice permissions for a role",
					Options: append(
						[]*dgo.ApplicationCommandOption{
							{
								Type:        dgo.ApplicationCommandOptionRole,
								Name:        "role",
								Description: "Role to remove",
								Required:    true,
							},
						},
						voicePermissionsList()...,
					),
				},
			},
		},
		guildID,
	)
	// Return Handler
	return types.Handler{
		Name: "role", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			utils.MatchSubcommand(i, s, roleSubcommands)
		},
	}
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
