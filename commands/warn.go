package commands

import (
	"github.com/AJGherardi/ManageBot/api"
)

// WarnHandler handles a warn command
type WarnHandler struct{}

func (h *WarnHandler) Name() string {
	return "warn"
}

func (h *WarnHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	// Get user
	user := c.GetUser(i.GetStringParm(0))
	// Send warning
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage(user.Mention() + " This is you final warning for " + i.GetStringParm(1))
}

func (h *WarnHandler) Regester() api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("warn", "Warn for user rule violation",
		api.MakeUserParmSinginture("User", "User to warn", true),
		api.MakeParmSingintureWithChoices("Violation", "Rule violated", true,
			api.Choice{Name: "Gore", Value: "Gore"},
			api.Choice{Name: "Harassment", Value: "Harassment"},
			api.Choice{Name: "Disrespecting staff", Value: "Disrespecting staff"},
			api.Choice{Name: "Sexually explicit content", Value: "Sexually explicit content"},
			api.Choice{Name: "Advertizing", Value: "Advertizing"},
			api.Choice{Name: "Spam", Value: "Spam"},
			api.Choice{Name: "Obsessive Pinging", Value: "Obsessive Pinging"},
			api.Choice{Name: "Hate Speech", Value: "Hate Speech"},
			api.Choice{Name: "Threatening People", Value: "Threatening People"},
			api.Choice{Name: "Sending Dangerous Links", Value: "Sending Dangerous Links"},
		),
	)
}
