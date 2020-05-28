package cmd

import (
	"os"

	"github.com/kcraley/octane-go/pkg/configuration"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const appName = "octane"

var (
	// Application configuration
	config *configuration.Configuration

	// Root CLI command
	rootCmd = &cobra.Command{
		Use:   "octane",
		Short: "a simple Discord administration bot",
		Long: `a simple Discord administration bot

Manages Discord servers using a chatops style of commands reminiscent
of old Battle.net days.  Enjoy managing the entire server from a single
chatops pane of glass.`,
	}
)

func init() {
	// Initialize
	cobra.OnInitialize(initConfig)

	// Create a new application configuration
	config = configuration.New()

	// Root level flags for the CLI
	rootCmd.PersistentFlags().StringVarP(&config.Token, "token", "t", "", "the Discord API token used to connect")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "enables additional verbose output for troubleshooting")
}

func initConfig() {
	// Marshal configuration from environment variables
	err := envconfig.Process(appName, &config)
	if err != nil {
		log.Fatalf("Failed parsing configuration from environment variables: %v", err)
	}

	// Configure logging
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	if config.Verbose {
		log.SetLevel(log.DebugLevel)
	}
}

// Execute is the main entrypoint for the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
