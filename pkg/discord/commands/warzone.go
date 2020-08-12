package commands

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
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
	SubCommands: []*command.Command{
		WarzoneProfileCmd,
		WarzoneStatsCmd,
	},
}

var (
	// WarzoneProfileCmd is a subcommand of warzone which is responsible
	// for looking up player profiles
	WarzoneProfileCmd = &command.Command{
		Name:        "profile",
		Description: "lookup a Call of Duty player profile",
		Usage:       "warzone profile",
		Example:     "warzone profile --username CoaXeD#1535 --platform battle",
		Handler:     WarzoneProfileCmdFunc,
	}
	// Variables used for parsing flags for the profile subcommand
	wzProfPlatform string
	wzProfUsername string
)

var (
	// WarzoneStatsCmd is a subcommand of warzone which is responsible
	// for getting a player's overall statistics
	WarzoneStatsCmd = &command.Command{
		Name:        "stats",
		Description: "retrieves a player's statistics",
		Usage:       "warzone stats",
		Example:     "warzone stats --username CoaXeD#1535 --platform battle",
		Handler:     WarzoneStatsCmdFunc,
	}
)

func init() {
	// TODO: These flags should be set at the warzone subcommand
	// layer and passed down to additional subcommands.  Between
	// our `warzone` subcommands the same variables are being used.
	//
	// Set flags for `warzone profile` subcommand
	WarzoneProfileCmd.Flags().StringVar(&wzProfPlatform, "platform", "battle", "The platform where the user's account is registered.")
	WarzoneProfileCmd.Flags().StringVar(&wzProfUsername, "username", "", "The username to lookup in Call of Duty's APIs.")

	// Set flags for `warzone stats` subcommand
	WarzoneStatsCmd.Flags().StringVar(&wzProfPlatform, "platform", "battle", "The platform where the user's account is registered.")
	WarzoneStatsCmd.Flags().StringVar(&wzProfUsername, "username", "", "The username to lookup in Call of Duty's APIs.")
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

		escapedProfile := url.QueryEscape(wzProfUsername)
		playerProfile, err := codClient.GetProfile("mw", wzProfPlatform, escapedProfile, "wz")
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

// WarzoneStatsCmdFunc is tha handler function for retrieving a Call of Duty Warzone player statistics
func WarzoneStatsCmdFunc(s *discordgo.Session, m *discordgo.Message) error {
	if os.Getenv("COD_USERNAME") != "" && os.Getenv("COD_PASSWORD") != "" {
		// Create a client to interact with the Call of Duty APIs
		codClient := callofduty.NewClient()
		err := codClient.Login(os.Getenv("COD_USERNAME"), os.Getenv("COD_PASSWORD"))
		if err != nil {
			return err
		}

		escapedProfile := url.QueryEscape(strings.ToLower(wzProfUsername))
		brStats, err := codClient.GetBRStats("mw", wzProfPlatform, escapedProfile, "wz")
		if err != nil {
			return err
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Type:        "rich",
			Title:       "Player Statistics - Mode: Battle Royal",
			Description: "Below is the lifetime statistics of a player's profile for Warzone",
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
				URL: "https://www.callofduty.com/content/dam/atvi/callofduty/cod-touchui/kronos/buy/bundle-features/mw-inc-xrk-desktop.jpg",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Username",
					Value:  strings.ToLower(wzProfUsername),
					Inline: true,
				},
				{
					Name:   "Platform",
					Value:  wzProfPlatform,
					Inline: true,
				},
				{
					Name:   "Games Played",
					Value:  strconv.Itoa(int(brStats.GamesPlayed)),
					Inline: true,
				},
				{
					Name:   "Wins",
					Value:  strconv.Itoa(int(brStats.Wins)),
					Inline: true,
				},
				{
					Name:   "Top Five",
					Value:  strconv.Itoa(int(brStats.TopFive)),
					Inline: true,
				},
				{
					Name:   "Top Ten",
					Value:  strconv.Itoa(int(brStats.TopTen)),
					Inline: true,
				},
				{
					Name:   "Top Twenty Five",
					Value:  strconv.Itoa(int(brStats.TopTwentyFive)),
					Inline: true,
				},
				{
					Name:   "Kills",
					Value:  strconv.Itoa(int(brStats.Kills)),
					Inline: true,
				},
				{
					Name:   "Deaths",
					Value:  strconv.Itoa(int(brStats.Deaths)),
					Inline: true,
				},
				{
					Name:   "K/D Ratio",
					Value:  fmt.Sprintf("%f%%", brStats.KdRatio*100/100),
					Inline: true,
				},
				{
					Name:   "Downs",
					Value:  strconv.Itoa(int(brStats.Downs)),
					Inline: true,
				},
				{
					Name:   "Score",
					Value:  strconv.Itoa(int(brStats.Score)),
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
