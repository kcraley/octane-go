package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Client is a custom Discord client containing the
// session and handlers.
type Client struct {
	Session  *discordgo.Session
	Handlers []interface{}
}

// NewClient Creates and returns a new discord client
func NewClient(token string) (*Client, error) {
	handlers := make([]interface{}, 0)
	handlers = append(handlers, handleReady)

	dcSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Session:  dcSession,
		Handlers: handlers,
	}
	return client, nil
}

// RegisterAllHandlers iterates through all existing Client
// handlers and adds them to the Session.
func (c *Client) RegisterAllHandlers() error {
	if len(c.Handlers) > 0 {
		for _, handler := range c.Handlers {
			c.Session.AddHandler(handler)
			return nil
		}
	}
	return errors.New("Client does not have any registered handlers")
}

// handleReady handles the Ready event
func handleReady(s *discordgo.Session, ev *discordgo.Ready) {
	s.UpdateListeningStatus("your commands")
}
