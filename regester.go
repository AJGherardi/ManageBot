package main

import (
	"github.com/AJGherardi/ManageBot/types"

	"github.com/AJGherardi/ManageBot/commands"
	dgo "github.com/bwmarrin/discordgo"
)

func regesterCommands(client *dgo.Session, guildID string) []types.Handler {
	channelHandler := commands.RegesterChannel(client, guildID)
	initHandler := commands.RegesterInit(client, guildID)
	inviteHandler := commands.RegesterInvite(client, guildID)
	kickHandler := commands.RegesterKick(client, guildID)
	nicknameHandler := commands.RegesterNickname(client, guildID)
	purgeHandler := commands.RegesterPurge(client, guildID)
	remindHandler := commands.RegesterRemind(client, guildID)
	rolesHandler := commands.RegesterRoles(client, guildID)
	sayHandler := commands.RegesterSay(client, guildID)
	statsHandler := commands.RegesterStats(client, guildID)
	voteHandler := commands.RegesterVote(client, guildID)
	warnHandler := commands.RegesterWarn(client, guildID)
	muteHandler := commands.RegesterMute(client, guildID)
	return []types.Handler{
		channelHandler,
		initHandler,
		inviteHandler,
		kickHandler,
		nicknameHandler,
		purgeHandler,
		remindHandler,
		rolesHandler,
		sayHandler,
		statsHandler,
		voteHandler,
		warnHandler,
		muteHandler,
	}
}
