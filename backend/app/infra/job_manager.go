package infra

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type JobManager interface {
	AddJob(job *Job)
	RemoveJob(jobID uuid.UUID)
	Run()

	AddWebSocketClient(conn *websocket.Conn)
	RemoveWebSocketClient(conn *websocket.Conn)
}

type jobManager struct {
	jobs      map[uuid.UUID]*Job
	jobsMutex sync.Mutex

	webSocketClients      map[*websocket.Conn]bool
	webSocketClientsMutex sync.Mutex
}

func NewJobManager() *jobManager {
	return &jobManager{
		jobs:             make(map[uuid.UUID]*Job),
		webSocketClients: make(map[*websocket.Conn]bool),
	}
}

func (m *jobManager) AddWebSocketClient(conn *websocket.Conn) {
	m.webSocketClientsMutex.Lock()
	m.webSocketClients[conn] = true
	m.webSocketClientsMutex.Unlock()
}

func (m *jobManager) RemoveWebSocketClient(conn *websocket.Conn) {
	m.webSocketClientsMutex.Lock()
	delete(m.webSocketClients, conn)
	m.webSocketClientsMutex.Unlock()
}

func (m *jobManager) AddJob(job *Job) {
	m.jobsMutex.Lock()
	m.jobs[uuid.New()] = job
	m.jobsMutex.Unlock()
}

func (m *jobManager) RemoveJob(jobID uuid.UUID) {
	m.jobsMutex.Lock()
	delete(m.jobs, jobID)
	m.jobsMutex.Unlock()
}

func (m *jobManager) GetJobs() map[uuid.UUID]*Job {
	return m.jobs
}

func (m *jobManager) Run() {
	for {
		for jobID, job := range m.jobs {
			if !job.started {
				go m.subscribeToWebSocket(jobID, job.logCh)
				go job.RunWithProgress()
			}
			if job.finished {
				m.RemoveJob(jobID)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (m *jobManager) subscribeToWebSocket(jobID uuid.UUID, logCh chan string) {
	for line := range logCh {
		m.broadcastLog(jobID, line)
	}
}

// Broadcast the log messages to all connected clients
func (m *jobManager) broadcastLog(jobID uuid.UUID, logMessage string) {
	m.webSocketClientsMutex.Lock()
	defer m.webSocketClientsMutex.Unlock()

	m.jobsMutex.Lock()
	defer m.jobsMutex.Unlock()

	for client := range m.webSocketClients {
		job := m.jobs[jobID]

		errorMessage := ""
		if job.err != nil {
			errorMessage = job.err.Error()
		}
		message := &JobMessage{
			JobID:        jobID.String(),
			Started:      job.started,
			Finished:     job.finished,
			ErrorMessage: errorMessage,
			Progress:     json.RawMessage(logMessage),
		}
		json, err := json.Marshal(message)
		if err != nil {
			fmt.Println(fmt.Errorf("unable to marshal job message: %w", err).Error())
			continue
		}
		err = client.WriteMessage(websocket.TextMessage, json)
		if err != nil {
			fmt.Println("Error sending message:", err)
			if err := client.Close(); err != nil {
				fmt.Println(fmt.Errorf("unable to close websocket: %w", err))
			}
			delete(m.webSocketClients, client)
		}
	}
}

type JobMessage struct {
	JobID        string          `json:"job_id"`
	Started      bool            `json:"started"`
	Finished     bool            `json:"finished"`
	ErrorMessage string          `json:"error,omitempty"`
	Progress     json.RawMessage `json:"progress"`
}
