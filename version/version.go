package version

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

const (
	stringFormat = "{BuildDate: %s, DiscordGo: %s, GitVersion: %s, GoOSArch: %s, GoVersion: %s}"
)

// Version variable containing all build information
var (
	buildDate  = ""
	gitVersion = ""

	Version = buildInformation{
		BuildDate:  buildDate,
		DiscordGo:  discordgo.VERSION,
		GitVersion: gitVersion,
		GoOSArch:   fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		GoVersion:  runtime.Version(),
	}
)

// buildInformation holds all information about the build
type buildInformation struct {
	BuildDate  string
	DiscordGo  string
	GitVersion string
	GoOSArch   string
	GoVersion  string
}

// String returns a string with all the
func (b *buildInformation) String() string {
	return fmt.Sprintf(
		stringFormat,
		b.BuildDate,
		b.DiscordGo,
		b.GitVersion,
		b.GoOSArch,
		b.GoVersion,
	)
}
