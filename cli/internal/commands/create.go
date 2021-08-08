package commands

import (
	"github.com/scriptonist/example-go-grpc-plugin/cli/internal/cli"
	"github.com/spf13/cobra"
)

func buildCreateCmd(c *cli.CLI) *cobra.Command {
	opts := cli.AddTodoOpts{}
	cmd := &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.AddTodo(opts)
		},
	}
	cmd.Flags().StringVar(&opts.Content, "content", "", "content for todo")

	return cmd
}
