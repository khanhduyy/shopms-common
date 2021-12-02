package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logrusx implements the Logger
type logrusx struct {
	opts option
}

// Formatter adapt to add additional fields
type Formatter struct {
	// Global customized fields
	fields map[string]interface{}
	// Source Formatter
	formatter logrus.Formatter
}

// Format customizes format entries before writing them on the target logger output.
func (l Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Additional the global fields
	for k, v := range l.fields {
		entry.Data[k] = v
	}
	return l.formatter.Format(entry)
}

func (l *logrusx) Load(options ...Option) error {
	for _, opt := range options {
		opt(&l.opts)
	}
	return nil
}

func (l *logrusx) Logf(level Level, format string, value ...interface{}) {
	switch level {
	case TraceLevel:
		logrus.Tracef(format, value...)
	case DebugLevel:
		logrus.Debugf(format, value...)
	case InfoLevel:
		logrus.Infof(format, value...)
	case WarnLevel:
		logrus.Warnf(format, value...)
	case ErrorLevel:
		logrus.Errorf(format, value...)
	case FatalLevel:
		logrus.Fatalf(format, value...)
	case PanicLevel:
		logrus.Panicf(format, value...)
	}
}

func (l *logrusx) Init() error {
	if level, err := logrus.ParseLevel(l.opts.level.String()); err != nil {
		// fallback to default INFO
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}

	switch l.opts.output {
	case StdErr:
		logrus.SetOutput(os.Stderr)
	case Std:
		logrus.SetOutput(os.Stdout)
	default:
		logrus.SetOutput(os.Stdout)
	}

	var fm logrus.Formatter
	switch l.opts.format {
	case JsonFormat:
		fm = new(logrus.JSONFormatter)
	default:
		fm = &logrus.TextFormatter{FullTimestamp: true}
	}

	logrus.SetFormatter(Formatter{
		fields:    l.opts.fields,
		formatter: fm,
	})

	return nil
}
