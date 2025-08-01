package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/joaopedropio/musiquera/app"
)

type SecurityController struct {
	a app.Application
}

func NewSecurityController(a app.Application) *SecurityController {
	return &SecurityController{
		a,
	}
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *SecurityController) Login(w http.ResponseWriter, r *http.Request) {
	loginInfo := &LoginInfo{}
	if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
		http.Error(w, fmt.Sprintf("unable to unmarshaw login info from body: %s", err.Error()), http.StatusInternalServerError)
	}

	if loginInfo.Username == "" {
		http.Error(w, "username can't be empty", http.StatusBadRequest)
		return
	}
	if loginInfo.Password == "" {
		http.Error(w, "password can't be empty", http.StatusBadRequest)
		return
	}
	fmt.Println("login info: " + loginInfo.Username + " " + loginInfo.Password)
	token, err := c.a.LoginService().Login(loginInfo.Username, loginInfo.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to login: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Set the token as an HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	_, err = fmt.Fprint(w, "Login Successful")
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to write response: %s", err.Error()), http.StatusInternalServerError)
	}
}

func (c *SecurityController) Logout(w http.ResponseWriter, r *http.Request) {
	// Overwrite the JWT cookie with an expired one
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt", // or whatever your JWT cookie is named
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0), // Expire immediately
		MaxAge:   -1,
	})

	// Optional: Redirect to login or home
	http.Redirect(w, r, "/loginPage", http.StatusSeeOther)
}

func (c *SecurityController) AuthCheck(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		fmt.Println("no token cookie found")
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	isLogged, err := c.a.LoginService().IsLogged(cookie.Value)
	if err != nil {
		fmt.Println(fmt.Errorf("unable to check if user is logged: %w", err).Error())
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	if !isLogged {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
