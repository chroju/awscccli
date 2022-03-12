package command

import (
	"github.com/spf13/cobra"
)

type types struct{}

func newTypesCommand(globalOption *GlobalOption) *cobra.Command {
	types := &types{}

	cmd := &cobra.Command{
		Use:   "types",
		Short: "",
		Args:  nil,
		RunE: func(cmd *cobra.Command, args []string) error {
			return types.Execute()
		},
	}

	cmd.SetOut(globalOption.StdOut)
	cmd.SetErr(globalOption.ErrOut)

	return cmd
}

func (o *types) Execute() error {
	return nil
}
