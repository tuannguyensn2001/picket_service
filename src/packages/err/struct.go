package errpkg

import "myclass_service/src/app"

type auth struct {
	CodeNotValid *app.Error `yaml:"codeNotValid"`
	Unauthorized *app.Error `yaml:"unauthorized"`
}

type general struct {
	Unauthorized *app.Error `yaml:"unauthorized"`
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
}

type rootErr struct {
	Auth    *auth    `yaml:"auth"`
	General *general `yaml:"general"`
}
