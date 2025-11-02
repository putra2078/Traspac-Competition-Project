package manager

import "errors"

var (
	ErrNIPExists 		= errors.New("NIP already in use")
	ErrEmailExists		= errors.New("Email already in use")
	ErrContactExists 	= errors.New("contact email already in use")
	ErrNilEntity 		= errors.New("manager, contact, user is nil")
)