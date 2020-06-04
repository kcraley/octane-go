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
	// Create a new trivia Game
	game := trivia.NewGame()

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
		questionMessage, err := s.ChannelMessageSend(m.ChannelID, question.Question)
		if err != nil {
			return err
		}

		// Set timeout and retrieve players' answers
		timeout := time.Now().Add(15 * time.Second).Unix()
		for {
			now := time.Now().Unix()
			if now >= timeout {
				break
			}

			answers, err := s.ChannelMessages(m.ChannelID, 100, "", questionMessage.ID, "")
			if err != nil {
				return err
			}
			for _, answer := range answers {
				// Validate player exists and if not, create one
				if !game.ContainsPlayer(answer.Author.Username) {
					game.AddPlayer(trivia.NewPlayer(answer.Author.Username))
				}

				// Validate the user's response against the answer
				if answer.Content == question.Answer {
					player := game.GetPlayer(answer.Author.Username)
					player.AddCorrectAnswer()
					player.AddScore(question.Value)
					break
				}
			}
		}

		// Send the answer
		if _, err := s.ChannelMessageSend(m.ChannelID, question.Answer); err != nil {
			return err
		}
	}

	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Game Over!",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Trivia Score:",
				Value: fmt.Sprintf("Players: %s, Score: %d, Correct Answers: %d",
					game.Players[0].Username,
					game.Players[0].Score,
					game.Players[0].CorrectAnswers,
				),
				Inline: true,
			},
		},
	}); err != nil {
		return err
	}
	return nil
}
