package main

import (
	"log"

	"github.com/scriptonist/example-go-grpc-plugin/cli/internal/cli"
	"github.com/scriptonist/example-go-grpc-plugin/cli/internal/commands"
)

func main() {
	cli, err := cli.NewCLI("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// defer cli.KillPlugin()

	cmd := commands.BuildRootCmd(cli)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
