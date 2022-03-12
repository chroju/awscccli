package command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/chroju/awscccli/awscc"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

type list struct {
	Manager awscc.AWSCCManager

	StdOut io.Writer
	ErrOut io.Writer

	typeName string
	format   string
	details  bool
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

			args = cmd.Flags().Args()
			list.details, err = cmd.Flags().GetBool("details")
			if err != nil {
				return err
			}

			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if format == "yaml" || format == "json" {
				list.format = format
			} else {
				return fmt.Errorf("--format must be yaml or json")
			}

			return list.Execute()
		},
	}

	cmd.Flags().String("format", "json", "output format (effective only with --details)")
	cmd.Flags().Bool("details", false, "show each resource details")
	cmd.SetOut(globalOption.StdOut)
	cmd.SetErr(globalOption.ErrOut)

	return cmd
}

func (list *list) Execute() error {
	resp, err := list.Manager.ListResources(list.typeName)
	if err != nil {
		return err
	}

	if list.details {
		resources := make([]map[string]interface{}, len(resp))
		var i int
		for _, value := range resp {
			var f map[string]interface{}
			if err = json.Unmarshal([]byte(*value), &f); err != nil {
				return err
			}
			resources[i] = f
			i++
		}

		var data []byte

		switch list.format {
		case "yaml":
			data, err = yaml.Marshal(resources)
		case "json":
			data, err = json.MarshalIndent(resources, "", "  ")
		}
		if err != nil {
			return err
		}

		fmt.Fprintf(list.StdOut, "%v\n", string(data))
	} else {
		for key := range resp {
			fmt.Fprintln(list.StdOut, *key)
		}
	}
	return nil
}
