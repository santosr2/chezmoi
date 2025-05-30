package cmd

import (
	"fmt"
	"os/exec"
)

type cmdOutputError struct {
	path   string
	args   []string
	output []byte
	err    error
}

func newCmdOutputError(cmd *exec.Cmd, output []byte, err error) *cmdOutputError {
	return &cmdOutputError{
		path:   cmd.Path,
		args:   cmd.Args,
		output: output,
		err:    err,
	}
}

func (e *cmdOutputError) Error() string {
	if len(e.output) == 0 {
		return fmt.Sprintf("%s: %v", shellQuoteCommand(e.path, e.args[1:]), e.err)
	}
	return fmt.Sprintf("%s: %v\n%s", shellQuoteCommand(e.path, e.args[1:]), e.err, e.output)
}

func (e *cmdOutputError) Unwrap() error {
	return e.err
}
