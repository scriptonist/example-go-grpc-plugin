module github.com/scriptonist/example-go-grpc-plugin/todo-fs

go 1.16

require (
	github.com/google/uuid v1.1.2
	github.com/hashicorp/go-plugin v1.4.2
	github.com/scriptonist/example-go-grpc-plugin/plugin v1.0.0
)

replace github.com/scriptonist/example-go-grpc-plugin/plugin => ../plugin
