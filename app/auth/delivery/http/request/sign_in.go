package request

type SignIn struct {
	Username string `json:"username" form:"username" validate:"required,gte=2"`
	Password string `json:"password" form:"password" validate:"required,gte=8"`
}
