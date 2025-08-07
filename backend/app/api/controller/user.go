package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaopedropio/musiquera/app"
)

type CreateInviteResponse struct {
	InviteLink string `json:"invite_link"`
}

type UserController struct {
	a app.Application
}

func NewUserController(a app.Application) *UserController {
	return &UserController{
		a: a,
	}
}

func (c *UserController) CreateInvite(w http.ResponseWriter, r *http.Request) {
	_, link, err := c.a.InviteService().CreateInvite()
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create invite: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(CreateInviteResponse{
		InviteLink: link,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to encode json response: %v", err), http.StatusInternalServerError)
		return
	}
}
