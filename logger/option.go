package logger

// option presents the logger options.
type option struct {
	level  Level
	format Format
	output Output
	fields map[string]interface{}
}

// Option customizes the logger options.
type Option func(opt *option)

// WithLevel set the level for the logger
func WithLevel(level Level) Option {
	return func(opt *option) {
		opt.level = level
	}
}

// WithFormat set the format for the logger
func WithFormat(format Format) Option {
	return func(opt *option) {
		opt.format = format
	}
}

// WithOutput set the output for the logger
func WithOutput(output Output) Option {
	return func(opt *option) {
		opt.output = output
	}
}

// WithFields set the customized fields for the logger
func WithFields(fields map[string]interface{}) Option {
	return func(opt *option) {
		opt.fields = fields
	}
}

type Level int8

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}
	return ""
}

func ToLevel(level string) Level {
	switch level {
	case TraceLevel.String():
		return TraceLevel
	case DebugLevel.String():
		return DebugLevel
	case InfoLevel.String():
		return InfoLevel
	case WarnLevel.String():
		return WarnLevel
	case ErrorLevel.String():
		return ErrorLevel
	case FatalLevel.String():
		return FatalLevel
	case PanicLevel.String():
		return PanicLevel
	}
	return InfoLevel
}

type Format int8

const (
	JsonFormat Format = iota
	TextFormat
)

func (f Format) String() string {
	switch f {
	case JsonFormat:
		return "json"
	case TextFormat:
		return "text"
	}
	return ""
}

func ToFormat(format string) Format {
	switch format {
	case JsonFormat.String():
		return JsonFormat
	case TextFormat.String():
		return TextFormat
	}
	return TextFormat
}

type Output int8

const (
	Std Output = iota
	StdErr
)

func (out Output) String() string {
	switch out {
	case Std:
		return "sdt"
	case StdErr:
		return "stderr"
	}
	return ""
}

func ToOutput(output string) Output {
	switch output {
	case Std.String():
		return Std
	case StdErr.String():
		return StdErr
	}
	return Std
}
