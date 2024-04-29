package env

import "os"

const (
	Pro         = "pro"
	Pre         = "pre"
	Test        = "test"
	Development = "development"
	Local       = "local"
)

// Environment ...
func Environment() string {
	return os.Getenv("APP_ENV")
}

// IsProduction ...
func IsProduction() bool {
	return Environment() == Pro
}

// IsPre ...
func IsPre() bool {
	return Environment() == Pre
}

// IsTest ...
func IsTest() bool {
	return Environment() == Test
}

// IsDevelopment ...
func IsDevelopment() bool {
	return Environment() == Development
}

// IsLocal ...
func IsLocal() bool {
	return Environment() == Local
}
