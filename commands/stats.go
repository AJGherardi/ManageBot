package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/api"
)

// StatsHandler handles a stat command
type StatsHandler struct{}

func (h *StatsHandler) Name() string {
	return "stats"
}

func (h *StatsHandler) Callback(i api.StandaloneCommandInvocation, c api.Connection) {
	// Get user list
	guild := c.GetGuild(i.GetGuildID())
	userIDs := guild.GetUserIDs()
	users := []api.User{}
	for _, userID := range userIDs {
		users = append(users, c.GetUser(userID))
	}
	// Get bot members
	bots := []api.User{}
	for _, user := range users {
		if user.IsBot() {
			bots = append(bots, user)
		}
	}
	// Check how many users are online start from one to count itself
	online := 1
	for _, user := range users {
		if user.IsOnline(i.GetGuildID()) {
			online++
		}
	}
	// Get boosters
	boosting := guild.GetBoosting()
	// Sends stats
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("There are " + fmt.Sprint(len(users)) + " members " + fmt.Sprint(online) + " of which are online")
	channel.SendEmbedMessage("There are " + fmt.Sprint(len(bots)) + " bots")
	channel.SendEmbedMessage("There are " + fmt.Sprint(boosting) + " people boosting your server")
}

func (h *StatsHandler) Regester(c api.Connection) api.StandaloneCommandSinginture {
	return api.MakeStandaloneCommandSinginture("stats", "Shows stats of a server or channel")
}
