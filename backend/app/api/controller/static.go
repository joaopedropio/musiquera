package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type StaticController struct {
	path string
}

func NewStaticController(path string) *StaticController {
	return &StaticController{
		path: path,
	}
}

func (c *StaticController) ServeStatic(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	if strings.Contains(requestPath, "media") {
		http.NotFound(w, r)
		return
	}
	filePath := filepath.Join(c.path, requestPath)

	// Prevent directory traversal
	if strings.Contains(requestPath, "..") {
		http.NotFound(w, r)
		return
	}

	// Serve file if it exists
	if fi, err := os.Stat(filePath); err == nil && !fi.IsDir() {
		http.ServeFile(w, r, filePath)
		return
	}

	// Serve index.html for React SPA routes
	http.ServeFile(w, r, filepath.Join(c.path, "index.html"))
}
