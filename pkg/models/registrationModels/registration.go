package registrationModels

type Register struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
