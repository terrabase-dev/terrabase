package services

import (
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
)

func mapError(err error) error {
	switch {
	case errors.Is(err, repos.ErrNotFound):
		return connect.NewError(connect.CodeNotFound, err)
	case errors.Is(err, repos.ErrAlreadyExists):
		return connect.NewError(connect.CodeAlreadyExists, err)
	case errors.Is(err, repos.ErrNoUpdatesProvided):
		return ErrNoUpdatesProvided
	default:
		return connect.NewError(connect.CodeInternal, err)
	}
}

func fieldRequiredError(fieldName string) error {
	return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("%s is required", fieldName))
}

func internalError(err error) error {
	return connect.NewError(connect.CodeInternal, err)
}

func unknownError(err error) error {
	return connect.NewError(connect.CodeUnknown, err)
}

var (
	ErrIdRequired        = fieldRequiredError("id")
	ErrNameRequired      = fieldRequiredError("name")
	ErrNoUpdatesProvided = connect.NewError(connect.CodeInvalidArgument, repos.ErrNoUpdatesProvided)
	ErrPermissionDenied  = connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	ErrUnauthenticated   = connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
)
