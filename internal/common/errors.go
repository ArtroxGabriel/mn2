package common

import "errors"

var (
	ErrInvalidStrategy = errors.New("invalid derivation strategy")
	ErrInvalidDerivate = errors.New("invalid derivation order")
)
