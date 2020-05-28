package version

import (
	"fmt"
	"runtime"
	"time"
)

const (
	stringFormat = "{BuildDate: %s, GitVersion: %s, GoOSArch: %s, GoVersion: %s}"
)

// Version variable containing all build information
var (
	buildDate  = time.Now().String()
	gitVersion = ""

	Version = buildInformation{
		BuildDate:  buildDate,
		GitVersion: gitVersion,
		GoOSArch:   fmt.Sprintf("%s, %s", runtime.GOOS, runtime.GOARCH),
		GoVersion:  runtime.Version(),
	}
)

// buildInformation holds all information about the build
type buildInformation struct {
	BuildDate  string
	GitVersion string
	GoOSArch   string
	GoVersion  string
}

// String returns a string with all the
func (b *buildInformation) String() string {
	return fmt.Sprintf(
		stringFormat,
		b.BuildDate,
		b.GitVersion,
		b.GoOSArch,
		b.GoVersion,
	)
}
