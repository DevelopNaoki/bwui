package main

import (
	"fmt"
	"os"

	"github.com/DevelopNaoki/bwui/config"
	"github.com/DevelopNaoki/bwui/server"
)

func main() {
	serverConfig, err := config.GetServerConfig()
	if err != nil {
		fmt.Print(err)
		fmt.Printf("\n")
		os.Exit(-1)
	}

	server.StartServer(serverConfig)
}
