package command

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	commandName      = "acc"
	shortDescription = ""
	longDescription  = ""
)

type GlobalOption struct {
	Profile   string
	Region    string
	IsNoColor bool

	StdOut io.Writer
	ErrOut io.Writer
}

func NewCommand(version string, stdOutWriter, errOutWriter io.Writer) (*cobra.Command, error) {
	o := &GlobalOption{}
	cmd := &cobra.Command{
		Use:          commandName,
		Version:      version,
		Short:        shortDescription,
		Long:         longDescription,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("Use subcommand: types, list, get")
		},
	}
	cmd.PersistentFlags().StringVarP(&o.Profile, "profile", "p", "", "AWS profile")
	// aws-sdk-go does not support the AWS_DEFAULT_REGION environment variable
	cmd.PersistentFlags().StringVar(&o.Region, "region", os.Getenv("AWS_DEFAULT_REGION"), "AWS region")
	// cmd.PersistentFlags().BoolVar(&o.IsNoColor, "no-color", false, "Turn off colored output")

	o.StdOut = stdOutWriter
	o.ErrOut = errOutWriter
	cmd.SetOut(stdOutWriter)
	cmd.SetErr(errOutWriter)

	cmd.AddCommand(
		newTypesCommand(o),
		newListCommand(o),
	)

	return cmd, nil
}
