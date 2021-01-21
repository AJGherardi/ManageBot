package commands

import (
	"fmt"
	"time"

	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

// HandleVote handles a vote command
func HandleVote(title, caption string, duration float64, i *dgo.InteractionCreate, s *dgo.Session) {
	// Send vote message
	voteMessage := sendVoteMessage(title, caption, i, s)
	// Wait some time for message to appear
	time.Sleep(500 * time.Millisecond)
	// Add reactions
	s.MessageReactionAdd(i.ChannelID, voteMessage, "✅")
	s.MessageReactionAdd(i.ChannelID, voteMessage, "❌")
	// Start timer
	time.AfterFunc(
		(time.Duration(duration) * time.Minute), func() {
			usersYes, _ := s.MessageReactions(i.ChannelID, voteMessage, "✅", 100, "", "")
			usersNo, _ := s.MessageReactions(i.ChannelID, voteMessage, "❌", 100, "", "")
			utils.SendResponse("Vote over "+title, i, s)
			utils.SendResponse("Users for "+fmt.Sprint(len(usersYes)-1), i, s)
			utils.SendResponse("Users against "+fmt.Sprint(len(usersNo)-1), i, s)
		},
	)
}

func sendVoteMessage(title, caption string, i *dgo.InteractionCreate, s *dgo.Session) string {
	voteMessage, _ := s.ChannelMessageSendEmbed(i.ChannelID, embed.NewGenericEmbed(title, caption))
	return voteMessage.ID
}
