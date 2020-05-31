package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kcraley/octane-go/pkg/configuration"
	"github.com/kcraley/octane-go/pkg/discord"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the Discord bot and serves commands",
	Long: `starts the Discord bot and serves commands

Start the bot and connect to the Discord servers.  This will allow
the bot to start handling traffic and commands.`,
	RunE: serveCmdFunc,
}

func init() {
	// Add `serve` subcommand to `octane`
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringVarP(&config.Prefix, "prefix", "p", configuration.DefaultPrefix, "the Discord API token used to connect")
	serveCmd.PersistentFlags().StringVarP(&config.Token, "token", "t", "", "the Discord API token used to connect")
}

// serveCmdFunc is the entrypoint for `octane serve`
func serveCmdFunc(cmd *cobra.Command, args []string) error {
	// Create a new client with the configuration token
	client, err := discord.NewClient(config.Token, config.Prefix)
	if err != nil {
		log.Errorf("Error creating client: %v", err)
		return err
	}
	log.Infof("Created a new client")

	// Initialize all handlers
	client.Router.Initialize(client.Session)

	// Open a websocket connection and listen
	err = client.Session.Open()
	if err != nil {
		log.Errorf("Failed opening a connection to Discord: %v", err)
	}
	defer client.Session.Close()
	log.WithFields(log.Fields{
		"sessionId": client.Session.State.SessionID,
		"user":      client.Session.State.User,
		"guilds":    client.Session.State.Guilds,
		"channels":  client.Session.State.PrivateChannels,
	}).Infof("Connected to Discord")

	// Listen and trap os signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sigChan

	// Cleanup
	log.Info("Recieved signal, shutting down...")
	return nil
}
