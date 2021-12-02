package grpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInternalFailed(t *testing.T) {
	//Given & When
	err := InternalFailed("Failed")
	//Then
	assert.Equal(t, "rpc error: code = Internal desc = Failed", err.Error())
}
