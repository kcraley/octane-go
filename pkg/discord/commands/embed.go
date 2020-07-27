package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kcraley/octane-go/pkg/command"
)

// EmbedCmd is a test chat command for sending embedded messages
var EmbedCmd = &command.Command{
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
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://github.com/kcraley/octane-go",
			Name:    "Octane",
			IconURL: "https://raw.githubusercontent.com/kcraley/octane-go/callofduty/docs/images/yellow-fire-button-medium.png",
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://raw.githubusercontent.com/kcraley/octane-go/callofduty/docs/images/yellow-fire-button-medium.png",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Line messgae",
				Value:  "This is a message on a whole line.",
				Inline: false,
			},
			{
				Name:   "Inline Message 1",
				Value:  "This is the first inline message",
				Inline: true,
			},
			{
				Name:   "Inline Message 2",
				Value:  "This is the second inline message",
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: "https://raw.githubusercontent.com/kcraley/octane-go/callofduty/docs/images/yellow-fire-button-medium.png",
			Text:    "Octane Discord Bot",
		},
	}); err != nil {
		return err
	}

	return nil
}
