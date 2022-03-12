package command

import (
	"encoding/json"
	"fmt"
	"io"

	yaml "gopkg.in/yaml.v3"

	"github.com/chroju/awscccli/awscc"
	"github.com/spf13/cobra"
)

type get struct {
	Manager awscc.AWSCCManager

	StdOut io.Writer
	ErrOut io.Writer

	typeName   string
	identifier string
	format     string
}

func newGetCommand(globalOption *GlobalOption) *cobra.Command {
	get := &get{}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			manager, err := awscc.New(globalOption.Profile, globalOption.Region)
			if err != nil {
				return err
			}
			get.Manager = manager

			get.StdOut = globalOption.StdOut
			get.ErrOut = globalOption.ErrOut

			args = cmd.Flags().Args()
			get.typeName = args[0]
			get.identifier = args[1]

			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if format == "yaml" || format == "json" {
				get.format = format
			} else {
				return fmt.Errorf("--format must be yaml or json")
			}

			return get.Execute()
		},
	}

	cmd.Flags().String("format", "json", "output format")
	cmd.SetOut(globalOption.StdOut)
	cmd.SetErr(globalOption.ErrOut)

	return cmd
}

func (get *get) Execute() error {
	resp, err := get.Manager.GetResource(get.typeName, get.identifier)
	if err != nil {
		return err
	}

	var f map[string]interface{}
	if err = json.Unmarshal([]byte(*resp), &f); err != nil {
		return err
	}

	var data []byte

	switch get.format {
	case "yaml":
		data, err = yaml.Marshal(f)
	case "json":
		data, err = json.MarshalIndent(f, "", "  ")
	}
	if err != nil {
		return err
	}

	fmt.Fprintf(get.StdOut, "%v\n", string(data))
	return nil
}
