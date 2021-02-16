package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

type Guild struct {
	guildID string
	c       *Connection
}

type Ban struct {
	Reason string
	UserID string
}

func (c *Connection) GetGuild(guildID string) Guild {
	return Guild{
		guildID: guildID,
		c:       c,
	}
}

func (g *Guild) GetChannelIDs() []string {
	channels, _ := g.c.client.GuildChannels(g.guildID)
	channelIDs := []string{}
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.ID)
	}
	return channelIDs
}

func (g *Guild) GetUserIDs() []string {
	users, _ := g.c.client.GuildMembers(g.guildID, "", 1000)
	userIDs := []string{}
	for _, user := range users {
		userIDs = append(userIDs, user.User.ID)
	}
	return userIDs
}

func (g *Guild) GetRoleIDs() []string {
	roles, _ := g.c.client.GuildRoles(g.guildID)
	roleIDs := []string{}
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}
	return roleIDs
}

func (g *Guild) GetEmojiIDs() []string {
	emojis, _ := g.c.client.GuildEmojis(g.guildID)
	emojiIDs := []string{}
	for _, emoji := range emojis {
		emojiIDs = append(emojiIDs, emoji.ID)
	}
	return emojiIDs
}

func (g *Guild) GetIntegrationIDs() []string {
	integrations, _ := g.c.client.GuildIntegrations(g.guildID)
	integrationIDs := []string{}
	for _, integration := range integrations {
		integrationIDs = append(integrationIDs, integration.ID)
	}
	return integrationIDs
}

func (g *Guild) GetAuditLog() {}

func (g *Guild) GetBans() []Ban {
	dgoBans, _ := g.c.client.GuildBans(g.guildID)
	bans := []Ban{}
	for _, dgoBan := range dgoBans {
		bans = append(bans, Ban{Reason: dgoBan.Reason, UserID: dgoBan.User.ID})
	}
	return bans
}

func (g *Guild) BanUser(userID, reason string, purgeDays int) {
	g.c.client.GuildBanCreateWithReason(g.guildID, userID, reason, purgeDays)
}

func (g *Guild) UnbanUser(userID string) {
	g.c.client.GuildBanDelete(g.guildID, userID)
}

func (g *Guild) KickUser(userID, reason string) {
	g.c.client.GuildMemberDeleteWithReason(g.guildID, userID, reason)
}

func (g *Guild) CreateEmoji(name, imageB64 string) string {
	em, _ := g.c.client.GuildEmojiCreate(g.guildID, name, imageB64, nil)
	return em.ID
}

func (g *Guild) DeleteEmoji(emojiID string) {
	g.c.client.GuildEmojiDelete(g.guildID, emojiID)
}

func (g *Guild) CreateCategory(name string) string {
	st, _ := g.c.client.GuildChannelCreateComplex(g.guildID, dgo.GuildChannelCreateData{
		Name: name,
		Type: dgo.ChannelTypeGuildCategory,
	})
	return st.ID
}

func (g *Guild) DeleteCategory(categoryID string) {
	g.c.client.ChannelDelete(categoryID)
}

func (g *Guild) CreateChannel(name, parentID string, channelType int, nsfw bool) string {
	st, _ := g.c.client.GuildChannelCreateComplex(g.guildID, dgo.GuildChannelCreateData{
		Name:     name,
		Type:     dgo.ChannelType(channelType),
		ParentID: parentID,
		NSFW:     nsfw,
	})
	return st.ID
}

func (g *Guild) DeleteChannel(channelID string) {
	g.c.client.ChannelDelete(channelID)
}

func (g *Guild) CreateRole(name string, colorDecimal, permissionInt int, mentionable bool) string {
	role, _ := g.c.client.GuildRoleCreate(g.guildID)
	g.c.client.GuildRoleEdit(g.guildID, role.ID, name, colorDecimal, true, permissionInt, mentionable)
	return role.ID
}

func (g *Guild) DeleteRole(roleID string) {
	g.c.client.GuildRoleDelete(g.guildID, roleID)
}

func (g *Guild) AssignRole(userID, roleID string) {
	g.c.client.GuildMemberRoleAdd(g.guildID, userID, roleID)
}

func (g *Guild) RevokeRole(userID, roleID string) {
	g.c.client.GuildMemberRoleRemove(g.guildID, userID, roleID)
}
