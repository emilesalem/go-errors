package custom_test

import (
	"emile/test/go-errors/custom"
	"errors"
	"fmt"
)

func ExampleNewInvalidPermissionsError() {
	var target custom.AuthorizationError

	err := custom.NewInvalidPermissionsError()

	if errors.As(err, &target) {
		fmt.Println(`As authorization error`)
	} else if errors.As(err, new(custom.DomainError)) {
		fmt.Printf(`As domain error `)
	} else {
		fmt.Println(`unknown error`)
	}
	// Output: As authorization error
}

func ExampleNewBlacklistedNamespaceError() {
	var target1 custom.AuthorizationError
	var target2 custom.DomainError

	err := custom.NewBlacklistedNamespaceError()

	if errors.As(err, &target1) {
		fmt.Println(`As authorization error`)
	} else if errors.As(err, &target2) {
		fmt.Printf(`As domain error `)
	} else {
		fmt.Println(`unknown error`)
	}
	// Output: As authorization error
}

func ExampleNewPeculiarDomainError() {
	var target custom.DomainError

	err := custom.NewPeculiarDomainError(`peculiar data`)

	if errors.As(err, new(custom.AuthorizationError)) {
		fmt.Println(`As authorization error`)
	} else if errors.As(err, &target) {
		fmt.Printf(`As domain error: %v`, err)
	} else {
		fmt.Println(`unknown error`)
	}
	// Output: As domain error: {peculiar data}: peculiar problem: domain error
}

func ExampleNewDomainError() {
	err := custom.NewDomainError(errors.New("error cause"))

	if errors.As(err, new(custom.AuthorizationError)) {
		fmt.Println(`As authorization error`)
	} else if errors.As(err, new(custom.DomainError)) {
		fmt.Printf(`As domain error `)
	} else {
		fmt.Println(`unknown error`)
	}
	// Output: As domain error
}

func ExampleNewAuthorizationError() {
	err := custom.NewAuthorizationError(errors.New("error cause"))

	if errors.As(err, new(custom.AuthorizationError)) {
		fmt.Println(`As authorization error`)
	} else if errors.As(err, new(custom.DomainError)) {
		fmt.Printf(`As domain error `)
	} else {
		fmt.Println(`unknown error`)
	}
	// Output: As authorization error
}
