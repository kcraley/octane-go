package discord

import "github.com/bwmarrin/discordgo"

// NewClient Creates and returns a new discord client
func NewClient(token string) (*discordgo.Session, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return client, nil
}
