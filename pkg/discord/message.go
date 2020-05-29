package discord

import (
	"github.com/bwmarrin/discordgo"
)

// handleMessageCreate handles message events
func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
