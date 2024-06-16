package model

type (
	Option struct {
		Local *Local `json:"local,omitempty"`
	}

	Local struct {
		Folder string `json:"folder"`
	}
)
