package types

import dgo "github.com/bwmarrin/discordgo"

type StandaloneCommand interface {
	Name() string
	Callback(i *dgo.InteractionCreate, s *dgo.Session)
	Regester(client *dgo.Session, guildID string)
}

type ParentCommand interface {
	Name() string
	Subcommands()
	Regester(client *dgo.Session, guildID string)
}

type Subcommand interface {
	Name() string
	Callback(i *dgo.InteractionCreate, s *dgo.Session)
	Regester(client *dgo.Session, guildID string)
}

type CommandInvocation struct {
	I *dgo.InteractionCreate
	S *dgo.Session
}

func (i *CommandInvocation) GetIntParm() {

}

func (i *CommandInvocation) GetStringParm() {

}

func (i *CommandInvocation) GetBoolParm() {

}

type SubcommandInvocation struct {
	I *dgo.InteractionCreate
}

func (i *SubcommandInvocation) GetIntParm() {

}

func (i *SubcommandInvocation) GetStringParm() {

}

func (i *SubcommandInvocation) GetBoolParm() {

}
