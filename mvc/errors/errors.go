package errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorStatus uint

type Error struct {
	Status  ErrorStatus
	Code    string
	Message string
	Caused  error
}

// Define values for status
const (
	InvalidArgs ErrorStatus = iota
	ResourceNotFound
	InternalFailed
)

func NewInvalidArgs(message string, err error) *Error {
	return &Error{
		Status:  InvalidArgs,
		Code:    ErrCodeBadRequestException,
		Message: message,
		Caused:  err,
	}
}

func NewResourceNotFound(message string, err error) *Error {
	return &Error{
		Status:  ResourceNotFound,
		Code:    ErrCodeNotFoundException,
		Message: message,
		Caused:  err,
	}
}

func NewInternalFailed(message string, err error) *Error {
	return &Error{
		Status:  InternalFailed,
		Code:    ErrCodeInternalServerErrorException,
		Message: message,
		Caused:  err,
	}
}
func (e *Error) ReplyStatus() error {
	return status.Errorf(e.Status.ToRpcStatus(), e.ToString())
}

func (e *Error) ToString() string {
	return fmt.Sprintf("%v, %v", e.Code, e.Message)
}

func (e *Error) Error() error {
	return errors.New(e.ToString())
}

func (e ErrorStatus) ToRpcStatus() codes.Code {
	switch e {
	case InvalidArgs:
		return codes.InvalidArgument
	case ResourceNotFound:
		return codes.NotFound
	case InternalFailed:
		return codes.Internal
	default:
		return codes.Internal
	}
}
