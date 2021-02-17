package api

// TODO: Deduplicate
import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

type Channel struct {
	channelID string
	c         *Connection
}

func (c *Connection) GetChannel(channelID string) Channel {
	return Channel{
		channelID: channelID,
		c:         c,
	}
}

func (st *Channel) GetName() string {
	channel, _ := st.c.client.Channel(st.channelID)
	return channel.Name
}

func (st *Channel) GetParentID() string {
	channel, _ := st.c.client.Channel(st.channelID)
	return channel.ParentID
}

func (st *Channel) SendMessage(text string) string {
	msg, _ := st.c.client.ChannelMessageSend(st.channelID, text)
	return msg.ID
}

func (st *Channel) SendEmbedMessage(text string) string {
	msg, _ := st.c.client.ChannelMessageSendEmbed(st.channelID, embed.NewGenericEmbed("", text))
	return msg.ID
}

func (st *Channel) DeleteMessage(msgID string) {
	st.c.client.ChannelMessageDelete(st.channelID, msgID)
}

func (st *Channel) DeleteMessages(number int) {
	msgs, _ := st.c.client.ChannelMessages(st.channelID, int(number)+1, "", "", "")
	// Get msg ids
	var msgIDs []string
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID)
	}
	// Delete msgs
	st.c.client.ChannelMessagesBulkDelete(st.channelID, msgIDs)
}

func (st *Channel) PinMessage(msgID string) {
	st.c.client.ChannelMessagePin(st.channelID, msgID)
}

func (st *Channel) UnpinMessage(msgID string) {
	st.c.client.ChannelMessageUnpin(st.channelID, msgID)
}

func (st *Channel) CreateInviteCode(maxUses int, temporary bool) string {
	invite, _ := st.c.client.ChannelInviteCreate(st.channelID, discordgo.Invite{
		MaxAge:    100,
		MaxUses:   maxUses,
		Temporary: temporary,
	})
	return invite.Code
}

func (st *Channel) Mention() string {
	channel, _ := st.c.client.State.Channel(st.channelID)
	return channel.Mention()
}

// TODO: Implement
// func (st *Channel) PermissionOverrideCreate() {}

// func (st *Channel) PermissionOverrideDelete() {}

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

// TODO: Implement
// func (st *Category) PermissionOverrideCreate() {}

// func (st *Category) PermissionOverrideDelete() {}

type DMChannel struct {
	channelID string
	c         *Connection
}

func (c *Connection) GetDMChannel(channelID string) DMChannel {
	return DMChannel{
		channelID: channelID,
		c:         c,
	}
}

func (st *DMChannel) GetName() string {
	channel, _ := st.c.client.Channel(st.channelID)
	return channel.Name
}

func (st *DMChannel) SendMessage(text string) string {
	msg, _ := st.c.client.ChannelMessageSend(st.channelID, text)
	return msg.ID
}

func (st *DMChannel) DeleteMessage(msgID string) {
	st.c.client.ChannelMessageDelete(st.channelID, msgID)
}

func (st *DMChannel) DeleteMessages(number int) {
	msgs, _ := st.c.client.ChannelMessages(st.channelID, int(number)+1, "", "", "")
	// Get msg ids
	var msgIDs []string
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID)
	}
	// Delete msgs
	st.c.client.ChannelMessagesBulkDelete(st.channelID, msgIDs)
}

func (st *DMChannel) PinMessage(msgID string) {
	st.c.client.ChannelMessagePin(st.channelID, msgID)
}

func (st *DMChannel) UnpinMessage(msgID string) {
	st.c.client.ChannelMessageUnpin(st.channelID, msgID)
}

func (st *DMChannel) Mention() string {
	channel, _ := st.c.client.State.Channel(st.channelID)
	return channel.Mention()
}
