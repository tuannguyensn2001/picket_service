package errpkg

import "myclass_service/src/app"

type auth struct {
	Unauthorized   *app.Error `yaml:"unauthorized"`
	AccountExisted *app.Error `yaml:"accountExisted"`
	LoginFailed    *app.Error `yaml:"loginFailed"`
	CodeNotValid   *app.Error `yaml:"codeNotValid"`
}

type test struct {
	TestHasContent *app.Error `yaml:"testHasContent"`
}

type general struct {
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
	NotFound     *app.Error `yaml:"notFound"`
	Unauthorized *app.Error `yaml:"unauthorized"`
}

type rootErr struct {
	Auth    *auth    `yaml:"auth"`
	Test    *test    `yaml:"test"`
	General *general `yaml:"general"`
}
