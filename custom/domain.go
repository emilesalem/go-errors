package custom

import (
	"errors"
	"fmt"
)

type DomainError struct {
	error
}

type PeculiarDomainError struct {
	error
	meta struct{ string }
}

func NewDomainError(cause error) DomainError {
	msg := "domain error"
	if cause != nil {
		return DomainError{fmt.Errorf("%s: %w", msg, cause)}
	}
	return DomainError{errors.New(msg)}
}

func NewPeculiarDomainError(peculiarData string) PeculiarDomainError {
	return PeculiarDomainError{
		fmt.Errorf("peculiar problem: %w", NewDomainError(nil)),
		struct{ string }{peculiarData},
	}
}

func (e PeculiarDomainError) Error() string {
	return fmt.Sprintf(`%v: %v`, e.meta, e.error.Error())
}

func (e PeculiarDomainError) Unwrap() error {
	return e.error
}
