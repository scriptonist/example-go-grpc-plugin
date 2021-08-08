package todo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/scriptonist/example-go-grpc-plugin/plugin/pkg/plugin"
)

// t *Todo
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

func (t *Todo) Create(_ context.Context, description string) error {
	id := uuid.New()
	log.Println(id.String())
	return ioutil.WriteFile(filepath.Join(t.dataDirectory, id.String()), []byte(description), 0655)
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
	return &plugin.TodoItem{
		Id:          id,
		Completed:   false,
		Description: string(contents),
	}, nil
}
