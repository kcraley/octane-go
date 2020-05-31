package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kcraley/octane-go/pkg/command"
)

var pingCmd = &command.Command{
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
