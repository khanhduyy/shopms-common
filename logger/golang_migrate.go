package logger

// golangMigrateLogger provides customized the default golang-migrate logging function.
type golangMigrateLogger struct {
}

func NewGolangMigrateLogger() *golangMigrateLogger {
	return &golangMigrateLogger{}
}

func (gl *golangMigrateLogger) Printf(format string, v ...interface{}) {
	Infof(format, v)
}

func (gl *golangMigrateLogger) Verbose() bool {
	return true
}
