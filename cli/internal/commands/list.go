package commands

import (
	"github.com/scriptonist/example-go-grpc-plugin/cli/internal/cli"
	"github.com/spf13/cobra"
)

type ListCmdOpts struct{}

func buildListCmd(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetTodo(cli.GetTodoOpts{
				Id: args[0],
			})
		},
	}

	return cmd
}
