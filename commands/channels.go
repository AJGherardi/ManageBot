package commands

import (
	"github.com/AJGherardi/ManageBot/api"
	dgo "github.com/bwmarrin/discordgo"
)

type ChannelHandler struct{}

func (h *ChannelHandler) Name() string {
	return "channel"
}

func (h *ChannelHandler) Subcommands() []api.Subcommand {
	return []api.Subcommand{
		&createChannelHandler{},
		&createGroupChannelHandler{},
		&deleteChannelHandler{},
	}
}

func (h *ChannelHandler) Regester(c api.Connection) api.ParentCommandSignature {
	return api.MakeParentCommandSignature("channel", "Manage channels")
}

type createChannelHandler struct{}

func (h *createChannelHandler) Name() string {
	return "create"
}

func (h *createChannelHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	guild := c.GetGuild(i.GetGuildID())
	guild.CreateChannel(i.GetStringParm(0), i.GetStringParm(1), i.GetIntParm(2), i.GetBoolParm(3))
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Added channel " + channel.Mention())
}

func (h *createChannelHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"create", "Adds a channel",
		api.MakeStringParmSignature("Name", "Name to give new channel", true),
		api.MakeChannelParmSignature("Category", "Category to add channel to", true),
		api.MakeParmSignatureWithChoices(
			"Type", "Type of new channel", true,
			api.Choice{Name: "Text", Value: dgo.ChannelTypeGuildText},
			api.Choice{Name: "Voice", Value: dgo.ChannelTypeGuildVoice},
		),
		api.MakeBoolParmSignature("NSFW", "Contains explicit material only applies to text channels", true),
	)
}

type createGroupChannelHandler struct{}

func (h *createGroupChannelHandler) Name() string {
	return "create-group"
}

func (h *createGroupChannelHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	guild := c.GetGuild(i.GetGuildID())
	guild.CreateCategory(i.GetStringParm(0))
	// Inform admin
	category := c.GetChannel(i.GetChannelID())
	category.SendEmbedMessage("Added channel group " + category.Mention())
}

func (h *createGroupChannelHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"create-group", "Adds a channel group",
		api.MakeStringParmSignature("Name", "Name to give new channel group", true),
	)
}

type deleteChannelHandler struct{}

func (h *deleteChannelHandler) Name() string {
	return "delete"
}

func (h *deleteChannelHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	guild := c.GetGuild(i.GetGuildID())
	guild.DeleteChannel(i.GetStringParm(0))
	// Inform admin
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Deleted channel " + channel.GetName())
}

func (h *deleteChannelHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"delete", "Remove a channel",
		api.MakeChannelParmSignature("Channel", "Channel to remove", true),
	)
}
