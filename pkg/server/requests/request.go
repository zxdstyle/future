package requests

import (
	"future-admin/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	RequestAble struct {
		*fiber.Ctx
	}
)

func New(ctx *fiber.Ctx) *RequestAble {
	return &RequestAble{
		Ctx: ctx,
	}
}

func (r *RequestAble) Validate(pointer interface{}) error {
	if err := r.Ctx.BodyParser(pointer); err != nil {
		return err
	}

	return validator.Validate(pointer)
}
