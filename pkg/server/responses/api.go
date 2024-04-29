package responses

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	response struct {
		status  int
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    *Meta       `json:"meta,omitempty"`
	}

	Meta struct {
		Paginator *Paginator `json:"paginator,omitempty"`
	}

	Paginator struct {
		CurrentPage int   `json:"current_page"`
		PerPage     int   `json:"per_page"`
		Total       int64 `json:"total"`
	}
)

func (r *response) StatusCode() int {
	if r.status > 0 {
		return r.status
	}
	return http.StatusOK
}

func (r *response) SetStatusCode(status int) {
	r.status = status
}

func (r *response) Respond(ctx *fiber.Ctx) error {
	return ctx.JSON(r)
}

func Failed(code int, message string) Response {
	return &response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func Success(data interface{}) Response {
	return &response{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func Empty() Response {
	return &response{
		status:  http.StatusNoContent,
		Code:    0,
		Message: "success",
		Data:    nil,
	}
}

func PaginatorList(data interface{}, current, pageSize int, total int64) Response {
	return &response{
		Code:    0,
		Message: "success",
		Data:    data,
		Meta: &Meta{
			Paginator: &Paginator{
				CurrentPage: current,
				PerPage:     pageSize,
				Total:       total,
			},
		},
	}
}
