package web

type UserAddRequest struct {
	Username  string `validate:"required,max=200,min=5" json:"username"`
	FirstName string `validate:"required,max=200,min=5" json:"firstName"`
	LastName  string `validate:"required,max=200,min=5" json:"lastName"`
	Phone     string `validate:"required,max=50,min=5" json:"phone"`
}
