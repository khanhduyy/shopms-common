package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InternalFailed provides the internal status error on given message.
func InternalFailed(message string) error {

	return status.Errorf(codes.Internal, message)
}
