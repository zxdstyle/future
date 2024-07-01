package local

type Addition struct {
	Path       string `json:"path" required:"true"`
	ShowHidden bool   `json:"show_hidden" default:"true" required:"false" help:"show hidden directories and files"`
}
