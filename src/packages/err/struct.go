package errpkg

import "myclass_service/src/app"

type general struct {
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
	Unauthorized *app.Error `yaml:"unauthorized"`
}

type rootErr struct {
	General *general `yaml:"general"`
}
