package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/scriptonist/example-go-grpc-plugin/plugin/pkg/plugin"
)

type Todo struct {
	dataDirectory string
}

func New() *Todo {
	os.MkdirAll("./data", 0755)
	d, _ := filepath.Abs("./data")
	return &Todo{
		dataDirectory: d,
	}
}

type Item struct {
	ID          string
	Description string
	Completed   bool
}

func (t *Todo) Create(_ context.Context, description string) error {
	id := uuid.New()
	item := Item{
		ID:          id.String(),
		Description: description,
		Completed:   false,
	}
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(t.dataDirectory, id.String()), b, 0655)
}

func (t *Todo) Read(_ context.Context, id string) (*plugin.TodoItem, error) {
	_, err := os.Stat(filepath.Join(t.dataDirectory, id))
	if err != nil {
		return nil, fmt.Errorf("item not found")
	}
	contents, err := ioutil.ReadFile(filepath.Join(t.dataDirectory, id))
	if err != nil {
		return nil, fmt.Errorf("reading contents failed")
	}
	var item Item
	err = json.Unmarshal(contents, &item)
	if err != nil {
		return nil, err
	}
	return &plugin.TodoItem{
		Id:          item.ID,
		Completed:   item.Completed,
		Description: item.Description,
	}, nil
}
