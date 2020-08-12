package command

import (
	"errors"
	"flag"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ExecutionHandler represents the function that is to be executed
type ExecutionHandler func(*discordgo.Session, *discordgo.Message) error

// Command represents a single command which can
// be triggered from end users
type Command struct {
	Name        string
	Description string
	Usage       string
	Example     string
	SubCommands []*Command
	Handler     ExecutionHandler
	flagSet     *flag.FlagSet
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
func (c *Command) GetFinalCommand(commands []string) (*Command, []string) {
	for _, subcommand := range c.SubCommands {
		if strings.HasPrefix(commands[0], "-") {
			continue
		} else if subcommand.Name == commands[0] {
			return subcommand.GetFinalCommand(commands[1:])
		}
	}
	return c, commands
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

// Flags returns the FlagSet for the current command.  If an
// existing FlagSet does not exist, a new one will be created
// and returned.
func (c *Command) Flags() *flag.FlagSet {
	if c.flagSet == nil {
		c.flagSet = flag.NewFlagSet(c.Name, flag.ContinueOnError)
	}
	return c.flagSet
}

// ParseFlags reads in all user provided input and updates
// the commands flags
func (c *Command) ParseFlags(input []string) error {
	err := c.Flags().Parse(input)
	if err != nil {
		return err
	}
	return nil
}

// Trigger executes the Command
func (c *Command) Trigger(s *discordgo.Session, m *discordgo.Message) error {
	if err := c.Handler(s, m); err != nil {
		return err
	}
	return nil
}
