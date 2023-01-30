package govalidator

import (
	"context"
	gov "github.com/go-playground/validator/v10"
)

type Validator struct {
	val *gov.Validate
}

func (v *Validator) Validate(ctx context.Context, s any) error {
	return v.val.StructCtx(ctx, s)
}

func New() *Validator {
	return &Validator{gov.New()}
}
