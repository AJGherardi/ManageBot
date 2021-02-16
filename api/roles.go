package api

import dgo "github.com/bwmarrin/discordgo"

type Role struct {
	guildID string
	roleID  string
	c       *Connection
}

type Permission int

// All permissions bit masks
const (
	PermissionViewChannel          Permission = dgo.PermissionViewChannel
	PermissionManageChannels       Permission = dgo.PermissionManageChannels
	PermissionManageRoles          Permission = dgo.PermissionManageRoles
	PermissionManageEmojis         Permission = dgo.PermissionManageEmojis
	PermissionViewAuditLogs        Permission = dgo.PermissionViewAuditLogs
	PermissionManageWebhooks       Permission = dgo.PermissionManageWebhooks
	PermissionManageServer         Permission = dgo.PermissionManageServer
	PermissionAdministrator        Permission = dgo.PermissionAdministrator
	PermissionCreateInstantInvite  Permission = dgo.PermissionCreateInstantInvite
	PermissionChangeNickname       Permission = dgo.PermissionChangeNickname
	PermissionManageNicknames      Permission = dgo.PermissionManageNicknames
	PermissionKickMembers          Permission = dgo.PermissionKickMembers
	PermissionBanMembers           Permission = dgo.PermissionBanMembers
	PermissionSendMessages         Permission = dgo.PermissionSendMessages
	PermissionEmbedLinks           Permission = dgo.PermissionEmbedLinks
	PermissionAttachFiles          Permission = dgo.PermissionAttachFiles
	PermissionAddReactions         Permission = dgo.PermissionAddReactions
	PermissionUseExternalEmojis    Permission = dgo.PermissionUseExternalEmojis
	PermissionMentionEveryone      Permission = dgo.PermissionMentionEveryone
	PermissionManageMessages       Permission = dgo.PermissionManageMessages
	PermissionReadMessageHistory   Permission = dgo.PermissionReadMessageHistory
	PermissionSendTTSMessages      Permission = dgo.PermissionSendTTSMessages
	PermissionVoiceConnect         Permission = dgo.PermissionVoiceConnect
	PermissionVoiceSpeak           Permission = dgo.PermissionVoiceSpeak
	PermissionVoiceUseVAD          Permission = dgo.PermissionVoiceUseVAD
	PermissionVoicePrioritySpeaker Permission = dgo.PermissionVoicePrioritySpeaker
	PermissionVoiceMuteMembers     Permission = dgo.PermissionVoiceMuteMembers
	PermissionVoiceDeafenMembers   Permission = dgo.PermissionVoiceDeafenMembers
	PermissionVoiceMoveMembers     Permission = dgo.PermissionVoiceMoveMembers
)

func (c *Connection) GetRole(guildID, roleID string) Role {
	return Role{
		guildID: guildID,
		roleID:  roleID,
		c:       c,
	}
}

func (r *Role) GetName() string {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	return role.Name
}

func (r *Role) SetColor(colorDecimal int) {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	r.c.client.GuildRoleEdit(r.guildID, r.roleID, role.Name, colorDecimal, true, role.Permissions, role.Mentionable)
}

func (r *Role) SetName(name string) {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	r.c.client.GuildRoleEdit(r.guildID, r.roleID, name, role.Color, true, role.Permissions, role.Mentionable)
}

func (r *Role) SetMentionable(mentionable bool) {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	r.c.client.GuildRoleEdit(r.guildID, r.roleID, role.Name, role.Color, true, role.Permissions, mentionable)
}

func (r *Role) SetPermission(value bool, permission Permission) {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	permissions := role.Permissions
	if value {
		permissions |= int(permission)
	} else {
		permissions &= ^int(permission)
	}
	r.c.client.GuildRoleEdit(r.guildID, r.roleID, role.Name, role.Color, true, permissions, role.Mentionable)
}

func (r *Role) CheckPermission(permission Permission) bool {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	permited := (role.Permissions & int(permission)) == int(permission)
	return permited
}

func (r *Role) Mention() string {
	role, _ := r.c.client.State.Role(r.guildID, r.roleID)
	return role.Mention()
}
