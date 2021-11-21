package commands

import (
	"fmt"
	"time"

	"github.com/AJGherardi/ManageBot/api"
)

type remind struct {
	Title string
	Timer *time.Timer
}

var reminders []remind

type RemindHandler struct{}

func (h *RemindHandler) Name() string {
	return "remind"
}

func (h *RemindHandler) Subcommands() []api.Subcommand {
	return []api.Subcommand{
		&setRemindHandler{},
		&viewRemindHandler{},
		&deleteRemindHandler{},
	}
}

func (h *RemindHandler) Regester(c api.Connection) api.ParentCommandSignature {
	return api.MakeParentCommandSignature("remind", "Manage reminders")
}

type setRemindHandler struct{}

func (h *setRemindHandler) Name() string {
	return "set"
}

func (h *setRemindHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Reminder set for " + i.GetStringParm(0) + " in " + fmt.Sprint(i.GetIntParm(1)) + " min")
	// Add a timer for the remind
	timer := time.AfterFunc(
		(time.Duration(i.GetIntParm(1)) * time.Minute), func() {
			channel := c.GetChannel(i.GetChannelID())
			channel.SendEmbedMessage("Reminder for " + i.GetStringParm(0))
		},
	)
	// Append remind to list
	reminders = append(reminders, remind{
		Title: i.GetStringParm(0),
		Timer: timer,
	})
}

func (h *setRemindHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"set", "Set a reminder",
		api.MakeStringParmSignature("Title", "Title of reminder", true),
		api.MakeIntParmSignature("Time", "How many min until reminder", true),
	)
}

type viewRemindHandler struct{}

func (h *viewRemindHandler) Name() string {
	return "view"
}

func (h *viewRemindHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("There is " + fmt.Sprint(len(reminders)))
	for index, reminder := range reminders {
		channel.SendEmbedMessage("There is a Reminder for " + reminder.Title + " at index " + fmt.Sprint(index))
	}
}

func (h *viewRemindHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature("view", "View reminders")
}

type deleteRemindHandler struct{}

func (h *deleteRemindHandler) Name() string {
	return "delete"
}

func (h *deleteRemindHandler) Callback(i api.SubcommandInvocation, c api.Connection) {
	channel := c.GetChannel(i.GetChannelID())
	channel.SendEmbedMessage("Deleted reminder")
	// Stop timmer for remind
	reminders[i.GetIntParm(0)].Timer.Stop()
	// Remove remind from slice
	reminders = removeRemind(reminders, i.GetIntParm(0))
}

func (h *deleteRemindHandler) Regester(c api.Connection) api.SubcommandSignature {
	return api.MakeSubcommandSignature(
		"delete", "Deletes the reminder at the given index",
		api.MakeIntParmSignature("Index", "Index of reminder", true),
	)
}

func removeRemind(slice []remind, s int) []remind {
	return append(slice[:s], slice[s+1:]...)
}
