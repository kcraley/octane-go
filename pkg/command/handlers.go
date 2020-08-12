package command

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
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
		log.WithFields(log.Fields{
			"prefix":  r.Prefix,
			"user":    m.Author.Username,
			"guild":   m.GuildID,
			"channel": m.GuildID,
			"content": m.Content,
		}).Info("Prefix triggered")
		// Clean the message content to pull the command
		content := strings.TrimPrefix(m.Content, r.Prefix)
		content = strings.TrimSpace(content)
		userCommands := strings.Split(content, " ")

		if len(userCommands) > 0 && userCommands[0] != "" {
			if command, err := r.GetCommand(userCommands[0]); err != nil {
				// TODO: refactor how the help command is registered
				// and triggered.  This will allow for a better way
				// to provide the user with help/usage messages
				s.ChannelMessageSend(m.ChannelID, "Sending help now...")
			} else {
				if command.HasSubcommands() && len(userCommands) > 1 {
					final, args := command.GetFinalCommand(userCommands[1:])
					final.ParseFlags(args)
					final.Trigger(s, m.Message)
				} else {
					if command.IsRunnable() {
						command.Trigger(s, m.Message)
					} else {
						// TODO: refactor how the help command is registered
						// and triggered.  This will allow for a better way
						// to provide the user with help/usage messages
						s.ChannelMessageSend(m.ChannelID, "Sending help now...")
					}
				}
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, randomDefaultMessage())
		}
	}
}

// randomDefaultMessage returns a random default message
func randomDefaultMessage() string {
	defaultMessages := []string{
		"You rang?",
		"What do you want?",
		"Hi!",
		"I was sleeping...",
		"Bruh",
		"At your command!",
		"Goliath online",
		"Executor?",
		"I haven't got all day...",
		"Make up your mind!",
		"This is Jimmy",
		"Recieving transmission",
		"Howdy!",
		"You got my attention",
		"Wanna turn up the heat?",
		"Call the shot",
		"Go ahead, TACCOM",
		"Jacked up and good to go",
		"Oh my god, he's whacked!",
		"Yes sir?",
		"I read you",
		"Reportin' for duty",
		"Ready to roll out",
	}

	entry := rand.Intn(len(defaultMessages) - 1)
	return defaultMessages[entry]
}
