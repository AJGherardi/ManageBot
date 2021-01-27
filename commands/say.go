package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleSay handles a say command
func HandleSay(message string, number float64, i *dgo.InteractionCreate, s *dgo.Session) {
	for r := 0; r < int(number); r++ {
		utils.SendResponse(message, i, s)
	}
}

// RegesterSay adds the say / command
func RegesterSay(client *dgo.Session, guildID string) {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "say",
			Description: "Repeats a message",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Message",
					Description: "Message to repeat",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionInteger,
					Name:        "Repeat",
					Description: "Number of times to repeat",
					Required:    true,
				},
			},
		},
		guildID,
	)
}
