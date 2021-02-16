package api

import dgo "github.com/bwmarrin/discordgo"

// StandaloneCommand defines the behavior of a standalone command
type StandaloneCommand interface {
	// Returns the name of the command which is used to match the command to its callback
	Name() string
	// Handels a invocation of this command
	Callback(i StandaloneCommandInvocation, c Connection)
	// Builds and returns a descscription of a application command and its parameters
	// which is regestred with the discord api
	Regester() StandaloneCommandSinginture
}

// ParentCommand defines the behavior of a parent command which contains subcommands
type ParentCommand interface {
	// Returns the name of the command which is used to match the command to its subcommands
	Name() string
	// Returns a list of the subcommands for this parent command
	Subcommands() []Subcommand
	// Builds and returns a descscription of a parent command
	// which is regestered at the same time as the parrent command
	Regester() ParentCommandSinginture
}

// Subcommand defines the behavior of a parent commands subcommand
type Subcommand interface {
	// Returns the name of the command which is used to match the subcommand to its callback
	Name() string
	// Handels a invocation of this command
	Callback(i SubcommandInvocation, c Connection)
	// Builds and returns a descscription of a subcommand and its parameters
	// which is regestered at the same time as the parrent command
	Regester() SubcommandSinginture
}

type invocation struct {
	i *dgo.InteractionCreate
}

// StandaloneCommandInvocation Wrapes a discord interaction and makes it easyer to get parameters
type StandaloneCommandInvocation struct {
	invocation
}

// GetGuildID returns the id of the guild that the command was invoked in
func (i *invocation) GetGuildID() string {
	return i.i.GuildID
}

// GetChannelID returns the id of the channel that the command was invoked in
func (i *invocation) GetChannelID() string {
	return i.i.ChannelID
}

// GetUserID returns the id of the user who invoked the command
func (i *invocation) GetUserID() string {
	return i.i.Member.User.ID
}

// GetIntParm gets a interger parameter
func (i *StandaloneCommandInvocation) GetIntParm(index int) int {
	return int(i.i.Interaction.Data.Options[index].Value.(float64))
}

// GetStringParm gets a interger parameter
func (i *StandaloneCommandInvocation) GetStringParm(index int) string {
	return i.i.Interaction.Data.Options[index].Value.(string)
}

// GetBoolParm gets a boolean parameter
func (i *StandaloneCommandInvocation) GetBoolParm(index int) bool {
	return i.i.Interaction.Data.Options[index].Value.(bool)
}

// SubcommandInvocation Wrapes a discord interaction and makes it easyer to get parameters
type SubcommandInvocation struct {
	invocation
}

// GetIntParm gets a interger parameter
func (i *SubcommandInvocation) GetIntParm(index int) int {
	return int(i.i.Interaction.Data.Options[0].Options[index].Value.(float64))
}

// GetStringParm gets a interger parameter
func (i *SubcommandInvocation) GetStringParm(index int) string {
	return i.i.Interaction.Data.Options[0].Options[index].Value.(string)
}

// GetBoolParm gets a boolean parameter
func (i *SubcommandInvocation) GetBoolParm(index int) bool {
	return i.i.Interaction.Data.Options[0].Options[index].Value.(bool)
}
