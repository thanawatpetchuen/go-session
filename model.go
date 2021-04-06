package gos

type User struct {
	Username string `json:"username"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
