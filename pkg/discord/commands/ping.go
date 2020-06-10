package commands

import (
	"github.com/kcraley/octane-go/pkg/command"

	"github.com/bwmarrin/discordgo"
)

// PingCmd is the command for sending a ping message back to the user
var PingCmd = &command.Command{
	Name:        "ping",
	Description: "Responds with `pong!`",
	Usage:       "ping",
	Example:     "ping",
	Handler:     pingCmdFunc,
}

// pingCmdFunc is the handler function for the ping command
func pingCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	if _, err := s.ChannelMessageSend(m.ChannelID, "pong!"); err != nil {
		return err
	}
	return nil
}
