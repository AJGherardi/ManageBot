package commands

import (
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

// HandleVote handles a vote command
func HandleVote(i *dgo.InteractionCreate, s *dgo.Session) {
	voteMessage, _ := s.ChannelMessageSend(i.ChannelID, "test")
	time.Sleep(1000 * time.Millisecond)
	s.MessageReactionAdd(i.ChannelID, voteMessage.ID, "✅")
	s.MessageReactionAdd(i.ChannelID, voteMessage.ID, "❌")
}
