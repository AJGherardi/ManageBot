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

func (h *RoleHandler) Regester(c api.Connection) api.ParentCommandSignature {
	return api.MakeParentCommandSignature("role", "Manage user roles")
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

func (h *createRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"create", "Makes a new role",
		api.MakeStringParmSignature("Name", "Role name", true),
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

func (h *deleteRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"delete", "Removes a role",
		api.MakeRoleParmSignature("Role", "Role to remove", true),
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

func (h *assignRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"assign", "Adds a role to a user",
		api.MakeUserParmSignature("User", "User to add role to", true),
		api.MakeRoleParmSignature("Role", "Role to add", true),
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

func (h *revokeRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"revoke", "Revokes a current role form a user",
		api.MakeUserParmSignature("User", "User to remove role from", true),
		api.MakeRoleParmSignature("Role", "Role to remove", true),
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

func (h *generalPermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"general-permissions-set", "Sets general permissions for a role",
		api.MakeRoleParmSignature("Role", "Role to edit", true),
		api.MakeBoolParmSignature("ViewChannels", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageChannels", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageRoles", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageEmojis", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ViewAuditLog", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageWebhooks", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageServer", "Check discord permissions list", true),
		api.MakeBoolParmSignature("Administrator", "Check discord permissions list", true),
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

func (h *membershipPermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"membership-permissions-set", "Sets membership permissions for a role",
		api.MakeRoleParmSignature("Role", "Role to edit", true),
		api.MakeBoolParmSignature("CreateInvite", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ChangeNicknames", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageNicknames", "Check discord permissions list", true),
		api.MakeBoolParmSignature("KickMembers", "Check discord permissions list", true),
		api.MakeBoolParmSignature("BanMembers", "Check discord permissions list", true),
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

func (h *textPermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"text-permissions-set", "Sets text permissions for a role",
		api.MakeRoleParmSignature("Role", "Role to edit", true),
		api.MakeBoolParmSignature("SendMessages", "Check discord permissions list", true),
		api.MakeBoolParmSignature("EmbedLinks", "Check discord permissions list", true),
		api.MakeBoolParmSignature("AttachFiles", "Check discord permissions list", true),
		api.MakeBoolParmSignature("AddReactions", "Check discord permissions list", true),
		api.MakeBoolParmSignature("UseExternalEmoji", "Check discord permissions list", true),
		api.MakeBoolParmSignature("MentionAllRoles", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ManageMessages", "Check discord permissions list", true),
		api.MakeBoolParmSignature("ReadMessageHistory", "Check discord permissions list", true),
		api.MakeBoolParmSignature("SendTTSMessages", "Check discord permissions list", true),
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

func (h *voicePermissionsSetRoleHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"voice-permissions-set", "Sets voice permissions for a role",
		api.MakeRoleParmSignature("Role", "Role to edit", true),
		api.MakeBoolParmSignature("Connect", "Check discord permissions list", true),
		api.MakeBoolParmSignature("Speak", "Check discord permissions list", true),
		api.MakeBoolParmSignature("UseVoiceActivity", "Check discord permissions list", true),
		api.MakeBoolParmSignature("PrioritySpeaker", "Check discord permissions list", true),
		api.MakeBoolParmSignature("MuteMembers", "Check discord permissions list", true),
		api.MakeBoolParmSignature("DeafenMembers", "Check discord permissions list", true),
		api.MakeBoolParmSignature("MoveMembers", "Check discord permissions list", true),
	)
}
