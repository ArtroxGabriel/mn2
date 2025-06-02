package common

import "errors"

var (
	ErrInvalidStrategy = errors.New("invalid derivation strategy")
	ErrInvalidDerivate = errors.New("invalid derivation order")
	ErrZeroValue       = errors.New("invalid input, must be greater than zero")
)
