package utils

import (
	dgo "github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

// SendResponse sends a embed msg in the current channel
func SendResponse(response string, i *dgo.InteractionCreate, s *dgo.Session) {
	s.ChannelMessageSendEmbed(i.ChannelID, embed.NewGenericEmbed("", response))
}
