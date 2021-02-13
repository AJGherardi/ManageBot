package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

// StandaloneCommandSinginture describes a command with no subcommands to be registered with the discord api
type StandaloneCommandSinginture struct {
	Name, Description string
	Parms             []ParmSinginture
}

// MakeStandaloneCommandSinginture Creates and returns a signiture for a StandaloneCommand
func MakeStandaloneCommandSinginture(name, description string, parms ...ParmSinginture) StandaloneCommandSinginture {
	return StandaloneCommandSinginture{
		Name:        name,
		Description: description,
		Parms:       parms,
	}
}

// ParentCommandSinginture describes a command with subcommands to be registered with the discord api
type ParentCommandSinginture struct {
	Name, Description string
}

// MakeParentCommandSinginture Creates and returns a signiture for a ParentCommand
func MakeParentCommandSinginture(name, description string) ParentCommandSinginture {
	return ParentCommandSinginture{
		Name:        name,
		Description: description,
	}
}

// SubcommandSinginture describes a subcommands to be registered with the discord api
type SubcommandSinginture struct {
	Name, Description string
	Parms             []ParmSinginture
}

// MakeSubcommandSinginture Creates and returns a signiture for a Subcommand
func MakeSubcommandSinginture(name, description string, parms ...ParmSinginture) SubcommandSinginture {
	return SubcommandSinginture{
		Name:        name,
		Description: description,
		Parms:       parms,
	}
}

// ParmSinginture describes a paramater to be registered with the discord api
type ParmSinginture interface {
	Build() *dgo.ApplicationCommandOption
}

// UnconstrainedParmSinginture is a parm singinture that dose not provide any required choices
type UnconstrainedParmSinginture struct {
	Name, Description string
	Type              uint8
	Required          bool
}

// Build Converts a ParmSinginture to application command option
func (p UnconstrainedParmSinginture) Build() *dgo.ApplicationCommandOption {
	return &dgo.ApplicationCommandOption{
		Name:        p.Name,
		Description: p.Description,
		Required:    p.Required,
		Type:        dgo.ApplicationCommandOptionType(p.Type),
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
	return UnconstrainedParmSinginture{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        parmType,
	}
}

// ConstrainedParmSinginture is a parm singinture that has a list of required choices
type ConstrainedParmSinginture struct {
	Name, Description string
	Type              uint8
	Required          bool
	Choices           []Choice
}

// Build Converts a ParmSinginture to application command option
func (p ConstrainedParmSinginture) Build() *dgo.ApplicationCommandOption {
	// Build choices
	parmChoices := []*dgo.ApplicationCommandOptionChoice{}
	for _, choice := range p.Choices {
		parmChoices = append(parmChoices, choice.Build())
	}
	return &dgo.ApplicationCommandOption{
		Name:        p.Name,
		Description: p.Description,
		Required:    p.Required,
		Type:        dgo.ApplicationCommandOptionType(p.Type),
		Choices:     parmChoices,
	}
}

// Choice returns a name/value pair
type Choice struct {
	Name  string
	Value interface{}
}

// Build converts a Choice to a Application command choice
func (c Choice) Build() *dgo.ApplicationCommandOptionChoice {
	return &dgo.ApplicationCommandOptionChoice{
		Name:  c.Name,
		Value: c.Value,
	}
}

// MakeIntParmSingintureWithChoices returns a parm with required choices
func MakeIntParmSingintureWithChoices(name, description string, required bool, choices ...Choice) ParmSinginture {
	return ConstrainedParmSinginture{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        uint8(dgo.ApplicationCommandOptionInteger),
		Choices:     choices,
	}
}
