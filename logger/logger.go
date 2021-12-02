package logger

// Logger  is an interface for logging methods
type Logger interface {
	Load(options ...Option) error
	Init() error
	Logf(level Level, format string, v ...interface{})
}

// logger global instance
var globalLogger Logger

func Tracef(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(TraceLevel, format, v)
	}
}

func Debugf(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(DebugLevel, format, v)
	}
}

func Infof(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(InfoLevel, format, v)
	}
}

func Warnf(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(WarnLevel, format, v)
	}
}

func Errorf(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(ErrorLevel, format, v)
	}
}

func Fatalf(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(FatalLevel, format, v)
	}
}

func Panicf(format string, v ...interface{}) {
	if globalLogger != nil {
		globalLogger.Logf(PanicLevel, format, v)
	}
}

// New creates the logger.
func New(cfg *Config) {

	var opts = []Option{
		WithLevel(ToLevel(cfg.Level)),
		WithFormat(ToFormat(cfg.Format)),
		WithOutput(ToOutput(cfg.Output)),
		WithFields(map[string]interface{}{
			"name": cfg.Name,
		}),
	}
	newLogger(&logrusx{}, opts...)
}

func newLogger(source Logger, opts ...Option) {
	err := source.Load(opts...)
	if err != nil {
		panic(err)
	}
	if err := source.Init(); err != nil {
		panic(err)
	}
	// initialize the global logger instance.
	globalLogger = source
}
