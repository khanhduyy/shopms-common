package logger

// Config describes the logger YAML configuration properties.
type Config struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Output string `yaml:"output"`
	Name   string `yaml:"name"`
}
