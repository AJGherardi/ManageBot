package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

// Connection Wraps the discord api for easy use
type Connection struct {
	client *dgo.Session
}

// ConnectToDiscord Opens a new connection to discord
func ConnectToDiscord(botToken, guildID string) Connection {
	// Creates a new client object
	client, _ := dgo.New("Bot " + botToken)
	// Set intents
	client.Identify.Intents = dgo.MakeIntent(
		dgo.IntentsAllWithoutPrivileged |
			dgo.IntentsGuildPresences |
			dgo.IntentsGuildMembers,
	)
	// Opens the connection
	client.Open()
	// Remove all commands
	deleteAllCommands(client, guildID)
	return Connection{client: client}
}
