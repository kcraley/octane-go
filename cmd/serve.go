package cmd

import "github.com/spf13/cobra"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the Discord bot and serves commands",
	Long: `starts the Discord bot and serves commands

Start the bot and connect to the Discord servers.  This will allow
the bot to start handling traffic and commands.`,
	RunE: serveCmdFunc,
}

// serveCmdFunc is the entrypoint for `octane serve`
func serveCmdFunc(cmd *cobra.Command, args []string) error {
	return nil
}
