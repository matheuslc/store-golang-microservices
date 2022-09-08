package customErrors

import (
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/contracts"
	"github.com/pkg/errors"
	"net/http"
)

func NewUnAuthorizedError(message string) error {
	ue := &unauthorizedError{
		WithStack: NewCustomErrorStack(nil, http.StatusUnauthorized, message),
	}

	return ue
}

func NewUnAuthorizedErrorWrap(err error, message string) error {
	ue := &unauthorizedError{
		WithStack: NewCustomErrorStack(err, http.StatusUnauthorized, message),
	}

	return ue
}

type unauthorizedError struct {
	contracts.WithStack
}

type UnauthorizedError interface {
	contracts.WithStack
	IsUnAuthorizedError() bool
	GetCustomError() CustomError
}

func (u *unauthorizedError) IsUnAuthorizedError() bool {
	return true
}

func (u *unauthorizedError) GetCustomError() CustomError {
	return GetCustomError(u)
}

func IsUnAuthorizedError(err error) bool {
	var unauthorizedError *unauthorizedError
	//us, ok := grpc_errors.Cause(err).(*unauthorizedError)
	if errors.As(err, &unauthorizedError) {
		return unauthorizedError.IsUnAuthorizedError()
	}

	return false
}