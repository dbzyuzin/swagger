package config

import "os"

var DebugMode bool = false
var ServerHost string = ":8080"

func init() {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		DebugMode = true
	}

	serverHost := os.Getenv("SERVER_HOST")
	if serverHost != "" {
		ServerHost = serverHost
	}
}