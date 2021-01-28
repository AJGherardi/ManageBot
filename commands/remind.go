package commands

import (
	"fmt"
	"time"

	"github.com/AJGherardi/ManageBot/types"
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

type remind struct {
	Title string
	Timer *time.Timer
}

var reminders []remind

var remindSubcommands []types.Subcommand = []types.Subcommand{
	{
		Name: "set",
		Callback: func(parms types.SubcommandParms) {
			handleRemindSet(
				parms.Option.Options[0].Value.(string),
				parms.Option.Options[1].Value.(float64),
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "view",
		Callback: func(parms types.SubcommandParms) {
			handleRemindView(
				parms.Interaction,
				parms.Session,
			)
		},
	},
	{
		Name: "delete",
		Callback: func(parms types.SubcommandParms) {
			handleRemindDelete(
				parms.Option.Options[0].Value.(float64),
				parms.Interaction,
				parms.Session,
			)
		},
	},
}

func handleRemindSet(title string, duration float64, i *dgo.InteractionCreate, s *dgo.Session) {
	utils.SendResponse("Reminder set for "+title+" in "+fmt.Sprint(duration)+" min", i, s)
	// Add a timer for the remind
	timer := time.AfterFunc(
		(time.Duration(duration) * time.Minute), func() {
			utils.SendResponse("Reminder for "+title, i, s)
		},
	)
	// Append remind to list
	reminders = append(reminders, remind{
		Title: title,
		Timer: timer,
	})
}

func handleRemindView(i *dgo.InteractionCreate, s *dgo.Session) {
	utils.SendResponse("There are "+fmt.Sprint(len(reminders)), i, s)
	for index, reminder := range reminders {
		utils.SendResponse("There is a Reminder for "+reminder.Title+" at index "+fmt.Sprint(index), i, s)
	}
}

func handleRemindDelete(index float64, i *dgo.InteractionCreate, s *dgo.Session) {
	utils.SendResponse("Deleted reminder", i, s)
	// Remove remind from slice
	reminders = removeRemind(reminders, int(index))
	// Stop timmer for remind
	reminders[int(index)].Timer.Stop()
}

func removeRemind(slice []remind, s int) []remind {
	return append(slice[:s], slice[s+1:]...)
}

// RegesterRemind adds the remind / command
func RegesterRemind(client *dgo.Session, guildID string) types.Handler {
	client.ApplicationCommandCreate(
		"",
		&dgo.ApplicationCommand{
			Name:        "remind",
			Description: "Manage reminders",
			Options: []*dgo.ApplicationCommandOption{
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "set",
					Description: "Set a reminder",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionString,
							Name:        "Title",
							Description: "Title of reminder",
							Required:    true,
						},
						{
							Type:        dgo.ApplicationCommandOptionInteger,
							Name:        "Time",
							Description: "How many min until reminder",
							Required:    true,
						},
					},
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "view",
					Description: "View reminders",
				},
				{
					Type:        dgo.ApplicationCommandOptionSubCommand,
					Name:        "delete",
					Description: "Deletes the reminder at the given index",
					Options: []*dgo.ApplicationCommandOption{
						{
							Type:        dgo.ApplicationCommandOptionInteger,
							Name:        "index",
							Description: "Index of reminder",
							Required:    true,
						},
					},
				},
			},
		},
		guildID,
	)
	// Return Handler
	return types.Handler{
		Name: "remind", Callback: func(i *dgo.InteractionCreate, s *dgo.Session) {
			utils.MatchSubcommand(i, s, remindSubcommands)
		},
	}
}
