package controller

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"

	"github.com/joaopedropio/musiquera/app"
	"github.com/joaopedropio/musiquera/app/infra"
)

type JobController struct {
	a app.Application
}

func NewJobController(a app.Application) *JobController {
	return &JobController{
		a,
	}
}

func (c *JobController) RunJobs(w http.ResponseWriter, r *http.Request) {
	jobManager := c.a.JobManager()

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
		jobManager.AddJob(infra.NewAddTrackJob(url))
	}
}

func (c *JobController) JobProgress(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // Allow all connections (or make more restrictive)
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer func() {
		if cerr := conn.Close(); err != nil {
			fmt.Println(fmt.Errorf("unable to close websocket: %w", cerr))
		}
	}()

	c.a.JobManager().AddWebSocketClient(conn)

	// Keep the connection alive
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			// Client disconnected
			c.a.JobManager().RemoveWebSocketClient(conn)
			break
		}
	}
}
