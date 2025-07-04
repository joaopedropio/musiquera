package controller

import (
	"fmt"
	"net/http"
)

func NewPingPong() *PingPongController {
	return &PingPongController{}
}

type PingPongController struct{}

func (c *PingPongController) Get(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "pong")
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to write to response: %s", err.Error()), 500)
	}
}
