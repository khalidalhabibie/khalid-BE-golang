package request

type Update struct {
	Name        *string `json:"name" form:"name" validate:"omitempty"`
	Type        *string `json:"type" form:"type" validate:"omitempty"`
	Description *string `json:"description" form:"description" validate:"omitempty"`
	NakesCount  *uint64 `json:"nakes_count" form:"nakes_count" validate:"omitempty"`
}
