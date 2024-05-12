package main

import (
	"github.com/amirrstm/go-auth/cmd/server"
	"github.com/amirrstm/go-auth/pkg/config"
)

func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")

	server.Serve()

}
