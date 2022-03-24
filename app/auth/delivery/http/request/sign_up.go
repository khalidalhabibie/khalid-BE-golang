package request

type SignUp struct {
	Username   string `json:"username" form:"username" validate:"required,gte=2"`
	Email      string `json:"email" form:"email" validate:"required,email"`
	Password   string `json:"password" form:"password" validate:"required,gte=8"`
	Repassword string `json:"repassword" form:"repassword" validate:"required,gte=8"`
}
