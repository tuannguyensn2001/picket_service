package auth_struct

type LoginGoogleOutput struct {
	AccessToken string
}

type RegisterInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Username string `validate:"required"`
}

type LoginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type LoginOutput struct {
	AccessToken string
}
