package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// handleReady is the handler for ready events
func (r *Router) handleReady(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateListeningStatus("your commands")
}

// handleMessageCreate is the handler for message create events
func (r *Router) handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages generated by bots or the bot itself
	if m.Author.Bot || m.Author.ID == s.State.User.ID {
		return
	}

	// Check for router prefix
	if strings.HasPrefix(m.Content, r.Prefix) {
		// Clean the message content to pull the command
		content := strings.TrimPrefix(m.Content, r.Prefix)
		content = strings.TrimSpace(content)

		for _, command := range r.Commands {
			if content == command.Name {
				command.Trigger()
			}
		}
	}
}
