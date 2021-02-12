package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/api"
	dgo "github.com/bwmarrin/discordgo"
)

type ChannelHandler struct{}

func (h *ChannelHandler) Name() string {
	return "channel"
}

func (h *ChannelHandler) Subcommands() []api.Subcommand {
	return []api.Subcommand{
		&createHandler{},
		&createGroupHandler{},
		&deleteHandler{},
	}
}

func (h *ChannelHandler) Regester() api.ParentCommandSinginture {
	return api.MakeParentCommandSinginture("channel", "Manage channels")
}

type createHandler struct{}

func (h *createHandler) Name() string {
	return "create"
}

func (h *createHandler) Callback(i *dgo.InteractionCreate, s *dgo.Session) {
	fmt.Println("Create called")
}

func (h *createHandler) Regester() api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"create", "Adds a channel",
		api.MakeStringParmSinginture("Name", "Name to give new channel", true),
		api.MakeChannelParmSinginture("Category", "Category to add channel to", true),
		api.MakeIntParmSingintureWithChoices(
			"Type", "Type of new channel", true,
			api.Choice{Name: "Text", Value: dgo.ChannelTypeGuildText},
		),
		api.MakeBoolParmSinginture("NSFW", "Contains explicit material only applys to text channels", true),
	)
}

type createGroupHandler struct{}

func (h *createGroupHandler) Name() string {
	return "create-group"
}

func (h *createGroupHandler) Callback(i *dgo.InteractionCreate, s *dgo.Session) {
	fmt.Println("create-group called")
}

func (h *createGroupHandler) Regester() api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"create-group", "Adds a channel group",
		api.MakeStringParmSinginture("Name", "Name to give new channel group", true),
	)
}

type deleteHandler struct{}

func (h *deleteHandler) Name() string {
	return "delete"
}

func (h *deleteHandler) Callback(i *dgo.InteractionCreate, s *dgo.Session) {
	fmt.Println("delete called")
}

func (h *deleteHandler) Regester() api.SubcommandSinginture {
	return api.MakeSubcommandSinginture(
		"delete", "Remove a channel",
		api.MakeChannelParmSinginture("Channel", "Channel to remove", true),
	)
}
