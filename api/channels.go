package api

type Channel struct {
	guildID   string
	channelID string
}

func (c *Connection) GetChannel(guildID, channelID string) {}

func (st *Channel) GetName() {}

func (st *Channel) GetParentID() {}

func (st *Channel) SendMessage() {}

func (st *Channel) DeleteMessage() {}

func (st *Channel) DeleteMessages() {}

func (st *Channel) PinMessage() {}

func (st *Channel) UnpinMessage() {}

func (st *Channel) PermissionOverrideCreate() {}

func (st *Channel) PermissionOverrideDelete() {}

type Category struct {
	guildID   string
	channelID string
}

func (c *Connection) GetCategory(guildID, channelID string) {}

func (st *Category) GetName() {}

func (st *Category) GetChannelIDs() {}

func (st *Category) PermissionOverrideCreate() {}

func (st *Category) PermissionOverrideDelete() {}

type DMChannel struct {
	channelID string
}

func (c *Connection) GetDMChannel(channelID string) {}

func (st *DMChannel) GetName() {}

func (st *DMChannel) SendMessage() {}

func (st *DMChannel) DeleteMessage() {}

func (st *DMChannel) DeleteMessages() {}

func (st *DMChannel) PinMessage() {}

func (st *DMChannel) UnpinMessage() {}
