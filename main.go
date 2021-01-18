package main

import (
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

var (
	botToken = "ODAwNDkxNTIzNzYxNTA0Mjg2.YAS50w.NPCoNQslj3u-fPlIpEp_rQiEFGE"
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
			Content: "Warning",
		}
		// Sends the inital response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
		time.Sleep(1 * time.Second)
		// Get user from parms
		userID := i.Interaction.Data.Options[0].Value.(string)
		user, _ := client.User(userID)
		s.InteractionResponseEdit("", i.Interaction, &dgo.WebhookEdit{
			Content: "**" + user.Username + "** This is you final warning for " + i.Interaction.Data.Options[1].Value.(string),
		})
	})
	// Opens the connection
	client.Open()
	// Remove all commands
	deleteAllCommands(client)
	// Regesters the command
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
