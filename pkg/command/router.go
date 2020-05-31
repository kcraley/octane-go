package command

import "github.com/bwmarrin/discordgo"

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
func (r *Router) GetCommand(name string) *Command {
	for _, command := range r.Commands {
		if name == command.Name {
			return command
		}
	}
	return nil
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
