package command

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// ExecutionHandler represents the function that is to be executed
// type ExecutionHandler func(*discordgo.Session, *discordgo.Message) error

// Command represents a single command which can
// be triggered from end users
type Command struct {
	Name        string
	Description string
	Usage       string
	Example     string
	Flags       []string
	SubCommands []*Command
	Handler     func(*discordgo.Session, *discordgo.Message) error
}

// HasSubcommands verifies if the command has additional subcommands
func (c *Command) HasSubcommands() bool {
	if len(c.SubCommands) > 0 {
		return true
	}
	return false
}

// IsRunnable validates the the command actually has a Handler
func (c *Command) IsRunnable() bool {
	return c.Handler != nil
}

// GetFinalCommand recursively iterates through a Command's
// subcommands and returns the final command for a list of
// commands
func (c *Command) GetFinalCommand(commands []string) (*Command, error) {
	for _, subcommand := range c.SubCommands {
		if subcommand.Name == commands[0] && subcommand.HasSubcommands() {
			return subcommand.GetFinalCommand(commands[1:])
		}
		return subcommand, nil
	}
	return nil, errors.New("Subcommand does not exist")
}

// GetSubCommand iterates through a Command's subcommands
// and returns a command if it matches a given name
func (c *Command) GetSubCommand(name string) (*Command, error) {
	for _, subcommand := range c.SubCommands {
		if subcommand.Name == name {
			return subcommand, nil
		}
	}
	return nil, errors.New("Subcommand does not exist")
}

// Trigger executes the Command
func (c *Command) Trigger(s *discordgo.Session, m *discordgo.Message) error {
	if err := c.Handler(s, m); err != nil {
		return err
	}
	return nil
}
