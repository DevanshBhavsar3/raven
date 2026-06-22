package message

import (
	"github.com/go-playground/validator/v10"
)

type CreateMessageRequest struct {
	Content string      `json:"content" validate:"required"`
	Role    MessageRole `json:"role" validate:"required,oneof=user agent"`
}

func (r *CreateMessageRequest) Validate(validate *validator.Validate) error {
	return validate.Struct(r)
}
