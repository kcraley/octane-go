package discord

import "github.com/bwmarrin/discordgo"

// Client is a custom Discord client containing the
// session and handlers.
type Client struct {
	Session  *discordgo.Session
	Handlers []interface{}
}

// NewClient Creates and returns a new discord client
func NewClient(token string) (*Client, error) {
	dcSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Session: dcSession,
	}
	return client, nil
}
