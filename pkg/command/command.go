package command

// Command represents a single command which can
// be triggered from end users
type Command struct {
	Name        string
	Description string
	Usage       string
	Example     string
	Flags       []string
	SubCommands []*Command
	Handler     interface{}
}

// Trigger executes the Command
func (c *Command) Trigger() error {
	return nil
}
