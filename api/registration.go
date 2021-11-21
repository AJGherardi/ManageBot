package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

// StartCommandHandler Registers all commands and begins command routing
func (c *Connection) StartCommandHandler(standaloneCommands []StandaloneCommand, parentCommands []ParentCommand, guildID string) {
	// Register all standalone commands
	registerStandaloneCommands(c, standaloneCommands, guildID)
	// Register all parent commands
	registerParentCommands(c, parentCommands, guildID)
	// Make handler
	handle := func(s *dgo.Session, i *dgo.InteractionCreate) {
		// Makes a response
		responseData := &dgo.InteractionApplicationCommandResponseData{
			TTS:     false,
			Content: "Please wait",
		}
		// Sends the initial response
		s.InteractionRespond(i.Interaction, &dgo.InteractionResponse{
			Type: dgo.InteractionResponseChannelMessage,
			Data: responseData,
		})
		// Check perms
		var authorized bool
		for _, roleID := range i.Interaction.Member.Roles {
			role, _ := s.State.Role(i.GuildID, roleID)
			permitted := (role.Permissions & dgo.PermissionAdministrator) == dgo.PermissionAdministrator
			if permitted {
				authorized = true
				break
			}
		}
		// Remove initial response
		s.InteractionResponseDelete("", i.Interaction)
		// Check if authorized
		if authorized == false {
			channel := c.GetChannel(i.ChannelID)
			channel.SendMessage("Not Authorized")
			return
		}
		// Route to application command handler if standalone command
		routeStandaloneCommand(standaloneCommands, i, *c)
		// Route to application command handler if parent command
		routeParentCommand(parentCommands, i, *c)
	}
	// Register handler
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

func registerStandaloneCommands(c *Connection, standaloneCommands []StandaloneCommand, guildID string) {
	for _, standaloneCommand := range standaloneCommands {
		// Get command signature
		standaloneCommandSignature := standaloneCommand.Regester(*c)
		// Register the command
		c.client.ApplicationCommandCreate(
			"",
			&dgo.ApplicationCommand{
				Name:        standaloneCommandSignature.Name,
				Description: standaloneCommandSignature.Description,
				Options:     convertToParamOptions(standaloneCommandSignature.Parms),
			},
			"",
		)
	}
}

func registerParentCommands(c *Connection, parentCommands []ParentCommand, guildID string) {
	for _, parentCommand := range parentCommands {
		// Get parent signature
		parentCommandSignature := parentCommand.Regester(*c)
		// Get subcommand signatures
		subcommandSignatures := []SubcommandSignature{}
		for _, subcommand := range parentCommand.Subcommands() {
			subcommandSignature := subcommand.Regester(*c)
			subcommandSignatures = append(subcommandSignatures, subcommandSignature)
		}
		// Register the command
		c.client.ApplicationCommandCreate(
			"",
			&dgo.ApplicationCommand{
				Name:        parentCommandSignature.Name,
				Description: parentCommandSignature.Description,
				Options:     convertToSubcommandOptions(subcommandSignatures),
			},
			guildID,
		)
	}
}

func convertToParamOptions(params []ParmSignature) []*dgo.ApplicationCommandOption {
	options := []*dgo.ApplicationCommandOption{}
	for _, paramSignature := range params {
		options = append(options, paramSignature.Build())
	}
	return options
}

func convertToSubcommandOptions(subcommands []SubcommandSignature) []*dgo.ApplicationCommandOption {
	subcommandOptions := []*dgo.ApplicationCommandOption{}
	for _, subcommandSignature := range subcommands {
		subcommandOptions = append(subcommandOptions, &dgo.ApplicationCommandOption{
			Name:        subcommandSignature.Name,
			Description: subcommandSignature.Description,
			Type:        dgo.ApplicationCommandOptionSubCommand,
			Options:     convertToParamOptions(subcommandSignature.Parms),
		})
	}
	return subcommandOptions
}

func deleteAllCommands(client *dgo.Session, guildID string) {
	apps, _ := client.Applications()
	for _, app := range apps {
		cmds, _ := client.ApplicationCommands(app.ID, guildID)
		for _, cmd := range cmds {
			client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
		}
	}
	cmds, _ := client.ApplicationCommands("", guildID)
	for _, cmd := range cmds {
		client.ApplicationCommandDelete(cmd.ApplicationID, cmd.ID, guildID)
	}
}
