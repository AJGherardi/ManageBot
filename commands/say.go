package commands

import (
	"github.com/AJGherardi/ManageBot/api"
)

type SayHandler struct{}

func (h *SayHandler) Name() string {
	return "say"
}

func (h *SayHandler) Subcommands() []api.Subcommand {
	return []api.Subcommand{
		&messageSayHandler{},
		&dmSayHandler{},
		&reactionRoleSayHandler{},
	}
}

func (h *SayHandler) Regester(c api.Connection) api.ParentCommandSignature {
	return api.MakeParentCommandSignature("say", "More powerful messages")
}

type messageSayHandler struct{}

func (h *messageSayHandler) Name() string {
	return "message"
}

func (h *messageSayHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	channel := c.GetChannel(i.GetChannelID())
	// Repeat sending messages
	for r := 0; r < i.GetIntParm(1); r++ {
		// Check if embed
		if i.GetBoolParm(2) {
			channel.SendEmbedMessage(i.GetStringParm(0))
		} else {
			channel.SendMessage(i.GetStringParm(0))
		}
	}
}

func (h *messageSayHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"message", "Repeats a message",
		api.MakeStringParmSignature("Message", "Message to repeat", true),
		api.MakeIntParmSignature("Repeat", "Number of times to repeat", true),
		api.MakeBoolParmSignature("Embed", "Sends message in a embed", true),
	)
}

type dmSayHandler struct{}

func (h *dmSayHandler) Name() string {
	return "dm"
}

func (h *dmSayHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	user := c.GetUser(i.GetStringParm(0))
	dmChannel := c.GetDMChannel(user.GetDMChannelID())
	// Repeat sending messages
	for r := 0; r < i.GetIntParm(2); r++ {
		// Check if embed
		if i.GetBoolParm(3) {
			dmChannel.SendEmbedMessage(i.GetStringParm(1))
		} else {
			dmChannel.SendMessage(i.GetStringParm(1))
		}
	}
}

func (h *dmSayHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"dm", "Sends a dm",
		api.MakeUserParmSignature("User", "User to dm", true),
		api.MakeStringParmSignature("Message", "Message to repeat", true),
		api.MakeIntParmSignature("Repeat", "Number of times to repeat", true),
		api.MakeBoolParmSignature("Embed", "Sends message in a embed", true),
	)
}

type reactionRole struct {
	EmojiID   string
	MessageID string
	RoleID    string
	GuildID   string
}

var reactionRoles []reactionRole

type reactionRoleSayHandler struct{}

func (h *reactionRoleSayHandler) Name() string {
	return "reaction-role"
}

func (h *reactionRoleSayHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	// Send message and get msgID
	channel := c.GetChannel(i.GetChannelID())
	msgID := channel.SendEmbedMessage(i.GetStringParm(0))
	// Add initial reaction
	channel.CreateReaction(msgID, i.GetStringParm(1))
	// Add reaction role
	reactionRoles = append(reactionRoles, reactionRole{
		EmojiID:   i.GetStringParm(1),
		MessageID: msgID,
		RoleID:    i.GetStringParm(2),
		GuildID:   i.GetGuildID(),
	})
}

func (h *reactionRoleSayHandler) Regester(c api.Connection) api.SubcommandSignature {
	// Start reaction handler
	c.StartReactionHandler(func(c api.Connection, msgID, userID, emojiName string) {
		// Match Reaction Role
		for _, rr := range reactionRoles {
			if rr.EmojiID == emojiName {
				if msgID == rr.MessageID {
					guild := c.GetGuild(rr.GuildID)
					guild.AssignRole(userID, rr.RoleID)
				}
			}
		}
	})
	return api.MakeSubcommandSignature(
		"reaction-role", "Grants role to whoever reacts using given emoji",
		api.MakeStringParmSignature("Message", "Message content", true),
		api.MakeStringParmSignature("Emoji", "Emoji for reaction", true),
		api.MakeRoleParmSignature("Role", "Role assigned", true),
	)
}
