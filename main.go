package main

import (
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

// botToken and guildID must be added to consts.go

func main() {
	// Creates a new client object
	client, _ := dgo.New("Bot " + botToken)
	// Regesters a event handeler for when the command is called
	client.AddHandler(commandHandler(client))
	// Opens the connection
	client.Open()
	// Remove all commands
	deleteAllCommands(client)
	// Regesters the commands
	regesterCommands(client)
	// Keep the app runing
	for {
	}
}

func commandHandler(client *dgo.Session) func(s *dgo.Session, i *dgo.InteractionCreate) {
	return func(s *dgo.Session, i *dgo.InteractionCreate) {
		// Makes a reaponse
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: "Pls wait",
		}
		// Sends the inital response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
		// Wait a sec
		time.Sleep(1 * time.Second)
		// Match command to handler function
		switch i.Interaction.Data.Name {
		case "warn":
			handleWarn(
				i.Interaction.Data.Options[0].Value.(string),
				i.Interaction.Data.Options[1].Value.(string),
				i,
				s,
			)
		}
	}
}

func handleWarn(userID, violation string, i *dgo.InteractionCreate, s *dgo.Session) {
	// Get user from parms
	user, _ := s.User(userID)
	s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
		Content: user.Mention() + " This is you final warning for " + violation,
	})
}

func regesterCommands(client *dgo.Session) {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "Warn",
			Description: "Warn for user rule violation",

			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionUser,
					Name:        "User",
					Description: "User to warn",
					Required:    true,
				},
				{
					Type:        dgo.ApplicationCommandOptionString,
					Name:        "Violation",
					Description: "Rules violated",
					Choices: []*dgo.ApplicationCommandOptionChoice{
						{
							Name:  "Gore",
							Value: "Gore",
						},
						{
							Name:  "Harassment",
							Value: "Harassment",
						},
						{
							Name:  "Disrespecting staff",
							Value: "Disrespecting staff",
						},
						{
							Name:  "Sexually explicit content",
							Value: "Sexually explicit content",
						},
						{
							Name:  "Advertizing",
							Value: "Advertizing",
						},
						{
							Name:  "Spam",
							Value: "Spam",
						},
						{
							Name:  "Obsessive pinging",
							Value: "Obsessive Pinging",
						},
						{
							Name:  "Hate Speech",
							Value: "Hate Speech",
						},
						{
							Name:  "Threatening People",
							Value: "Threatening People",
						},
						{
							Name:  "Sending Dangerous Links",
							Value: "Sending Dangerous Links",
						},
					},
					Required: true,
				},
			},
		},
		guildID,
	)
}

func deleteAllCommands(client *dgo.Session) {
	apps, _ := client.Applications()
	for _, app := range apps {
		cmds, _ := client.ApplicationCommands(app.ID, guildID)
		for _, cmd := range cmds {
			client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
		}
	}
	cmds, _ := client.ApplicationCommands("", guildID)
	for _, cmd := range cmds {
		client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
	}
}
