package commands

import (
	"fmt"

	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleStats handles a stat command
func HandleStats(i *dgo.InteractionCreate, s *dgo.Session) {
	// Gathers stats
	// Gathers number of members
		MemberCount int `json:"member_count"`
	// Gathers numbers of people with staff role
		numberOfStaff := 0
	// Gathers numbers of people with staff role that are online
		numberOfStaffOnline := 0
		//if /*has role staff*/ && /*is online*/ && /*!marked searched for staff*/{
			//find some way to mark off that user
			//numberOfStaffOnline ++
			//}
			//else /*skip user*/
	// Finds the current member amount online
		ApproximatePresenceCount int `json:"approximate_presence_count"`
	// Finds number of current boosters
		PremiumSubscriptionCount int `json:"premium_subscription_count"`
	// Finds number of bots
		numberOfBots := 0
		//if /*is bot*/ && /*!marked searched for bot*/{
			//find some way to mark off that user
			//}
			//else /*skip user*/
	// Sends stats
	utils.SendResponse("There are "+fmt.Sprint(MemberCount)+" many users "+fmt.Sprint(ApproximatePresenceCount)+" of which are online", i, s)
	utils.SendResponse("There are "+fmt.Sprint(numberOfStaff)+" many users "+fmt.Sprint(numberOfStaffOnline)+" of which are online", i, s)
	utils.SendResponse("There are "+fmt.Sprint(numberOfBots)+" many bots", i, s)
	utils.SendResponse("There are "+fmt.Sprint(PremiumSubscriptionCount)+" current boosters", i, s)
}

