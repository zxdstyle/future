package handler

import (
	"errors"
	"future-admin/pkg/server/responses"
	va "future-admin/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *fiber.Ctx, er error) error {
	var (
		code = fiber.StatusInternalServerError
		msg  = er.Error()
	)

	var e *fiber.Error
	if errors.As(er, &e) {
		code = e.Code
		msg = e.Message
	}

	if errors.Is(er, gorm.ErrRecordNotFound) {
		code = fiber.StatusNotFound
		msg = er.Error()
	}

	var validatorErr validator.ValidationErrors
	if errors.As(er, &validatorErr) {
		code = fiber.StatusUnprocessableEntity
		msg = validatorErr[0].Translate(va.Trans)
	}

	ctx.Status(code)

	return ctx.JSON(responses.Failed(0, msg))
}
