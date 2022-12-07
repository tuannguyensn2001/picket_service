package errpkg

import "picket/src/app"

type test struct {
	TestHasContent *app.Error `yaml:"testHasContent"`
}

type answersheet struct {
	UserDoingTest *app.Error `yaml:"userDoingTest"`
	TimeNotValid  *app.Error `yaml:"timeNotValid"`
}

type general struct {
	Unauthorized *app.Error `yaml:"unauthorized"`
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
	NotFound     *app.Error `yaml:"notFound"`
}

type auth struct {
	AccountExisted *app.Error `yaml:"accountExisted"`
	LoginFailed    *app.Error `yaml:"loginFailed"`
	CodeNotValid   *app.Error `yaml:"codeNotValid"`
	Unauthorized   *app.Error `yaml:"unauthorized"`
}

type rootErr struct {
	Test        *test        `yaml:"test"`
	Answersheet *answersheet `yaml:"answersheet"`
	General     *general     `yaml:"general"`
	Auth        *auth        `yaml:"auth"`
}
