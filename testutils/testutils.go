package testutils

import (
	"bytes"
	"io"

	"github.com/spf13/cobra"
)

func Exec(cmd *cobra.Command, args ...string) (out []byte, err error) {
	b := bytes.NewBuffer(nil)
	cmd.SetOut(b)
	cmd.SetArgs(args)
	err = cmd.Execute()
	if err != nil {
		return nil, err
	}

	return io.ReadAll(b)
}
