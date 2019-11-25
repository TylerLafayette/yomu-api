package main

import (
	"os"
)

// DefaultPort is the app's default port it will listen on if the
// environment variable is not set.
const DefaultPort = "8080"

// Config contains configuration information for the app.
type Config struct {
	port string // the port the app will run on
}

// NewConfig creates and returns a new Config struct.
func NewConfig(port string) *Config {
	return &Config{port}
}

// CreateConfig constructs a new Config struct using default values.
func CreateConfig() *Config {
	// Get the PORT environment variable.
	port := os.Getenv("PORT")

	if port == "" {
		// Use the default port if no environment variable is found.
		port = DefaultPort
	}

	port = ":" + port

	return NewConfig(port)
}
