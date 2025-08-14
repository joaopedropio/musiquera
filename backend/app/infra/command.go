package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/joaopedropio/musiquera/app/utils"
)

const YTPDLPMP3Format = YTDLPFormat("mp3")

type YTDLPFormat string

func (f YTDLPFormat) String() string {
	return string(f)
}

type Progress struct {
	Percentage string `json:"percentage"`
	Size       string `json:"size"`
	Speed      string `json:"speed"`
	ETA        string `json:"eta"`
}

func (p *Progress) String() string {
	return fmt.Sprintf("Progress: %s%% | Size: %s | Speed: %s | ETA: %s",
		p.Percentage, p.Size, p.Speed, p.ETA)
}

func (p *Progress) JSON() ([]byte, error) {
	return json.Marshal(p)
}

type Command interface {
	Execute(logCh chan string) error
}

type formatTrackCommand struct {
}

func NewFormatTrackCommand() Command {
	return &formatTrackCommand{}
}

func (c *formatTrackCommand) Execute(logCh chan string) error {
	return nil
}

type downloadTrackCommand struct {
	videoURL  string
	outputDir string
	format    string
}

func NewDownloadTrackCommand(videoURL, outputDir, format string) Command {
	return &downloadTrackCommand{
		videoURL:  videoURL,
		outputDir: outputDir,
		format:    format,
	}
}

func (c *downloadTrackCommand) Execute(logCh chan string) error {
	cmd := exec.Command("yt-dlp",
		"-x",
		"--newline",
		"--progress-template", "[download] %(progress._percent_str)s;%(progress._total_bytes_str)s;%(progress._speed_str)s;%(progress._eta_str)s",
		"--audio-format", c.format,
		"-o", fmt.Sprintf("%s/%%(title)s.%%(ext)s", c.outputDir),
		c.videoURL,
	)

	return utils.RunCommandWithProgress(cmd, logCh, c.handleOutput)
}

func (c *downloadTrackCommand) handleOutput(cmdLineOutput string) string {
	progress, err := c.parseDownloadProgessLine(cmdLineOutput)
	if err != nil {
		return ""
	}
	json, err := progress.JSON()
	if err != nil {
		return ""
	}
	return string(json)
}

func (c *downloadTrackCommand) parseDownloadProgessLine(line string) (*Progress, error) {
	if strings.Contains(line, "[download]") {
		line = strings.ReplaceAll(line, "[download] ", "")
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			return nil, fmt.Errorf("line contains [download] but does not have 4 fields, have %d fields", len(fields))
		}
		return &Progress{
			Percentage: strings.Trim(fields[0], " "),
		}, nil
	}
	return nil, errors.New("unable to parse download progress line")
}
