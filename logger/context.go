package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

const (
	UserIdKey = "sso_id"
)

type (
	user struct{}
)

func NewUserCtx(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, user{}, userId)
}

// WithContext Use context create entry
func WithContext(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}

	if v := FromUserIdContext(ctx); v != "" {
		fields[UserIdKey] = v
	}

	return logrus.WithContext(ctx).WithFields(fields)
}

func FromUserIdContext(ctx context.Context) string {
	v := ctx.Value(user{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
