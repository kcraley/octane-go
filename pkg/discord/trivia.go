package discord

import (
	"fmt"
	"strings"
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
	// Create a new trivia Game and fetch querstions
	game := trivia.NewGame()
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
		// NOTE: This loop is most likely being rate limited
		timeout := time.Now().Add(15 * time.Second).Unix()
		for {
			now := time.Now().Unix()
			if now >= timeout {
				break
			}

			// Constantly get messages since the question was answered
			answers, err := s.ChannelMessages(m.ChannelID, 100, "", questionMessage.ID, "")
			if err != nil {
				return err
			}

			// Iterate through each answer and handle it
			for _, answer := range answers {
				// Create a new Player in the Game if they don't exist
				if !game.ContainsPlayer(answer.Author.Username) {
					game.AddPlayer(trivia.NewPlayer(answer.Author.Username))
				}

				// Clean the answers
				answer.Content = strings.ToLower(answer.Content)
				answer.Content = strings.TrimSpace(answer.Content)

				// Validate the user's response against the answer
				// TODO: check if the user's answer is a partial match
				// and give partial points.
				if answer.Content == strings.ToLower(question.Answer) {
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
				Name:   "Trivia Score:",
				Value:  game.SprintScoreboard(),
				Inline: true,
			},
		},
	}); err != nil {
		return err
	}
	return nil
}
