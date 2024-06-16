package model

type (
	IDriver interface {
		List() error
		Config() DriverConfig
		GetAddition() any
	}

	DriverConfig struct {
		Name string
		Slug string
	}

	DriverInfo struct {
		Slug      string               `json:"slug"`
		Name      string               `json:"name"`
		Additions []DriverAdditionItem `json:"additions"`
	}

	DriverAdditionItem struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Default  string `json:"default"`
		Options  string `json:"options"`
		Required bool   `json:"required"`
		Help     string `json:"help"`
	}
)
