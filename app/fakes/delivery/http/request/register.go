package request

type Register struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Type        string `json:"type" form:"type" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	NakesCount  uint64 `json:"nakes_count" form:"nakes_count" validate:"required"`
}
