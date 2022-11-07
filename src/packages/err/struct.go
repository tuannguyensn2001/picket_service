package errpkg

import "myclass_service/src/app"

type general struct {
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
	Unauthorized *app.Error `yaml:"unauthorized"`
}

type auth struct {
	CodeNotValid   *app.Error `yaml:"codeNotValid"`
	Unauthorized   *app.Error `yaml:"unauthorized"`
	AccountExisted *app.Error `yaml:"accountExisted"`
	LoginFailed    *app.Error `yaml:"loginFailed"`
}

type rootErr struct {
	General *general `yaml:"general"`
	Auth    *auth    `yaml:"auth"`
}
