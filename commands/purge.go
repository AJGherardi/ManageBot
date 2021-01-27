package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandlePurge handles a purge command
func HandlePurge(number float64, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get msgs
	msgs, _ := s.ChannelMessages(i.ChannelID, int(number)+1, "", "", "")
	// Get msg ids
	var msgIDs []string
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID)
	}
	// Delete msgs
	s.ChannelMessagesBulkDelete(i.ChannelID, msgIDs)
	utils.SendResponse("Removed "+fmt.Sprint(number)+" messages", i, s)
}

// RegesterPurge adds the kick / command
func RegesterPurge(client *dgo.Session, guildID string) {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "purge",
			Description: "Removes specified number of msgs from current channel",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Number",
					Description: "Number of messages to remove",
					Required:    true,
				},
			},
		},
		guildID,
	)
}
