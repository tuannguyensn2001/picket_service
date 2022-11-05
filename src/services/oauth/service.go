package oauth_service

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"myclass_service/src/config"
	"strings"
)

type service struct {
	config config.IConfig
}

func New(config config.IConfig) *service {
	return &service{config: config}
}

func (s *service) GetAccessTokenFromCode(ctx context.Context, code string) (string, error) {
	client := resty.New()

	body := map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"client_id":     s.config.GetGoogleClientId(),
		"client_secret": s.config.GetGoogleClientSecret(),
		"redirect_uri":  s.config.GetClientUrl(),
	}

	type ResponseError struct {
		Error string `json:"error"`
	}

	type ResponseSuccess struct {
		AccessToken string `json:"access_token"`
	}

	var respErr *ResponseError

	resp, err := client.R().
		SetFormData(body).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetError(respErr).
		Post("https://oauth2.googleapis.com/token")
	if err != nil {
		return "", err
	}

	if respErr != nil {
		return "", errors.New(respErr.Error)
	}

	var respSuccess ResponseSuccess
	f := bufio.NewReader(strings.NewReader(string(resp.Body())))
	err = json.NewDecoder(f).Decode(&respSuccess)
	if err != nil {
		return "", err
	}

	return respSuccess.AccessToken, nil
}
