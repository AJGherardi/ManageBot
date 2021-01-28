package utils

import (
	"github.com/AJGherardi/ManageBot/types"
	dgo "github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

// SendResponse sends a embed msg in the current channel
func SendResponse(response string, i *dgo.InteractionCreate, s *dgo.Session) {
	s.ChannelMessageSendEmbed(i.ChannelID, embed.NewGenericEmbed("", response))
}

// MatchSubcommand searches a list of subcommands for a matching handler for a / command
func MatchSubcommand(i *dgo.InteractionCreate, s *dgo.Session, subcommands []types.Subcommand) {
	for _, subcommand := range subcommands {
		if subcommand.Name == i.Interaction.Data.Options[0].Name {
			subcommand.Callback(types.SubcommandParms{
				Interaction: i,
				Session:     s,
				Option:      i.Data.Options[0],
			})
		}
	}
}
