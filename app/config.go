package app

import flags "github.com/jessevdk/go-flags"

// Config holds the app configuration
type Config struct {
	AppName     string
	AppVersion  string
	Environment string `short:"e" long:"environment" description:"Sets the app environment" default:"test" env:"GOWORK_ENV"`
}

// NewConfig returns a new instance of app's configuration
func NewConfig() (*Config, error) {
	var config Config
	var parser = flags.NewParser(&config, flags.IgnoreUnknown)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	config.AppName = appName
	config.AppVersion = appVersion

	return &config, nil
}
