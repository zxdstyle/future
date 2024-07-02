package responses

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Bytes(data []byte) Response {
	return &bytesResponse{
		status: http.StatusOK,
		data:   data,
	}
}

type bytesResponse struct {
	status int
	data   []byte
}

func (v *bytesResponse) StatusCode() int {
	return v.status
}

func (v *bytesResponse) SetStatusCode(i int) {
	v.status = i
}

func (v *bytesResponse) Respond(ctx *fiber.Ctx) error {
	_, err := ctx.Write(v.data)
	return err
}
