package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"

	todoplugin "github.com/scriptonist/example-go-grpc-plugin/plugin/pkg/plugin"
)

type CLI struct {
	pluginClient *plugin.Client
	todo         todoplugin.Todo
}

func (c *CLI) KillPlugin() {
	c.pluginClient.Kill()
}

func NewCLI(apiServerAddr string) (*CLI, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: todoplugin.Handshake,
		Plugins:         todoplugin.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("TODO_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
			plugin.ProtocolNetRPC,
		},
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense("todo")
	if err != nil {
		return nil, err
	}
	c := &CLI{}
	c.pluginClient = client
	c.todo = raw.(todoplugin.Todo)
	// if todo, ok := raw.(todoplugin.Todo); !ok {
	// 	return nil, fmt.Errorf("cannot get instance of plugin")
	// } else {
	// 	c.todo = todo
	// }

	return c, nil
}

type AddTodoOpts struct {
	Content string
}

func (c *CLI) AddTodo(opts AddTodoOpts) error {
	return c.todo.Create(context.Background(), opts.Content)
}

type GetTodoOpts struct {
	Id string
}

func (c *CLI) GetTodo(opts GetTodoOpts) error {
	todo, err := c.todo.Read(context.Background(), opts.Id)
	if err != nil {
		return err
	}
	fmt.Println(todo)
	return nil
}
