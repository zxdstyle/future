package responses

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Bytes(data bytes.Buffer) Response {
	return &bytesResponse{
		status: http.StatusOK,
		data:   data,
	}
}

type bytesResponse struct {
	status int
	data   bytes.Buffer
}

func (v *bytesResponse) StatusCode() int {
	return v.status
}

func (v *bytesResponse) SetStatusCode(i int) {
	v.status = i
}

func (v *bytesResponse) Respond(ctx *fiber.Ctx) error {
	_, err := ctx.Write(v.data.Bytes())
	return err
}
