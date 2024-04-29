package handler

import (
	"context"
	"future-admin/pkg/server/requests"
	"future-admin/pkg/server/responses"
	"github.com/gofiber/fiber/v2"
)

type Handler func(ctx context.Context, req *requests.RequestAble) (responses.Response, error)

func WrapHandler(handlers ...Handler) (res []fiber.Handler) {
	for _, handler := range handlers {
		res = append(res, func(ctx *fiber.Ctx) error {
			resp, err := handler(ctx.UserContext(), requests.New(ctx))
			if err != nil {
				return err
			}
			ctx = ctx.Status(resp.StatusCode())

			return resp.Respond(ctx)
		})
	}
	return
}
