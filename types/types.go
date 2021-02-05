package types

import (
	dgo "github.com/bwmarrin/discordgo"
)

// Handler holds a refrence to the handler function for a application command
type Handler struct {
	Name     string
	Callback func(i *dgo.InteractionCreate, s *dgo.Session)
}

// Subcommand holds a refrence to the handler function for a application subcommand
type Subcommand struct {
	Name     string
	Callback func(parms SubcommandParms)
}

// SubcommandParms holds all neccesary parms to invoke a subcommand
type SubcommandParms struct {
	Interaction *dgo.InteractionCreate
	Session     *dgo.Session
	Option      *dgo.ApplicationCommandInteractionDataOption
}

// ServerData holds all information that is needed to manage a discord guild
type ServerData struct {
	GuildID string `bson:"guildID,omitempty"`
	Name    string `bson:"name,omitempty"`
	// CommandHistory
	// Tickets
	// Warnings
}

// GetGuildID returns the servers id
func (s *ServerData) GetGuildID() string {
	return s.GuildID
}

// ChangeName changes the servers name
func (s *ServerData) ChangeName(name string) {
	s.Name = name
}

// GetName retuens the servers name
func (s *ServerData) GetName() string {
	return s.Name
}
