package request

type SignInConfirmation struct {
	Username string `json:"username" form:"username" validate:"required,gte=2"`
	Code  string `json:"code" form:"code" validate:"required"`
}
