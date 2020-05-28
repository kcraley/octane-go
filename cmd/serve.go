package cmd

import (
	"os"
	"os/signal"
	"syscall"

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
}

// serveCmdFunc is the entrypoint for `octane serve`
func serveCmdFunc(cmd *cobra.Command, args []string) error {
	// Create a new client with the configuration token
	client, err := discord.NewClient(config.Token)
	if err != nil {
		log.Errorf("Error creating client: %v", err)
		return err
	}
	log.WithFields(log.Fields{
		"sessionId": client.State.SessionID,
		"user":      client.State.User,
		"guilds":    client.State.Guilds,
		"channels":  client.State.PrivateChannels,
	}).Infof("Created a new client client")

	// TODO: Execute real logic here

	// Listen and trap os signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sigChan

	// Cleanup
	log.Info("Recieved signal, shutting down...")
	client.Close()
	return nil
}
