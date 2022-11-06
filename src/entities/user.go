package entities

type Profile struct {
	Avatar string `json:"avatar"`
}

type User struct {
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Profile  *Profile `json:"profile"`
	Type     int      `json:"type"`
	Status   int      `json:"status"`
}
