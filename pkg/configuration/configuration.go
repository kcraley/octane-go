package configuration

const (
	// DefaultPrefix is the default prefix which the bot listens to
	DefaultPrefix = "!octane"
)

// Configuration contains the application configuration settings
type Configuration struct {
	Prefix  string ``
	Token   string ``
	Verbose bool   ``
}

// New creates and returns a new empty Configuration
func New() *Configuration {
	return &Configuration{}
}
