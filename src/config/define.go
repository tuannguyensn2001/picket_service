package config

type IConfig interface {
	GetEnv() string
	GetGrpcAddress() string
	GetHttpAddress() string
	GetGoogleClientId() string
	GetGoogleClientSecret() string
	GetClientUrl() string
}

func (c config) GetGoogleClientId() string {
	return c.GoogleClientId
}

func (c config) GetGoogleClientSecret() string {
	return c.GoogleClientSecret
}

func (c config) GetClientUrl() string {
	return c.ClientUrl
}
