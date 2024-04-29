package responses

import "github.com/gofiber/fiber/v2"

type viewResponse struct {
	status  int
	view    string
	bind    interface{}
	layouts []string
}

func (v *viewResponse) StatusCode() int {
	return v.status
}

func (v *viewResponse) SetStatusCode(i int) {
	v.status = i
}

func (v *viewResponse) Respond(ctx *fiber.Ctx) error {
	return ctx.Render(v.view, v.bind)
}

func View(view string, bind interface{}, layouts ...string) Response {
	return &viewResponse{
		status:  fiber.StatusOK,
		view:    view,
		bind:    bind,
		layouts: layouts,
	}
}

func Abort(code int, message string) Response {
	return &viewResponse{
		status: code,
		view:   "error",
		bind:   map[string]string{"ErrorMsg": message},
	}
}
