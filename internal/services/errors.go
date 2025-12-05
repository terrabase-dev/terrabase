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

var (
	ErrIdRequired        = fieldRequiredError("id")
	ErrNoUpdatesProvided = connect.NewError(connect.CodeInvalidArgument, repos.ErrNoUpdatesProvided)
)
