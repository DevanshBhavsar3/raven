package models

import "github.com/go-playground/validator/v10"

type CreateConversationRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

func (c *CreateConversationRequest) Validate(validate *validator.Validate) error {
	return validate.Struct(c)
}
