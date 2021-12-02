package logger

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestLogrusLogger(t *testing.T) {
	//Given & When
	New(&Config{Name: "gvs-common", Level: "error"})
	//Then
	globalLogger.Logf(ErrorLevel, "logsomething %d", 4711)
	//Given & When
	New(&Config{Name: "gvs-common", Level: "invalid"})
	//When
	assert.Equal(t, "info", logrus.InfoLevel.String())
}
