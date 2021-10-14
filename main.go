package main

import (
	_errors "emile/test/go-errors/errors"
	"errors"
	"fmt"
)

func main() {
	err := _errors.NewInvalidPermissionsError()

	if errors.As(err, new(_errors.AuthorizationError)) {
		fmt.Printf("1st error : %v\n", err)
	} else {
		fmt.Println("unknown error")
	}

	cause := errors.New("some other error")

	err2 := _errors.NewAuthorizationError(cause)

	fmt.Printf("2nd error: %v\n", err2)
}
