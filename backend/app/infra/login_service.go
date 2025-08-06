package infra

import (
	"fmt"

	"github.com/go-chi/jwtauth/v5"

	"github.com/joaopedropio/musiquera/app/utils"
)

type LoginService interface {
	JWTAuth() *jwtauth.JWTAuth
	Login(username, password string) (string, error)
	IsLogged(token string) (bool, error)
}

func NewLoginService(jwtSecret string, userRepo UserRepo) LoginService {
	return &loginService{
		jwtAuth:  createJwtAuth(jwtSecret),
		userRepo: userRepo,
	}
}

type loginService struct {
	jwtAuth  *jwtauth.JWTAuth
	userRepo UserRepo
}

func (s *loginService) JWTAuth() *jwtauth.JWTAuth {
	return s.jwtAuth
}

func createJwtAuth(jwtSecret string) *jwtauth.JWTAuth {
	algorithm := "HS256"
	signKey := []byte(jwtSecret)
	return jwtauth.New(algorithm, signKey, nil)
}

func (s *loginService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("unable to get user by username: %w", err)
	}
	verified, err := utils.VerifyPassword(password, user.Password())
	if err != nil {
		return "", fmt.Errorf("unable to verify password: %w", err)
	}
	if !verified {
		return "", fmt.Errorf("password does not match")
	}

	_, jwt, _ := s.jwtAuth.Encode(map[string]interface{}{"username": user.Username()})

	return jwt, nil
}

func (s *loginService) IsLogged(t string) (bool, error) {
	token, err := s.jwtAuth.Decode(t)
	if err != nil {
		return false, fmt.Errorf("unable to decode jwt token: %w", err)
	}

	value, ok := token.Get("username")
	if !ok {
		return false, fmt.Errorf("jwt token (besides valid) does not have username filed")
	}
	username := value.(string)
	if username == "" {
		return false, nil
	}

	return true, nil
}
