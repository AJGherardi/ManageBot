package commands

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleStats handles a stat command
func HandleStats(i *dgo.InteractionCreate, s *dgo.Session) {
	// Gathers stats
	// Gathers number of bots
	numberOfBots := 0
	// Gathers number of members
	numberOfPariticipants := 0
	// Gathers number of members online
	numberOfPariticipantsOnline := 0
	// Gathers numbers of people with staff role
	numberOfStaff := 0
	// Gathers numbers of people with staff role that are online
	numberOfStaffOnline := 0
	// Does math to find the current member amount
	numberOfMembers := int(numberOfPariticipants) - int(numberOfBots)
	// Does math to find the current member amount online
	numberOfMembersOnline := int(numberOfPariticipantsOnline) - int(numberOfBots)
	// Sends stats
	utils.SendResponse("There are "+fmt.Sprint(numberOfBots)+" many bots", i, s)
	utils.SendResponse("There are "+fmt.Sprint(numberOfMembers)+" many users. "+fmt.Sprint(numberOfMembersOnline)+" of which are online", i, s)
	utils.SendResponse("There are "+fmt.SprintnumberOfStaff)+" many users. "+fmt.Sprint(numberOfStaffOnline)+" of which are online", i, s)
}
