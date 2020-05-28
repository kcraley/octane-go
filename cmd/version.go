package cmd

import (
	"fmt"

	"github.com/kcraley/octane-go/version"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "prints the version and build information for the binary",
		Long:  `prints the version and build information for the binary`,
		Run:   versionCmdFunc,
	}
)

func init() {
	// Add `version` subcommand to `octane`
	rootCmd.AddCommand(versionCmd)
}

// versionCmdFunc is the entrypoint for `octane version`
func versionCmdFunc(cmd *cobra.Command, args []string) {
	// Print binary version information
	fmt.Println(version.Version.String())
}
