package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kcraley/octane-go/pkg/command"
)

// WarzoneCmd is the bot command to lookup Call of Duty Warzone players
var WarzoneCmd = &command.Command{
	Name:        "warzone",
	Description: "lookup a Call of Duty player",
	Usage:       "warzone",
	Example:     "warzone <player>",
	Handler:     warzoneCmdFunc,
}

// warzoneCmdFunc is the handler function for looking up Call of Duty Warzone players
func warzoneCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}
