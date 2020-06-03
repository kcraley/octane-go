package discord

import (
	"fmt"
	"time"

	"github.com/kcraley/octane-go/pkg/command"
	"github.com/kcraley/octane-go/pkg/trivia"

	"github.com/bwmarrin/discordgo"
)

var triviaCmd = &command.Command{
	Name:        "trivia",
	Description: "Starts a trivia session",
	Usage:       "trivia",
	Example:     "trivia",
	Handler:     triviaCmdFunc,
}

// triviaCmdFunc is the handler function for the trivia command
func triviaCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	// Retrieve a random trivia question
	questions, err := trivia.GetRandomQuestions(1)
	if err != nil {
		return err
	}

	for _, question := range questions {
		// Inform the players of the category
		if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The category is... %s", question.Category.Title)); err != nil {
			return err
		}

		// Let the players process the category
		time.Sleep(3 * time.Second)

		// Ask the trivia question
		if _, err := s.ChannelMessageSend(m.ChannelID, question.Question); err != nil {
			return err
		}

		// Let the players make guesses
		time.Sleep(10 * time.Second)

		// Send the answer
		if _, err := s.ChannelMessageSend(m.ChannelID, question.Answer); err != nil {
			return err
		}
	}
	return nil
}
