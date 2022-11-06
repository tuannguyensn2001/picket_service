package oauth_service

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"myclass_service/src/config"
	"myclass_service/src/entities"
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

	resp, err := client.R().
		SetContext(ctx).
		SetFormData(body).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetError(&ResponseError{}).
		Post("https://oauth2.googleapis.com/token")

	if err != nil {
		zap.S().Error(err)
		return "", err
	}

	if val, ok := resp.Error().(*ResponseError); ok {
		zap.S().Error(val)
		return "", errors.New(val.Error)
	}

	var respSuccess ResponseSuccess
	f := bufio.NewReader(strings.NewReader(string(resp.Body())))
	err = json.NewDecoder(f).Decode(&respSuccess)
	if err != nil {
		return "", err
	}

	return respSuccess.AccessToken, nil
}

func (s *service) GetUserProfileByAccessToken(ctx context.Context, accessToken string) (*entities.User, error) {
	client := resty.New()

	type ResponseError struct {
		Error string `json:"error"`
	}
	var respErr *ResponseError

	resp, err := client.R().SetQueryParam("access_token", accessToken).SetError(respErr).Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	if respErr != nil {
		return nil, errors.New(respErr.Error)
	}

	type ResponseSuccess struct {
		Name    string `json:"name"`
		Picture string `json:"picture"`
		Email   string `json:"email"`
	}

	var response ResponseSuccess
	f := bufio.NewReader(strings.NewReader(string(resp.Body())))
	err = json.NewDecoder(f).Decode(&response)
	if err != nil {
		return nil, err
	}

	result := entities.User{
		Username: response.Name,
		Email:    response.Email,
		Profile: &entities.Profile{
			Avatar: response.Picture,
		},
	}
	return &result, nil

}
