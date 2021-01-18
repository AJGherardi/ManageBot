package main

import (
	"log"

	dgo "github.com/bwmarrin/discordgo"
)

var (
	botToken = "ODAwNDkxNTIzNzYxNTA0Mjg2.YAS50w.FvJlTQtvp0DbP4_mUxUbbOadvAY"
	guild    = "799794515443318814"
)

func main() {
	client, _ := dgo.New("Bot " + botToken)

	client.AddHandler(func(s *dgo.Session, i *dgo.InteractionCreate) {
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: i.Interaction.Data.Options[0].Value.(string),
		}
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
	})

	if err := client.Open(); err != nil {
		log.Fatal(err)
	}

	client.ApplicationCommandCreate("", &dgo.ApplicationCommand{
		Name:        "order",
		Description: "pls input your order",
		Options: []*dgo.ApplicationCommandOption{
			{
				Type:        dgo.ApplicationCommandOptionString,
				Name:        "Choices",
				Description: "posible orders",
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
	}, guild)

	for {

	}
	if err := client.Close(); err != nil {
		log.Fatal(err)
	}
}
