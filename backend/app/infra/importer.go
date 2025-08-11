package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
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

type importer struct {
}

func NewImporter() *importer {
	return &importer{}
}

func (i *importer) CreateDownloadAudioWithProgressJob(videoURL, outputDir, format string) *Command {
	cmd := exec.Command("yt-dlp",
		"-x",
		"--newline",
		"--progress-template", "[download] %(progress._percent_str)s;%(progress._total_bytes_str)s;%(progress._speed_str)s;%(progress._eta_str)s",
		"--audio-format", format,
		"-o", fmt.Sprintf("%s/%%(title)s.%%(ext)s", outputDir),
		videoURL,
	)

	return &Command{
		CMD: cmd,
		JsonProgressOutputFunc: func(cmdLineOutput string) []byte {
			progress, err := parseDownloadProgessLine(cmdLineOutput)
			if err != nil {
				return nil
			}
			json, err := progress.JSON()
			if err != nil {
				return nil
			}
			return json
		},
	}
}

func parseDownloadProgessLine(line string) (*Progress, error) {
	if strings.Contains(line, "[download]") {
		line = strings.ReplaceAll(line, "[download] ", "")
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			return nil, fmt.Errorf("line contains [download] but does not have 4 fields, have %d fields", len(fields))
		}
		return &Progress{
			Percentage: strings.Trim(fields[0], " "),
			Size:       strings.Trim(fields[1], " "),
			Speed:      strings.Trim(fields[2], " "),
			ETA:        strings.Trim(fields[3], " "),
		}, nil
	}
	fmt.Printf("not progress line: %s\n", line)
	return nil, errors.New("not a download progress line")
}
