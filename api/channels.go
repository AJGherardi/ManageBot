package api

import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

type channel struct {
	channelID string
	c         *Connection
}

func (st *channel) GetName() string {
	channel, _ := st.c.client.Channel(st.channelID)
	return channel.Name
}

func (st *channel) SendMessage(text string) string {
	msg, _ := st.c.client.ChannelMessageSend(st.channelID, text)
	return msg.ID
}

func (st *channel) SendEmbedMessage(text string) string {
	msg, _ := st.c.client.ChannelMessageSendEmbed(st.channelID, embed.NewGenericEmbed("", text))
	return msg.ID
}

func (st *channel) SendEmbedMessageWithTitle(title, caption string) string {
	msg, _ := st.c.client.ChannelMessageSendEmbed(st.channelID, embed.NewGenericEmbed(title, caption))
	return msg.ID
}

func (st *channel) DeleteMessage(msgID string) {
	st.c.client.ChannelMessageDelete(st.channelID, msgID)
}

func (st *channel) DeleteMessages(number int) {
	msgs, _ := st.c.client.ChannelMessages(st.channelID, int(number)+1, "", "", "")
	// Get msg ids
	var msgIDs []string
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID)
	}
	// Delete msgs
	st.c.client.ChannelMessagesBulkDelete(st.channelID, msgIDs)
}

func (st *channel) PinMessage(msgID string) {
	st.c.client.ChannelMessagePin(st.channelID, msgID)
}

func (st *channel) UnpinMessage(msgID string) {
	st.c.client.ChannelMessageUnpin(st.channelID, msgID)
}

func (st *channel) CreateReaction(msgID, emoji string) {
	st.c.client.MessageReactionAdd(st.channelID, msgID, emoji)
}

func (st *channel) GetReactions(msgID, emoji string) []string {
	users, _ := st.c.client.MessageReactions(st.channelID, msgID, emoji, 100, "", "")
	userIDs := []string{}
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}
	return userIDs
}

func (st *channel) Mention() string {
	channel, _ := st.c.client.State.Channel(st.channelID)
	return channel.Mention()
}

type GuildChannel struct {
	channel
}

func (c *Connection) GetChannel(channelID string) GuildChannel {
	return GuildChannel{
		channel: channel{
			channelID: channelID,
			c:         c,
		},
	}
}

func (st *GuildChannel) GetParentID() string {
	channel, _ := st.c.client.Channel(st.channelID)
	return channel.ParentID
}

func (st *GuildChannel) CreateInviteCode(maxUses int, temporary bool) string {
	invite, _ := st.c.client.ChannelInviteCreate(st.channelID, discordgo.Invite{
		MaxAge:    100,
		MaxUses:   maxUses,
		Temporary: temporary,
	})
	return invite.Code
}

// TODO: Implement
// func (st *Channel) PermissionOverrideCreate() {}

// func (st *Channel) PermissionOverrideDelete() {}

type DMChannel struct {
	channel
}

func (c *Connection) GetDMChannel(dmChannelID string) DMChannel {
	return DMChannel{
		channel: channel{
			channelID: dmChannelID,
			c:         c,
		},
	}
}

type Category struct {
	channelID string
	c         *Connection
}

func (c *Connection) GetCategory(channelID string) Category {
	return Category{
		channelID: channelID,
		c:         c,
	}
}

func (st *Category) GetName() string {
	channel, _ := st.c.client.Channel(st.channelID)
	return channel.Name
}

func (st *Category) Mention() string {
	channel, _ := st.c.client.State.Channel(st.channelID)
	return channel.Mention()
}
