package infra

import (
	"fmt"
)

type LoginService interface {
	Login (username, password string) (string, error)
}

func NewLoginService(passwordService PasswordService, userRepo UserRepo) LoginService {
	return &loginService{
		passwordService: passwordService,
		userRepo: userRepo,
	}
}

type loginService struct {
	passwordService PasswordService
	userRepo UserRepo
}

func (s *loginService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("unable to get user by username: %w", err)
	}
	fmt.Println("saved pass: " + user.Password())
	fmt.Println("password: " + password)
	verified, err := s.passwordService.VerifyPassword(password, user.Password())
	if err != nil {
		return "", fmt.Errorf("unable to verify password: %w", err)
	}
	if !verified {
		return "", fmt.Errorf("password does not match")
	}

	_, jwt, _ := s.passwordService.JWTAuth().Encode(map[string]interface{}{"username":user.Username()})

	return jwt, nil
}
