package request

type SignUpConfirmation struct {
	Email string `json:"email" form:"email" validate:"required,email"`
	Code  string `json:"code" form:"code" validate:"required"`
}
