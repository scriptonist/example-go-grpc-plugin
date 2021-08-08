package plugin

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/scriptonist/example-go-grpc-plugin/plugin/pkg/proto"
)

type TodoGRPCClient struct {
	client proto.TodoClient
}

func (c TodoGRPCClient) Create(ctx context.Context, description string) error {
	_, err := c.client.Create(ctx, &proto.CreateRequest{Description: description})
	if err != nil {
		return err
	}
	return nil
}

func (c TodoGRPCClient) Read(ctx context.Context, id string) (*TodoItem, error) {
	item, err := c.client.Read(ctx, &proto.ReadRequest{Id: id})
	if err != nil {
		return nil, err
	}
	resp := &TodoItem{
		Id:          item.Id,
		Completed:   item.Completed,
		Description: item.Description,
	}
	return resp, nil
}

type TodoGRPCServer struct {
	Impl Todo
	proto.UnimplementedTodoServer
}

func (s TodoGRPCServer) Create(ctx context.Context, req *proto.CreateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.Impl.Create(ctx, req.Description)
}
func (s TodoGRPCServer) Read(ctx context.Context, req *proto.ReadRequest) (*proto.TodoItem, error) {
	item, err := s.Impl.Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.TodoItem{Id: item.Id, Description: item.Description, Completed: item.Completed}, nil
}
