package app

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Error struct {
	Code       int    `json:"code" yaml:"code"`
	Message    string `json:"message" yaml:"message"`
	StatusCode int    `json:"status_code" yaml:"statusCode"`
	GrpcCode   int    `json:"grpc_code" yaml:"grpcCode"`
}

//General.Forbidden

func (e Error) Error() string {
	return e.Message
}

func LoadErr(url string) (map[string]map[string]Error, error) {
	yamlFile, err := ioutil.ReadFile(url)
	if err != nil {
		return nil, err
	}

	m := make(map[string]map[string]Error)

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
