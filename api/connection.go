package api

import (
	"github.com/AJGherardi/ManageBot/utils"
	dgo "github.com/bwmarrin/discordgo"
)

// Connection Wraps the discord api for easy use
type Connection struct {
	client *dgo.Session
}

// ConnectToDiscord Opens a new connection to discord
func ConnectToDiscord(botToken, guildID string) Connection {
	// Creates a new client object
	client, _ := dgo.New("Bot " + botToken)
	// Set intents
	client.Identify.Intents = dgo.MakeIntent(
		dgo.IntentsAllWithoutPrivileged |
			dgo.IntentsGuildPresences |
			dgo.IntentsGuildMembers,
	)
	// Opens the connection
	client.Open()
	// Remove all commands
	deleteAllCommands(client, guildID)
	return Connection{client: client}
}

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
			utils.SendResponse("Not authorized", i, s)
			return
		}
		// Route to appliction command handler if standalone command
		routeStandaloneCommand(standaloneCommands, i, s)
		// Route to application command handler if parent command
		routeParentCommand(parentCommands, i, s)
	}
	// Regester handler
	c.client.AddHandler(handle)
}

func routeParentCommand(parentCommands []ParentCommand, i *dgo.InteractionCreate, s *dgo.Session) {
	for _, parentCommand := range parentCommands {
		// Match parent command
		if parentCommand.Name() == i.Interaction.Data.Name {
			// Match subcommand
			for _, subcommand := range parentCommand.Subcommands() {
				if subcommand.Name() == i.Interaction.Data.Options[0].Name {
					subcommand.Callback(i, s)
				}
			}
		}
	}
}

func routeStandaloneCommand(standaloneCommands []StandaloneCommand, i *dgo.InteractionCreate, s *dgo.Session) {
	for _, standaloneCommand := range standaloneCommands {
		if standaloneCommand.Name() == i.Interaction.Data.Name {
			standaloneCommand.Callback(i, s)
		}
	}
}

func regesterStandaloneCommands(c *Connection, standaloneCommands []StandaloneCommand, guildID string) {
	for _, standaloneCommand := range standaloneCommands {
		// Get command signature
		standaloneCommandSinginture := standaloneCommand.Regester()
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
		parentCommandSinginture := parentCommand.Regester()
		// Get subcommand singintures
		subcommandSingintures := []SubcommandSinginture{}
		for _, subcommand := range parentCommand.Subcommands() {
			subcommandSinginture := subcommand.Regester()
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
		options = append(options, &dgo.ApplicationCommandOption{
			Name:        parmSinginture.Name,
			Description: parmSinginture.Description,
			Required:    parmSinginture.Required,
			Type:        dgo.ApplicationCommandOptionType(parmSinginture.Type),
		})
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
