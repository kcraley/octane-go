package configuration

// Configuration contains the application configuration settings
type Configuration struct {
	Token   string ``
	Verbose bool   ``
}

// New creates and returns a new empty Configuration
func New() *Configuration {
	return &Configuration{}
}
