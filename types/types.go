package types

import (
	dgo "github.com/bwmarrin/discordgo"
)

// Handler holds a refrence to the handler function for a application command
type Handler struct {
	Name     string
	Callback func(i *dgo.InteractionCreate, s *dgo.Session)
}

// Handler holds a refrence to the handler function for a application command
type Subcommand struct {
	Name     string
	Callback func(i *dgo.InteractionCreate, s *dgo.Session, option *dgo.ApplicationCommandInteractionDataOption)
}
