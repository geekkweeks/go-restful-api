package web

type LoginRequest struct {
	Username string `validate:"required,max=200,min=5" json:"username"`
	Password string `validate:"required,max=20,min=5" json:"password"`
}
