package entities

type Profile struct {
	Avatar string `json:"avatar"`
}

type User struct {
	Email    string   `json:"email,omitempty"`
	Username string   `json:"username,omitempty"`
	Profile  *Profile `json:"profile,omitempty"`
	Type     int      `json:"type,omitempty"`
	Status   int      `json:"status,omitempty"`
	Id       int      `json:"id,omitempty"`
	Password string   `json:"-"`
}
