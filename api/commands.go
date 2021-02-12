package api

import dgo "github.com/bwmarrin/discordgo"

// StandaloneCommand defines the behavior of a standalone command
type StandaloneCommand interface {
	// Returns the name of the command which is used to match the command to its callback
	Name() string
	// Handels a invocation of this command
	Callback(i *dgo.InteractionCreate, s *dgo.Session)
	// Builds and returns a descscription of a application command and its parameters
	// which is regestred with the discord api
	Regester() StandaloneCommandSinginture
}

// ParentCommand defines the behavior of a parent command which contains subcommands
type ParentCommand interface {
	// Returns the name of the command which is used to match the command to its subcommands
	Name() string
	// Returns a list of the subcommands for this parent command
	Subcommands()
	// Builds and returns a descscription of a parent command
	// which is regestered at the same time as the parrent command
	Regester() ParentCommandSinginture
}

// Subcommand defines the behavior of a parent commands subcommand
type Subcommand interface {
	// Returns the name of the command which is used to match the subcommand to its callback
	Name() string
	// Handels a invocation of this command
	Callback(i *dgo.InteractionCreate, s *dgo.Session)
	// Builds and returns a descscription of a subcommand and its parameters
	// which is regestered at the same time as the parrent command
	Regester() SubcommandSinginture
}

// StandaloneCommandInvocation Wrapes a discord interaction and makes it easyer to get parameters
type StandaloneCommandInvocation struct {
	I *dgo.InteractionCreate
}

// GetIntParm gets a interger parameter
func (i *StandaloneCommandInvocation) GetIntParm(index int) int {
	return 0
}

// GetStringParm gets a interger parameter
func (i *StandaloneCommandInvocation) GetStringParm(index int) string {
	return ""
}

// GetBoolParm gets a boolean parameter
func (i *StandaloneCommandInvocation) GetBoolParm(index int) bool {
	return false
}

// SubcommandInvocation Wrapes a discord interaction and makes it easyer to get parameters
type SubcommandInvocation struct {
	I *dgo.InteractionCreate
}

// GetIntParm gets a interger parameter
func (i *SubcommandInvocation) GetIntParm(index int) int {
	return 0
}

// GetStringParm gets a interger parameter
func (i *SubcommandInvocation) GetStringParm(index int) string {
	return ""
}

// GetBoolParm gets a boolean parameter
func (i *SubcommandInvocation) GetBoolParm(index int) bool {
	return false
}
