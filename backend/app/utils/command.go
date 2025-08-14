package utils

import (
	"bufio"
	"fmt"
	"os/exec"
)

type HandleCMDOutput func(line string) string

func RunCommandWithProgress(cmd *exec.Cmd, logCh chan string, handleOutput HandleCMDOutput) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("unable to get stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("unable to start command: %w", err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		logCh <- handleOutput(line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("unable to read stdout scanner: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("unable to wait cmd: %w", err)
	}

	return nil
}
