package custom

import (
	"errors"
	"fmt"
)

type AuthorizationError struct {
	error
}

type InvalidPermmissionsError struct {
	error
}

type BlacklistedNamespaceError struct {
	error
}

func NewAuthorizationError(cause error) AuthorizationError {
	msg := "authorization error"
	if cause != nil {
		return AuthorizationError{fmt.Errorf("%s: %w", msg, cause)}
	}
	return AuthorizationError{errors.New(msg)}
}

func NewInvalidPermissionsError() InvalidPermmissionsError {
	return InvalidPermmissionsError{
		fmt.Errorf("invalid permissions %w", NewAuthorizationError(nil)),
	}
}

func (e InvalidPermmissionsError) Unwrap() error {
	return e.error
}

func NewBlacklistedNamespaceError() BlacklistedNamespaceError {
	return BlacklistedNamespaceError{
		fmt.Errorf("blacklisted namespace %w", NewAuthorizationError(nil)),
	}
}

func (e BlacklistedNamespaceError) Unwrap() error {
	return e.error
}
