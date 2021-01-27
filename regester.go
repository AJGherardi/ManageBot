package main

import (
	"github.com/AJGherardi/ManageBot/commands"
	dgo "github.com/bwmarrin/discordgo"
)

func regesterCommands(client *dgo.Session, guildID string) {
	commands.RegesterInit(client, guildID)
	commands.RegesterInvite(client, guildID)
	commands.RegesterKick(client, guildID)
	commands.RegesterPurge(client, guildID)
	commands.RegesterRemind(client, guildID)
	commands.RegesterRoles(client, guildID)
	commands.RegesterSay(client, guildID)
	commands.RegesterStats(client, guildID)
	commands.RegesterVote(client, guildID)
	commands.RegesterWarn(client, guildID)
}
