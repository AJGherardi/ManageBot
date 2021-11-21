package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

// StandaloneCommandSignature describes a command with no subcommands to be registered with the discord api
type StandaloneCommandSignature struct {
	Name, Description string
	Parms             []ParmSignature
}

// MakeStandaloneCommandSignature Creates and returns a signature for a StandaloneCommand
func MakeStandaloneCommandSignature(name, description string, parms ...ParmSignature) StandaloneCommandSignature {
	return StandaloneCommandSignature{
		Name:        name,
		Description: description,
		Parms:       parms,
	}
}

// ParentCommandSignature describes a command with subcommands to be registered with the discord api
type ParentCommandSignature struct {
	Name, Description string
}

// MakeParentCommandSignature Creates and returns a signature for a ParentCommand
func MakeParentCommandSignature(name, description string) ParentCommandSignature {
	return ParentCommandSignature{
		Name:        name,
		Description: description,
	}
}

// SubcommandSignature describes a subcommands to be registered with the discord api
type SubcommandSignature struct {
	Name, Description string
	Parms             []ParmSignature
}

// MakeSubcommandSignature Creates and returns a signature for a Subcommand
func MakeSubcommandSignature(name, description string, parms ...ParmSignature) SubcommandSignature {
	return SubcommandSignature{
		Name:        name,
		Description: description,
		Parms:       parms,
	}
}

// ParmSignature describes a paramater to be registered with the discord api
type ParmSignature interface {
	Build() *dgo.ApplicationCommandOption
}

// UnconstrainedParmSignature is a parm signature that dose not provide any required choices
type UnconstrainedParmSignature struct {
	Name, Description string
	Type              uint8
	Required          bool
}

// Build Converts a ParmSignature to application command option
func (p UnconstrainedParmSignature) Build() *dgo.ApplicationCommandOption {
	return &dgo.ApplicationCommandOption{
		Name:        p.Name,
		Description: p.Description,
		Required:    p.Required,
		Type:        dgo.ApplicationCommandOptionType(p.Type),
	}
}

// MakeIntParmSignature Creates and returns a signature for a integer parm
func MakeIntParmSignature(name, description string, required bool) ParmSignature {
	return makeParmSignatureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionInteger))
}

// MakeStringParmSignature Creates and returns a signature for a string parm
func MakeStringParmSignature(name, description string, required bool) ParmSignature {
	return makeParmSignatureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionString))
}

// MakeBoolParmSignature Creates and returns a signature for a bool parm
func MakeBoolParmSignature(name, description string, required bool) ParmSignature {
	return makeParmSignatureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionBoolean))
}

// MakeChannelParmSignature Creates and returns a signature for a channel parm
func MakeChannelParmSignature(name, description string, required bool) ParmSignature {
	return makeParmSignatureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionChannel))
}

// MakeUserParmSignature Creates and returns a signature for a user parm
func MakeUserParmSignature(name, description string, required bool) ParmSignature {
	return makeParmSignatureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionUser))
}

// MakeRoleParmSignature Creates and returns a signature for a role parm
func MakeRoleParmSignature(name, description string, required bool) ParmSignature {
	return makeParmSignatureWithType(name, description, required, uint8(dgo.ApplicationCommandOptionRole))
}

func makeParmSignatureWithType(name, description string, required bool, parmType uint8) ParmSignature {
	return UnconstrainedParmSignature{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        parmType,
	}
}

// ConstrainedParmSignature is a parm signature that has a list of required choices
type ConstrainedParmSignature struct {
	Name, Description string
	Type              uint8
	Required          bool
	Choices           []Choice
}

// Build Converts a ParmSignature to application command option
func (p ConstrainedParmSignature) Build() *dgo.ApplicationCommandOption {
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

// MakeParmSignatureWithChoices returns a parm with required choices
func MakeParmSignatureWithChoices(name, description string, required bool, choices ...Choice) ParmSignature {
	return ConstrainedParmSignature{
		Name:        name,
		Description: description,
		Required:    required,
		Type:        uint8(dgo.ApplicationCommandOptionInteger),
		Choices:     choices,
	}
}
