package commands

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/kcraley/octane-go/pkg/callofduty"
	"github.com/kcraley/octane-go/pkg/command"

	"github.com/bwmarrin/discordgo"
)

// WarzoneCmd is the bot command to interact with the Call of Duty APIs
var WarzoneCmd = &command.Command{
	Name:        "warzone",
	Description: "lookup a Call of Duty player",
	Usage:       "warzone",
	Example:     "warzone <player>",
	SubCommands: []*command.Command{WarzoneProfileCmd},
}

// WarzoneProfileCmd is a subcommand of warzone which is responsible
// for looking up player profiles
var WarzoneProfileCmd = &command.Command{
	Name:        "profile",
	Description: "lookup a Call of Duty player profile",
	Usage:       "warzone profile",
	Example:     "warzone profile <player>",
	Handler:     WarzoneProfileCmdFunc,
}

// WarzoneProfileCmdFunc is the handler function for looking up Call of Duty Warzone player profiles
func WarzoneProfileCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	// Login to the Call of Duty APIs if credentials are set
	if os.Getenv("COD_USERNAME") != "" && os.Getenv("COD_PASSWORD") != "" {
		// Create a client to interact with the Call of Duty APIs
		codClient := callofduty.NewClient()
		err := codClient.Login(os.Getenv("COD_USERNAME"), os.Getenv("COD_PASSWORD"))
		if err != nil {
			return err
		}

		playerProfile, err := codClient.GetProfile("mw", "battle", "CoaXeD%231353", "wz")
		if err != nil {
			return err
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Type:        "rich",
			Title:       "Player Profile",
			Description: "This is a simple player profile.",
			URL:         "https://www.callofduty.com/",
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       16105282,
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://www.callofduty.com/",
				Name:    "Call of Duty",
				IconURL: "https://www.callofduty.com/content/dam/atvi/callofduty/hub/touch-ui/common/kronos-icon.png",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://modernwarfarediscordbot.com/images/gamemodes/br.png",
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://www.callofduty.com/content/dam/atvi/callofduty/cod-touchui/kronos/buy/bundle-features/mw-inc-wz-desktop.jpg",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Username",
					Value:  playerProfile.Username,
					Inline: true,
				},
				{
					Name:   "Platform",
					Value:  playerProfile.Platform,
					Inline: true,
				},
				{
					Name:   "Level",
					Value:  strconv.Itoa(int(playerProfile.Level)),
					Inline: true,
				},
				{
					Name:   "Total XP",
					Value:  strconv.Itoa(int(playerProfile.TotalXp)),
					Inline: true,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: "https://www.callofduty.com/content/dam/atvi/callofduty/cod-touchui/hub/wtb-portal/touts/battlepass-touts/Call-of-Duty_codpoints-coin.png",
				Text:    "Season Four Reloaded: New Modes and Content Available!",
			},
		})
		return nil
	}

	return errors.New("Unable to connect to Call of Duty APIs")
}
