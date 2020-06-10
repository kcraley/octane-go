package commands

import (
	"strings"

	"github.com/kcraley/octane-go/pkg/command"
	"github.com/kcraley/octane-go/version"

	"github.com/bwmarrin/discordgo"
)

// VersionCmd is the command to reply with Octane build information
var VersionCmd = &command.Command{
	Name:        "version",
	Description: "prints the version of Octane",
	Usage:       "version",
	Example:     "version",
	Handler:     versionCmdFunc,
}

// versionCmdFunc is the handler function for the ping command
func versionCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	versionString := version.Version.String()
	versionString = strings.Trim(versionString, "{")
	versionString = strings.Trim(versionString, "}")
	if _, err := s.ChannelMessageSend(m.ChannelID, versionString); err != nil {
		return err
	}
	return nil
}
