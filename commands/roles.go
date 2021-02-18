package commands

import (
	"github.com/AJGherardi/ManageBot/api"
)

type RoleHandler struct{}

func (h *RoleHandler) Name() string {
	return "role"
}

func (h *RoleHandler) Subcommands() []api.Subcommand {
	return []api.Subcommand{
		&createRoleHandler{},
		&deleteRoleHandler{},
		&assignRoleHandler{},
		&revokeRoleHandler{},
		&generalPermissionsSetRoleHandler{},
		&membershipPermissionsSetRoleHandler{},
		&textPermissionsSetRoleHandler{},
		&voicePermissionsSetRoleHandler{},
	}
}

func (h *RoleHandler) Regester(c api.Connection) api.ParentCommandSinginture {
	return api.MakeParentCommandSinginture("role", "Manage user roles")
}

type createRoleHandler struct{}

func (h *createRoleHandler) Name() string {
	return "create"
}

func (h *createRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Make a new role
	guild := c.GetGuild(i.GetGuildID())
	roleID := guild.CreateRole(i.GetStringParm(0), 50, 0, true)
	role := c.GetRole(i.GetGuildID(), roleID)
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Added role " + role.Mention())
}

func (h *createRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"create", "Makes a new role",
		api.MakeStringParmSinginture("Name", "Role name", true),
	)
}

type deleteRoleHandler struct{}

func (h *deleteRoleHandler) Name() string {
	return "delete"
}

func (h *deleteRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get role name
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(0))
	roleName := role.GetName()
	// Delete role
	guild := c.GetGuild(i.GetGuildID())
	guild.DeleteRole(i.GetStringParm(0))
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Deleted role " + roleName)
}

func (h *deleteRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"delete", "Removes a role",
		api.MakeRoleParmSinginture("Role", "Role to remove", true),
	)
}

type assignRoleHandler struct{}

func (h *assignRoleHandler) Name() string {
	return "assign"
}

func (h *assignRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get user and role from parms
	user := c.GetUser(i.GetStringParm(0))
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(1))
	// Assign role to user
	guild := c.GetGuild(i.GetGuildID())
	guild.AssignRole(i.GetStringParm(0), i.GetStringParm(1))
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Added role " + role.Mention() + " to " + user.Mention())
}

func (h *assignRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"assign", "Adds a role to a user",
		api.MakeUserParmSinginture("User", "User to add role to", true),
		api.MakeRoleParmSinginture("Role", "Role to add", true),
	)
}

type revokeRoleHandler struct{}

func (h *revokeRoleHandler) Name() string {
	return "revoke"
}

func (h *revokeRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get user and role from parms
	user := c.GetUser(i.GetStringParm(0))
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(1))
	// Revoked role to user
	guild := c.GetGuild(i.GetGuildID())
	guild.RevokeRole(i.GetStringParm(0), i.GetStringParm(1))
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Removed role " + role.Mention() + " to " + user.Mention())
}

func (h *revokeRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"revoke", "Revokes a current role form a user",
		api.MakeUserParmSinginture("User", "User to remove role from", true),
		api.MakeRoleParmSinginture("Role", "Role to remove", true),
	)
}

type generalPermissionsSetRoleHandler struct{}

func (h *generalPermissionsSetRoleHandler) Name() string {
	return "general-permissions-set"
}

func (h *generalPermissionsSetRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get role from parms
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(0))
	// Set permissions
	role.SetPermission(i.GetBoolParm(1), api.PermissionViewChannel)
	role.SetPermission(i.GetBoolParm(2), api.PermissionManageChannels)
	role.SetPermission(i.GetBoolParm(3), api.PermissionManageRoles)
	role.SetPermission(i.GetBoolParm(4), api.PermissionManageEmojis)
	role.SetPermission(i.GetBoolParm(5), api.PermissionViewAuditLogs)
	role.SetPermission(i.GetBoolParm(6), api.PermissionManageWebhooks)
	role.SetPermission(i.GetBoolParm(7), api.PermissionManageServer)
	role.SetPermission(i.GetBoolParm(8), api.PermissionAdministrator)
	// Commit permissions
	role.CommitPermissions()
}

func (h *generalPermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"general-permissions-set", "Sets general permissions for a role",
		api.MakeRoleParmSinginture("Role", "Role to edit", true),
		api.MakeBoolParmSinginture("ViewChannels", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageChannels", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageRoles", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageEmojis", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ViewAuditLog", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageWebhooks", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageServer", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("Administrator", "Check discord permissions list", true),
	)
}

