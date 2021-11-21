package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/api"
)

// PurgeHandler handles a purge command
type PurgeHandler struct{}

func (h *PurgeHandler) Name() string {
	return "purge"
}

func (h *PurgeHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	channel := c.GetChannel(i.GetChannelID())
	channel.DeleteMessages(i.GetIntParm(0))
	// Inform admin
	channel.SendEmbedMessage("Removed " + fmt.Sprint(i.GetIntParm(0)) + " messages")
}

func (h *PurgeHandler) Regester(c api.Connection) api.StandaloneCommandSignature {
	return api.MakeStandaloneCommandSignature("purge", "Removes specified number of msgs from current channel",
		api.MakeIntParmSignature("Number", "Number of messages to remove", true),
	)
}
