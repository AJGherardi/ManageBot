package stats


// HandleStats handles a stat command
func HandleStats(i *dgo.InteractionCreate, s *dgo.Session) {
// Gathers stats
	// Gathers number of bots
		numberOfBots := "holderoftheplaces"
	// Gathers number of members
		numberOfPariticipants := "holderoftheplaces"
	// Gathers number of members online
		numberOfPariticipantsOnline := "holderoftheplaces"
	// Gathers numbers of people with staff role
		numberOfStaff := "holderoftheplaces"
	// Gathers numbers of people with staff role that are online
		numberOfStaffOnline := "holderoftheplaces"
	// Sends stats
	utils.SendResponse("There are " numberOfBots " many bots", i, s)
	utils.SendResponse("There are " numberOfPariticipants-numberOfBots " many users. " numberOfPariticipantsOnline " of which are online", i, s)
	utils.SendResponse("There are " numberOfStaff " many users. " numberOfStaffOnline " of which are online", i, s)
} 
