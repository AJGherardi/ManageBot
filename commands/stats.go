package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleStats handles a stat command
func HandleStats(i *dgo.InteractionCreate, s *dgo.Session) {
	// Get members list
	members, _ := s.GuildMembers(i.GuildID, "", 100)
	// Get bot members
	bots := []*dgo.Member{}
	for _, member := range members {
		if member.User.Bot {
			bots = append(bots, member)
		}
	}
	// Sends stats
	utils.SendResponse("There are "+fmt.Sprint(len(members))+" many members", i, s)
	utils.SendResponse("There are "+fmt.Sprint(len(bots))+" many bots", i, s)
}
