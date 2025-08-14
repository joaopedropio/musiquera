package infra

import (
	"fmt"
)

type Job struct {
	started     bool
	finished    bool
	currentTask string
	err         error
	logCh       chan string

	videoURL string
}

func NewAddTrackJob(videoURL string) *Job {
	return &Job{
		logCh:    make(chan string),
		videoURL: videoURL,
	}
}

func (j *Job) RunWithProgress() {
	j.start()
	j.currentTask = "Downloading track from YouTube"
	download := NewDownloadTrackCommand(j.videoURL, "./", YTPDLPMP3Format.String())
	if err := download.Execute(j.logCh); err != nil {
		j.finish(fmt.Errorf("unable to download track: %w", err))
		return
	}

	j.currentTask = "Formatting track to MPEG Dash"
	format := NewFormatTrackCommand()
	if err := format.Execute(j.logCh); err != nil {
		j.finish(fmt.Errorf("unable to format audio: %w", err))
	}
	j.finish(nil)
}

func (j *Job) start() {
	j.currentTask = "Job started"
	j.started = true
	j.finished = false
	j.err = nil
}

func (j *Job) finish(err error) {
	if err != nil {
		j.currentTask = "Finished Successfully"
	} else {
		j.currentTask = "Finished Unsuccessfully"
	}
	j.finished = true
	j.err = err
}
