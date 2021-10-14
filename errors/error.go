package errors

import (
	"errors"
	"fmt"
)

type AuthorizationError error

type InvalidPermmissionsError struct {
	AuthorizationError
}

type BlacklistedNamespaceError struct {
	AuthorizationError
}

func NewAuthorizationError(cause error) AuthorizationError {
	if cause != nil {
		return fmt.Errorf("authorization error: %w", cause)
	}
	return errors.New("authorization error")
}

func NewInvalidPermissionsError() InvalidPermmissionsError {
	return InvalidPermmissionsError{fmt.Errorf("invalid permissions %w", NewAuthorizationError(nil))}
}

func NewBlacklistedNamespaceError() BlacklistedNamespaceError {
	return BlacklistedNamespaceError{fmt.Errorf("blacklisted namespace %w", NewAuthorizationError(nil))}
}
