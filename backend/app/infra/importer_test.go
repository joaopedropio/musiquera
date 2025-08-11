package infra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/infra"
)

func TestImporter_X(t *testing.T) {
	t.Skip("This test should be only called locally")
	importer := infra.NewImporter()
	videoURL := "https://www.youtube.com/watch?v=J1p_OCoHwBg"
	outputDir := "."
	command := importer.CreateDownloadAudioWithProgressJob(videoURL, outputDir, infra.YTPDLPMP3Format.String())
	assert.NotNil(t, command)
}
