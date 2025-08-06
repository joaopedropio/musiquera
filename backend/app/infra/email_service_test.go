package infra_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/infra"
)

func TestEmailService_EmailAddressValidation(t *testing.T) {
	service := infra.NewEmailService(nil, "", "")
	assert.True(t, service.IsValidEmailAddress("musiquera@musiquera.uk"))
	assert.True(t, service.IsValidEmailAddress("example@gmail.com"))
	assert.False(t, service.IsValidEmailAddress("email"))
	assert.False(t, service.IsValidEmailAddress("email@@gmail.com"))
	assert.False(t, service.IsValidEmailAddress("email@exx gmail"))
	assert.False(t, service.IsValidEmailAddress("email@.gmail"))
	assert.False(t, service.IsValidEmailAddress("@gmail"))
	assert.False(t, service.IsValidEmailAddress("email @gmai"))
}

func TestEmailService_ShouldContainValidBody(t *testing.T) {
	// Arrange
	providerURL := "https://api.postmarkapp.com/email/withTemplate"
	mfaCode := "123456"
	toEmailAddress := "to@email.com"
	fromEmailAddress := "from@email.com"
	apiToken := "api_token"
	clientMock := NewHTTPClientMock(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, providerURL, req.URL.String())
		assert.Equal(t, req.Method, "POST")

		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, apiToken, req.Header.Get("X-Postmark-Server-Token"))

		body, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"From":"from@email.com","To":"to@email.com","TemplateAlias":"mfa_code","TemplateModel":{"mfa_code":"123456"}}`, string(body))

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"success":true}`)),
			Header:     make(http.Header),
		}, nil
	})
	emailService := infra.NewEmailService(clientMock, apiToken, fromEmailAddress)

	// Act
	err := emailService.SendCreateAccountMFACode(mfaCode, toEmailAddress)

	// Assert
	assert.NoError(t, err)
}

func NewHTTPClientMock(handler func(req *http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: &MockRoundTripper{
			RoundTripFunc: handler,
		},
	}
}

type MockRoundTripper struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}
