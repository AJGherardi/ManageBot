package commands

import (
	"fmt"
	"time"

	"github.com/AJGherardi/ManageBot/api"
)

type VoteHandler struct{}

func (h *VoteHandler) Name() string {
	return "vote"
}

func (h *VoteHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	// Send vote message
	channel := c.GetChannel(i.GetChannelID())
	voteMessage := channel.SendEmbedMessageWithTitle(i.GetStringParm(0), i.GetStringParm(1))
	// Wait some time for message to appear
	time.Sleep(500 * time.Millisecond)
	// Add reactions
	channel.CreateReaction(voteMessage, "✅")
	channel.CreateReaction(voteMessage, "❌")
	// Start timer
	time.AfterFunc(
		(time.Duration(i.GetIntParm(2)) * time.Minute), func() {
			channel := c.GetChannel(i.GetChannelID())
			usersYes := channel.GetReactions(voteMessage, "✅")
			usersNo := channel.GetReactions(voteMessage, "❌")
			channel.SendEmbedMessage("Vote over " + i.GetStringParm(0))
			channel.SendEmbedMessage("Users for " + fmt.Sprint(len(usersYes)-1))
			channel.SendEmbedMessage("Users against " + fmt.Sprint(len(usersNo)-1))
		},
	)
}

func (h *VoteHandler) Regester(c api.Connection) api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture(
		"vote", "Make a vote",
		api.MakeStringParmSinginture("Title", "Title of vote message", true),
		api.MakeStringParmSinginture("Caption", "Caption for vote message", true),
		api.MakeIntParmSinginture("Time", "Time till end of vote in min", true),
	)
}
