package utils

import "os/exec"

func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	return cmd.Output()
}
