package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/mail"
)

const emailServiceAPIUrl = "https://api.postmarkapp.com/email/withTemplate"

type EmailService interface {
	SendCreateAccountMFACode(mfaCode, toAddress string) error
	IsValidEmailAddress(email string) bool
}

func NewEmailService(client *http.Client, emailAPIToken string, fromAddress string) EmailService {
	return &emailService{
		client:        client,
		emailAPIToken: emailAPIToken,
		fromAddress:   fromAddress,
	}
}

type emailService struct {
	client        *http.Client
	emailAPIToken string
	fromAddress   string
}

type templateModel struct {
	MfaCode string `json:"mfa_code"`
}

type mfaCodeBodyRequest struct {
	From          string `json:"From"`
	To            string `json:"To"`
	TemplateAlias string `json:"TemplateAlias"`
	TemplateModel *templateModel
}

func (s *emailService) IsValidEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *emailService) SendCreateAccountMFACode(mfaCode, toAddress string) error {
	data, err := s.createMfaCodeBodyRequest(mfaCode, s.fromAddress, toAddress)
	if err != nil {
		return fmt.Errorf("unable to create mfa code body request: %w", err)
	}
	req, err := http.NewRequest("POST", emailServiceAPIUrl, data)
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}
	defer func() {
		if cerr := req.Body.Close(); err != nil {
			err = fmt.Errorf("unable to close body: %w", cerr)
		}
	}()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Postmark-Server-Token", s.emailAPIToken)

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to send request to email provider: %w", err)
	}

	// Print response status
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Status: ", resp.Status)
	fmt.Println("Body: ", string(body))

	return nil
}

func (s *emailService) createMfaCodeBodyRequest(mfaCode string, fromAddress, toAddress string) (io.Reader, error) {
	data := &mfaCodeBodyRequest{
		From:          fromAddress,
		To:            toAddress,
		TemplateAlias: "mfa_code",
		TemplateModel: &templateModel{
			MfaCode: mfaCode,
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json body: %w", err)
	}
	return bytes.NewBuffer(jsonData), nil
}
