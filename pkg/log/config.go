package log

type Config struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
	Name              string // service name
	Writers           string
	LoggerDir         string
	LogFormatText     bool
	LogRollingPolicy  string
	LogBackupCount    uint
}

type Option func(*Config)

// WithFilename set log filename
func WithFilename(filename string) Option {
	return func(cfg *Config) {
		cfg.Name = filename
	}
}

// WithLogDir set log dir
func WithLogDir(dir string) Option {
	return func(cfg *Config) {
		cfg.LoggerDir = dir
	}
}
