package commands

import (
	"fmt"
	"time"

	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// HandleRemind handles a remind command duration is in min
func HandleRemind(title string, duration float64, i *dgo.InteractionCreate, s *dgo.Session) {
	utils.SendResponse("Reminder set for "+title+" in "+fmt.Sprint(duration)+" min", i, s)
	time.AfterFunc(
		(time.Duration(duration) * time.Minute), func() {
			utils.SendResponse("Reminder for "+title, i, s)
		},
	)
}
