package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
