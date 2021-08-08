module github.com/scriptonist/example-go-grpc-plugin/cli

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-plugin v1.4.2
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/scriptonist/example-go-grpc-plugin/plugin v0.0.1
	github.com/scriptonist/grpc-todo-example/service v0.0.0-20210801052723-50410e30c6e1
	github.com/spf13/cobra v1.2.1
)

replace github.com/scriptonist/example-go-grpc-plugin/plugin => ../plugin
