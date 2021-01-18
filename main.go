package main

import (
	dgo "github.com/bwmarrin/discordgo"
)

var (
	botToken = "ODAwNDkxNTIzNzYxNTA0Mjg2.YAS50w.pIKeG2AYPWdHW9Gzyb3V82Gh2iA"
	guildID  = "799794515443318814"
)

// Execution starts here
func main() {
	// Creates a new client object
	client, _ := dgo.New("Bot " + botToken)
	// Regesters a event handeler for when the command is called
	client.AddHandler(func(s *dgo.Session, i *dgo.InteractionCreate) {
		// Makes a reaponse
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: i.Interaction.Data.Options[0].Value.(string),
		}
		// Sends the response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
	})
	// Opens the connection
	client.Open()
	// Regesters the command
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "order",
			Description: "pls input your order",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Choices",
					Description: "possible orders",
					Choices: []*dgo.ApplicationCommandOptionChoice{
						{
							Name:  "Food",
							Value: "Food",
						},
						{
							Name:  "Drink",
							Value: "Drink",
						},
					},
					Required: true,
				},
			},
		},
		guildID,
	)
	// Keep the app runing
	for {

	}
	// if err := client.Close(); err != nil {
	// 	log.Fatal(err)
	// }
} // Execution ends here
