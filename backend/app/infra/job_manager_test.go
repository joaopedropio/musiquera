package infra_test

import (
	"testing"

	"github.com/joaopedropio/musiquera/app/infra"
)

func TestJobManager(t *testing.T) {
	t.Skip("This test should be called only locally")
	jobManager := infra.NewJobManager()
	importer := infra.NewImporter()
	videoURLs := []string{
		"https://www.youtube.com/watch?v=V5ftrI8hrcw",
		"https://www.youtube.com/watch?v=nO3yGPfZg24",
		"https://www.youtube.com/watch?v=1-zjtRghCQE",
		"https://www.youtube.com/watch?v=FEGXbB9b1lE",
		"https://www.youtube.com/watch?v=AS0yjHPXYI8",
		"https://www.youtube.com/watch?v=qpHWsWySwmc",
		"https://www.youtube.com/watch?v=YYkBSyoZxXQ",
		"https://www.youtube.com/watch?v=E9nrKitD05g",
	}
	for _, url := range videoURLs {
		cmd := importer.CreateDownloadAudioWithProgressJob(url, "./", infra.YTPDLPMP3Format.String())
		jobManager.AddJob(infra.NewJob(cmd))
	}
	jobManager.Run()
}
