package discord

import (
	"github.com/kcraley/octane-go/pkg/command"

	"github.com/bwmarrin/discordgo"
)

// Client is a custom Discord client containing the
// session and handlers.
type Client struct {
	Session *discordgo.Session
	Router  *command.Router
}

// NewClient Creates and returns a new discord client
func NewClient(token string, prefix string) (*Client, error) {
	// Create a new discordgo session
	dcSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	// Create a new command Router
	router := command.NewRouter(prefix)

	// Create and return a new instance of a Client
	client := &Client{
		Session: dcSession,
		Router:  router,
	}
	return client, nil
}
