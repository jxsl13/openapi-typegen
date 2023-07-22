package testutils

import (
	"bytes"
	"io"

	"github.com/spf13/cobra"
)

// Exec executes the cobra command and collects the output into a buffer for further analysis.
func Exec(cmd *cobra.Command, args ...string) (out []byte, err error) {
	b := bytes.NewBuffer(nil)
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs(args)
	err = cmd.Execute()
	if err != nil {
		return nil, err
	}

	return io.ReadAll(b)
}

// ExecWithStdout writes its output to stdout/stderr as well as whatever is previously defined
// to be the target (like a file)
func ExecWithStdout(cmd *cobra.Command, args ...string) (out []byte, err error) {
	b := bytes.NewBuffer(nil)
	cmd.SetOut(io.MultiWriter(cmd.OutOrStdout(), b))
	cmd.SetErr(io.MultiWriter(cmd.OutOrStderr(), b))
	cmd.SetArgs(args)
	err = cmd.Execute()
	if err != nil {
		return nil, err
	}

	return io.ReadAll(b)
}
