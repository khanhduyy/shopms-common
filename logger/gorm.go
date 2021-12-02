package logger

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// gormLogger provides customized the default gorm logging function.
type gormLogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

// NewGorm creates the gorm logger
func NewGorm() *gormLogger {
	return &gormLogger{
		SkipErrRecordNotFound: true,
	}
}

func (l *gormLogger) LogMode(glogger.LogLevel) glogger.Interface {
	return l
}

func (l *gormLogger) Info(_ context.Context, s string, args ...interface{}) {
	Infof(s, args)
}

func (l *gormLogger) Warn(_ context.Context, s string, args ...interface{}) {
	Warnf(s, args)
}

func (l *gormLogger) Error(_ context.Context, s string, args ...interface{}) {
	Errorf(s, args)
}

func (l *gormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}

	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		Errorf("%s [%s]", sql, elapsed)
		return
	}
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		Warnf("%s [%s]", sql, elapsed)
		return
	}
	Debugf("%s [%s]", sql, elapsed)
}
