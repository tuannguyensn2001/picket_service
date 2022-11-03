package config

type config struct {
	Env         string
	GrpcAddress string
}

func GetConfig() config {
	structure := bootstrap()
	result := config{
		Env:         structure.App.Env,
		GrpcAddress: structure.App.GrpcAddress,
	}

	return result
}

func (c config) GetEnv() string {
	return c.Env
}

func (c config) GetGrpcAddress() string {
	return c.GrpcAddress
}
