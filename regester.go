package main

import (
	"github.com/AJGherardi/ManageBot/commands"
	"github.com/AJGherardi/ManageBot/types"
)

func getCommands() []types.StandaloneCommand {
	initHandler := commands.InitHandler{}
	return []types.StandaloneCommand{
		&initHandler,
	}
}
