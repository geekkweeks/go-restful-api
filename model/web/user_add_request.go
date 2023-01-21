package web

type UserAddRequest struct {
	Username  string `validate:"required, max=200, min=5"`
	FirstName string `validate:"required, max=200, min=5"`
	LastName  string `validate:"required, max=200, min=5"`
	Phone     string `validate:"required, max=50, min=5"`
}
