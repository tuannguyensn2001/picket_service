package config

type config struct {
	Env                string
	GrpcAddress        string
	HttpAddress        string
	GoogleClientId     string
	GoogleClientSecret string
	ClientUrl          string
}

func GetConfig() config {
	structure := bootstrap()
	result := config{
		Env:                structure.App.Env,
		GrpcAddress:        structure.App.GrpcAddress,
		HttpAddress:        structure.App.HttpAddress,
		GoogleClientId:     structure.OAuth2.Google.ClientId,
		GoogleClientSecret: structure.OAuth2.Google.ClientSecret,
		ClientUrl:          structure.Client.Url,
	}

	return result
}

func (c config) GetEnv() string {
	return c.Env
}

func (c config) GetGrpcAddress() string {
	return c.GrpcAddress
}

func (c config) GetHttpAddress() string {
	return c.HttpAddress
}
