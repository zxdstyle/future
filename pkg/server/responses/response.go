package responses

import "github.com/gofiber/fiber/v2"

type Response interface {
	StatusCode() int
	SetStatusCode(int)
	Respond(ctx *fiber.Ctx) error
}