type membershipPermissionsSetRoleHandler struct{}

func (h *membershipPermissionsSetRoleHandler) Name() string {
	return "membership-permissions-set"
}

func (h *membershipPermissionsSetRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get role from parms
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(0))
	// Set permissions
	role.SetPermission(i.GetBoolParm(1), api.PermissionCreateInstantInvite)
	role.SetPermission(i.GetBoolParm(2), api.PermissionChangeNickname)
	role.SetPermission(i.GetBoolParm(3), api.PermissionManageNicknames)
	role.SetPermission(i.GetBoolParm(4), api.PermissionKickMembers)
	role.SetPermission(i.GetBoolParm(5), api.PermissionBanMembers)
	// Commit permissions
	role.CommitPermissions()
}

func (h *membershipPermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"membership-permissions-set", "Sets membership permissions for a role",
		api.MakeRoleParmSinginture("Role", "Role to edit", true),
		api.MakeBoolParmSinginture("CreateInvite", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ChangeNicknames", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageNicknames", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("KickMembers", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("BanMembers", "Check discord permissions list", true),
	)
}

type textPermissionsSetRoleHandler struct{}

func (h *textPermissionsSetRoleHandler) Name() string {
	return "text-permissions-set"
}

func (h *textPermissionsSetRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get role from parms
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(0))
	// Set permissions
	role.SetPermission(i.GetBoolParm(1), api.PermissionSendMessages)
	role.SetPermission(i.GetBoolParm(2), api.PermissionEmbedLinks)
	role.SetPermission(i.GetBoolParm(3), api.PermissionAttachFiles)
	role.SetPermission(i.GetBoolParm(4), api.PermissionAddReactions)
	role.SetPermission(i.GetBoolParm(5), api.PermissionUseExternalEmojis)
	role.SetPermission(i.GetBoolParm(6), api.PermissionMentionEveryone)
	role.SetPermission(i.GetBoolParm(7), api.PermissionManageMessages)
	role.SetPermission(i.GetBoolParm(8), api.PermissionReadMessageHistory)
	role.SetPermission(i.GetBoolParm(9), api.PermissionSendTTSMessages)
	// Commit permissions
	role.CommitPermissions()
}

func (h *textPermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"text-permissions-set", "Sets text permissions for a role",
		api.MakeRoleParmSinginture("Role", "Role to edit", true),
		api.MakeBoolParmSinginture("SendMessages", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("EmbedLinks", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("AttachFiles", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("AddReactions", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("UseExternalEmoji", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("MentionAllRoles", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ManageMessages", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("ReadMessageHistory", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("SendTTSMessages", "Check discord permissions list", true),
	)
}

type voicePermissionsSetRoleHandler struct{}

func (h *voicePermissionsSetRoleHandler) Name() string {
	return "voice-permissions-set"
}

func (h *voicePermissionsSetRoleHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Get role from parms
	role := c.GetRole(i.GetGuildID(), i.GetStringParm(0))
	// Set permissions
	role.SetPermission(i.GetBoolParm(1), api.PermissionVoiceConnect)
	role.SetPermission(i.GetBoolParm(2), api.PermissionVoiceSpeak)
	role.SetPermission(i.GetBoolParm(3), api.PermissionVoiceUseVAD)
	role.SetPermission(i.GetBoolParm(4), api.PermissionVoicePrioritySpeaker)
	role.SetPermission(i.GetBoolParm(5), api.PermissionVoiceMuteMembers)
	role.SetPermission(i.GetBoolParm(6), api.PermissionVoiceDeafenMembers)
	role.SetPermission(i.GetBoolParm(7), api.PermissionVoiceMoveMembers)
	// Commit permissions
	role.CommitPermissions()
}

func (h *voicePermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"voice-permissions-set", "Sets voice permissions for a role",
		api.MakeRoleParmSinginture("Role", "Role to edit", true),
		api.MakeBoolParmSinginture("Connect", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("Speek", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("UseVoiceActivity", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("PrioritySpeeker", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("MuteMembers", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("DeafenMembers", "Check discord permissions list", true),
		api.MakeBoolParmSinginture("MoveMembers", "Check discord permissions list", true),
	)
}
