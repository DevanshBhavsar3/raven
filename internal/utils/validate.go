package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Valitable interface {
	Validate(*validator.Validate) error
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func BindAndValidate(r *http.Request, payload Valitable) error {
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return fmt.Errorf("error decoding request body: %v", err)
	}

	if err := payload.Validate(validate); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	return nil
}
