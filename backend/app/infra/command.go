package infra

import "os/exec"

type Command struct {
	CMD                    *exec.Cmd
	JsonProgressOutputFunc func(cmdLineOutput string) []byte
}
