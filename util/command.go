package util

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
)

func NewBashCommand(ctx context.Context, command string) *exec.Cmd {
	return NewPgidCommand(ctx, "bash", "-c", command)
}

// Why won't Go kill a child process correctly?
// https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly/78584235#78584235
func NewPgidCommand(ctx context.Context, name string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)

	// Killing a child process and all of its children in Go
	// https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Mutate and override cmd.Cancel
	cmd.Cancel = func() error {
		var errors []error
		if err := cmd.Process.Kill(); err != nil {
			errors = append(errors, err)
		}
		if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
			errors = append(errors, err)
		}
		if len(errors) > 0 {
			return fmt.Errorf("error cancelling pid=%d, errors=%v", cmd.Process.Pid, errors)
		}
		return nil
	}
	return cmd
}
