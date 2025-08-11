package infra

import (
	"bufio"
	"fmt"
)

type Job struct {
	command  *Command
	started  bool
	finished bool
	err      error
	logCh    chan string
}

func NewJob(cmd *Command) *Job {
	return &Job{
		command:  cmd,
		logCh:    make(chan string),
		err:      nil,
		finished: false,
	}
}

func (j *Job) RunWithProgress() {
	j.start()
	if err := runWithProgress(j.logCh, j.command); err != nil {
		j.finish(fmt.Errorf("unable to run job with progress: %w", err))
		return
	}
	j.finish(nil)
}

func runWithProgress(logCh chan string, command *Command) error {
	cmd := command.CMD
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
		logCh <- string(command.JsonProgressOutputFunc(line))
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("unable to read stdout scanner: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("unable to wait cmd: %w", err)
	}

	return nil
}

func (j *Job) start() {
	j.started = true
	j.finished = false
	j.err = nil
}

func (j *Job) finish(err error) {
	close(j.logCh)
	j.finished = true
	j.err = err
}
