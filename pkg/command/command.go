package command

import "github.com/bwmarrin/discordgo"

// ExecutionHandler represents the function that is to be executed
type ExecutionHandler func(*discordgo.Session, *discordgo.Message) error

// Command represents a single command which can
// be triggered from end users
type Command struct {
	Name        string
	Description string
	Usage       string
	Example     string
	Flags       []string
	SubCommands []*Command
	Handler     ExecutionHandler
}

// Trigger executes the Command
func (c *Command) Trigger(s *discordgo.Session, m *discordgo.Message) error {
	if err := c.Handler(s, m); err != nil {
		return err
	}
	return nil
}
