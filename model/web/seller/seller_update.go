package web

type UserUpdateServiceRequest struct {
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required,email" json:"email"`
	//Password string `validate:"required" json:"password"`
}
