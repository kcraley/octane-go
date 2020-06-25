package command

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Router represents a command router for organizing
// and executing commands registered.
type Router struct {
	Prefix   string
	Commands []*Command
}

// NewRouter creates and returns a new Router
func NewRouter(prefix string) *Router {
	return &Router{
		Prefix: prefix,
	}
}

// GetCommand returns an existing Command given the name
func (r *Router) GetCommand(name string) (*Command, error) {
	for _, command := range r.Commands {
		if name == command.Name {
			return command, nil
		}
	}
	return nil, errors.New("Command not registered")
}

// ContainsCommand verifies that the Router has a command regsitered
// with the given name
func (r *Router) ContainsCommand(name string) bool {
	for _, command := range r.Commands {
		if name == command.Name {
			return true
		}
	}
	return false
}

// Initialize adds all handler functions to a *discordgo.Session
// All handlers should be registered here so they can be triggered
func (r *Router) Initialize(session *discordgo.Session) {
	session.AddHandler(r.handleReady)
	session.AddHandler(r.handleMessageCreate)
}

// RegisterCommand adds a new Command to the Router
// which can be triggered
func (r *Router) RegisterCommand(cmd *Command) {
	r.Commands = append(r.Commands, cmd)
}

// RegisterHelpCommand adds the help command for usage information
// on all registered commands
func (r *Router) RegisterHelpCommand() {
	var (
		usageOutput string

		helpCmdFunc = func(s *discordgo.Session, m *discordgo.Message) error {
			for _, command := range r.Commands {
				usageOutput += fmt.Sprintf("%s\t%s\n", command.Name, command.Description)
			}
			if _, err := s.ChannelMessageSend(m.ChannelID, usageOutput); err != nil {
				return err
			}
			return nil
		}

		helpCmd = &Command{
			Name:        "help",
			Description: "Provides usage and additional information of registered commands",
			Usage:       "help",
			Example:     "help [subcommand]",
			Handler:     helpCmdFunc,
		}
	)

	r.RegisterCommand(helpCmd)
}
