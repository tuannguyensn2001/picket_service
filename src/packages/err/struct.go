package err

import "myclass_service/src/app"

type general struct {
	Forbidden  *app.Error `yaml:"forbidden"`
	BadRequest *app.Error `yaml:"badRequest"`
}

type rootErr struct {
	General *general `yaml:"general"`
}
