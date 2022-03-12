package command

import (
	"fmt"
	"io"

	"github.com/chroju/awscccli/awscc"
	"github.com/spf13/cobra"
)

type list struct {
	Manager awscc.AWSCCManager

	StdOut io.Writer
	ErrOut io.Writer

	typeName string
}

func newListCommand(globalOption *GlobalOption) *cobra.Command {
	list := &list{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			args = cmd.Flags().Args()
			if len(args) != 0 {
				list.typeName = args[0]
			}
			manager, err := awscc.New(globalOption.Profile, globalOption.Region)
			if err != nil {
				return err
			}
			list.Manager = manager

			list.StdOut = globalOption.StdOut
			list.ErrOut = globalOption.ErrOut

			return list.Execute()
		},
	}

	cmd.SetOut(globalOption.StdOut)
	cmd.SetErr(globalOption.ErrOut)

	return cmd
}

func (list *list) Execute() error {
	resp, err := list.Manager.ListResources(list.typeName)
	if err != nil {
		return err
	}
	for _, v := range resp {
		fmt.Fprintln(list.StdOut, *v)
	}
	return nil
}
