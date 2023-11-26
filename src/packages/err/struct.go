package errpkg

import "picket/src/app"

type auth struct {
	CodeNotValid   *app.Error `yaml:"codeNotValid"`
	Unauthorized   *app.Error `yaml:"unauthorized"`
	AccountExisted *app.Error `yaml:"accountExisted"`
	LoginFailed    *app.Error `yaml:"loginFailed"`
}

type test struct {
	TestHasContent   *app.Error `yaml:"testHasContent"`
	QuestionNotValid *app.Error `yaml:"questionNotValid"`
	TestNotValid     *app.Error `yaml:"testNotValid"`
}

type answersheet struct {
	UserDoingTest *app.Error `yaml:"userDoingTest"`
	TimeNotValid  *app.Error `yaml:"timeNotValid"`
}

type general struct {
	BadRequest   *app.Error `yaml:"badRequest"`
	NotFound     *app.Error `yaml:"notFound"`
	Unauthorized *app.Error `yaml:"unauthorized"`
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
}

type rootErr struct {
	Auth        *auth        `yaml:"auth"`
	Test        *test        `yaml:"test"`
	Answersheet *answersheet `yaml:"answersheet"`
	General     *general     `yaml:"general"`
}
