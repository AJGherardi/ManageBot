package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

// StartCommandHandler Regesters all commands and begines command routing
func (c *Connection) StartCommandHandler(standaloneCommands []StandaloneCommand, parentCommands []ParentCommand, guildID string) {
	// Regester all standalone commands
	regesterStandaloneCommands(c, standaloneCommands, guildID)
	// Regester all parent commands
	regesterParentCommands(c, parentCommands, guildID)
	// Make handler
	handle := func(s *dgo.Session, i *dgo.InteractionCreate) {
		// Makes a reaponse
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: "Please wait",
		}
		// Sends the inital response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
		// Chack perms
		var authorized bool
		for _, roleID := range i.Interaction.Member.Roles {
			role, _ := s.State.Role(i.GuildID, roleID)
			permited := (role.Permissions & dgo.PermissionAdministrator) == dgo.PermissionAdministrator
			if permited {
				authorized = true
				break
			}
		}
		// Remove initial reaponse
		s.InteractionResponseDelete("", i.Interaction)
		// Check if authorized
		if authorized == false {
			channel := c.GetChannel(i.ChannelID)
			channel.SendMessage("Not Authorized")
			return
		}
		// Route to appliction command handler if standalone command
		routeStandaloneCommand(standaloneCommands, i, *c)
		// Route to application command handler if parent command
		routeParentCommand(parentCommands, i, *c)
	}
	// Regester handler
	c.client.AddHandler(handle)
}

func routeParentCommand(parentCommands []ParentCommand, i *dgo.InteractionCreate, c Connection) {
	for _, parentCommand := range parentCommands {
		// Match parent command
		if parentCommand.Name() == i.Interaction.Data.Name {
			// Match subcommand
			for _, subcommand := range parentCommand.Subcommands() {
				if subcommand.Name() == i.Interaction.Data.Options[0].Name {
					subcommand.Callback(SubcommandInvocation{invocation: invocation{i: i}}, c)
				}
			}
		}
	}
}

func routeStandaloneCommand(standaloneCommands []StandaloneCommand, i *dgo.InteractionCreate, c Connection) {
	for _, standaloneCommand := range standaloneCommands {
		if standaloneCommand.Name() == i.Interaction.Data.Name {
			standaloneCommand.Callback(StandaloneCommandInvocation{invocation: invocation{i: i}}, c)
		}
	}
}

func regesterStandaloneCommands(c *Connection, standaloneCommands []StandaloneCommand, guildID string) {
	for _, standaloneCommand := range standaloneCommands {
		// Get command signature
		standaloneCommandSinginture := standaloneCommand.Regester(*c)
		// Regester the command
		c.client.ApplicationCommandCreate(
			"",
			&dgo.ApplicationCommand{
				Name:        standaloneCommandSinginture.Name,
				Description: standaloneCommandSinginture.Description,
				Options:     convertToParmOptions(standaloneCommandSinginture.Parms),
			},
			guildID,
		)
	}
}

func regesterParentCommands(c *Connection, parentCommands []ParentCommand, guildID string) {
	for _, parentCommand := range parentCommands {
		// Get parent signature
		parentCommandSinginture := parentCommand.Regester(*c)
		// Get subcommand singintures
		subcommandSingintures := []SubcommandSinginture{}
		for _, subcommand := range parentCommand.Subcommands() {
			subcommandSinginture := subcommand.Regester(*c)
			subcommandSingintures = append(subcommandSingintures, subcommandSinginture)
		}
		// Regester the command
		c.client.ApplicationCommandCreate(
			"",
			&dgo.ApplicationCommand{
				Name:        parentCommandSinginture.Name,
				Description: parentCommandSinginture.Description,
				Options:     convertToSubcommandOptions(subcommandSingintures),
			},
			guildID,
		)
	}
}

func convertToParmOptions(parms []ParmSinginture) []*dgo.ApplicationCommandOption {
	options := []*dgo.ApplicationCommandOption{}
	for _, parmSinginture := range parms {
		options = append(options, parmSinginture.Build())
	}
	return options
}

func convertToSubcommandOptions(subcommands []SubcommandSinginture) []*dgo.ApplicationCommandOption {
	subcommandOptions := []*dgo.ApplicationCommandOption{}
	for _, subcommandSinginture := range subcommands {
		subcommandOptions = append(subcommandOptions, &dgo.ApplicationCommandOption{
			Name:        subcommandSinginture.Name,
			Description: subcommandSinginture.Description,
			Type:        dgo.ApplicationCommandOptionSubCommand,
			Options:     convertToParmOptions(subcommandSinginture.Parms),
		})
	}
	return subcommandOptions
}

// func deleteAllCommands(client *dgo.Session, guildID string) {
// 	apps, _ := client.Applications()
// 	for _, app := range apps {
// 		cmds, _ := client.ApplicationCommands(app.ID, guildID)
// 		for _, cmd := range cmds {
// 			client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
// 		}
// 	}
// 	cmds, _ := client.ApplicationCommands("", guildID)
// 	for _, cmd := range cmds {
// 		client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
// 	}
// }
