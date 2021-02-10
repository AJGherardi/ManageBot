package utils

import (
	dgo "github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

// SendDM sends a embed msg in the user dm
func SendDM(response, channelID string, s *dgo.Session) *dgo.Message {
	message, _ := s.ChannelMessageSendEmbed(channelID, embed.NewGenericEmbed("", response))
	return message
}

// SendResponse sends a embed msg in the current channel
func SendResponse(response string, i *dgo.InteractionCreate, s *dgo.Session) *dgo.Message {
	message, _ := s.ChannelMessageSendEmbed(i.ChannelID, embed.NewGenericEmbed("", response))
	return message
}
