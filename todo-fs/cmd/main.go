package main

import (
	"github.com/hashicorp/go-plugin"
	todoplugin "github.com/scriptonist/example-go-grpc-plugin/plugin/pkg/plugin"
	"github.com/scriptonist/example-go-grpc-plugin/service/internal/todo"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: todoplugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			"todo": &todoplugin.TodoPlugin{Impl: todo.New()},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}
