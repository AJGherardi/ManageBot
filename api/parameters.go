package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

// SubcommandSinginture describes a subcommands to be registered with the discord api
type SubcommandSinginture struct {
	Name, Description string
	Parms             []ParmSinginture
}

// ParentCommandSinginture describes a command with subcommands to be registered with the discord api
type ParentCommandSinginture struct {
	Name, Description string
}

// StandaloneCommandSinginture describes a command with no subcommands to be registered with the discord api
type StandaloneCommandSinginture struct {
	Name, Description string
	Parms             []ParmSinginture
}

// ParmSinginture describes a paramater to be registered with the discord api
type ParmSinginture struct {
	Name, Description string
	Type              uint8
	Required          bool
}

// MakeStandaloneCommandSinginture Creates and returns a signiture for a StandaloneCommand
func MakeStandaloneCommandSinginture(name, description string, parms ...ParmSinginture) StandaloneCommandSinginture {
	return StandaloneCommandSinginture{
		Name:        name,
		Description: description,
		Parms:       parms,
	}
}

// MakeParentCommandSinginture Creates and returns a signiture for a ParentCommand
func MakeParentCommandSinginture(name, description string) ParentCommandSinginture {
	return ParentCommandSinginture{
		Name:        name,
		Description: description,
	}
}

// MakeSubcommandSinginture Creates and returns a signiture for a Subcommand
func MakeSubcommandSinginture(name, description string, parms ...ParmSinginture) SubcommandSinginture {
	return SubcommandSinginture{
		Name:        name,
		Description: description,
		Parms:       parms,
	}
}

// MakeIntParmSinginture Creates and returns a signiture for a intager parm
func MakeIntParmSinginture(name, description string, required bool) ParmSinginture {
	return makeParmSingitureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionInteger))
}

// MakeStringParmSinginture Creates and returns a signiture for a string parm
func MakeStringParmSinginture(name, description string, required bool) ParmSinginture {
	return makeParmSingitureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionString))
}

// MakeBoolParmSinginture Creates and returns a signiture for a bool parm
func MakeBoolParmSinginture(name, description string, required bool) ParmSinginture {
	return makeParmSingitureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionBoolean))
}

// MakeChannelParmSinginture Creates and returns a signiture for a channel parm
func MakeChannelParmSinginture(name, description string, required bool) ParmSinginture {
	return makeParmSingitureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionChannel))
}

// MakeUserParmSinginture Creates and returns a signiture for a user parm
func MakeUserParmSinginture(name, description string, required bool) ParmSinginture {
	return makeParmSingitureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionUser))
}

// MakeRoleParmSinginture Creates and returns a signiture for a role parm
func MakeRoleParmSinginture(name, description string, required bool) ParmSinginture {
	return makeParmSingitureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionRole))
}

func makeParmSingitureWithType(name, description string, required bool, parmType uint8) ParmSinginture {
	return ParmSinginture{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        parmType,
	}
}
