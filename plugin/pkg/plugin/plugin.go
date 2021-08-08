package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/scriptonist/example-go-grpc-plugin/plugin/pkg/proto"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var PluginMap = map[string]plugin.Plugin{
	"todo": &TodoPlugin{},
}

type TodoItem struct {
	Id          string
	Completed   bool
	Description string
}

type Todo interface {
	Create(context.Context, string) error
	Read(context.Context, string) (*TodoItem, error)
}

type TodoPlugin struct {
	plugin.Plugin
	Impl Todo
}

// GRPCServer should register this plugin for serving with the
// given GRPCServer. Unlike Plugin.Server, this is only called once
// since gRPC plugins serve singletons.
func (t *TodoPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterTodoServer(s, &TodoGRPCServer{Impl: t.Impl})
	return nil
}

// GRPCClient should return the interface implementation for the plugin
// you're serving via gRPC. The provided context will be canceled by
// go-plugin in the event of the plugin process exiting.
func (t *TodoPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &TodoGRPCClient{client: proto.NewTodoClient(c)}, nil
}

var _ plugin.GRPCPlugin = &TodoPlugin{}
