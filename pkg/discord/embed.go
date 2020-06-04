package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kcraley/octane-go/pkg/command"
)

var embedCmd = &command.Command{
	Name:        "embed",
	Description: "send an embedded test message",
	Usage:       "embed",
	Example:     "embed",
	Handler:     embedCmdFunc,
}

// embedCmdFunc is the handler function for the embed command
func embedCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "embed",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Message",
				Value:  "```This is a test embedded message.```",
				Inline: false,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}
