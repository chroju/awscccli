package command

import (
	"fmt"
	"io"

	"github.com/chroju/awscccli/awscc"
	"github.com/spf13/cobra"
)

type types struct {
	Manager awscc.AWSCCManager

	StdOut io.Writer
	ErrOut io.Writer
}

func newTypesCommand(globalOption *GlobalOption) *cobra.Command {
	types := &types{}

	cmd := &cobra.Command{
		Use:   "types",
		Short: "",
		Args:  nil,
		RunE: func(cmd *cobra.Command, args []string) error {
			manager, err := awscc.New(globalOption.Profile, globalOption.Region)
			if err != nil {
				return err
			}
			types.Manager = manager

			types.StdOut = globalOption.StdOut
			types.ErrOut = globalOption.ErrOut

			return types.Execute()
		},
	}

	cmd.SetOut(globalOption.StdOut)
	cmd.SetErr(globalOption.ErrOut)

	return cmd
}

func (types *types) Execute() error {
	resp, err := types.Manager.ListTypes()
	if err != nil {
		return err
	}
	for _, v := range resp {
		fmt.Fprintln(types.StdOut, *v)
	}
	return nil
}
