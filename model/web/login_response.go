package web

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
