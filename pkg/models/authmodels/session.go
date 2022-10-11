package authmodels

type Session struct {
	UserId   int    `json:"userId,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}
