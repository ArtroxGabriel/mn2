// Package derivation implements numerical differentiation methods using various finite difference strategies.
// It provides a unified interface for calculating first, second, and third derivatives using
// forward, backward, and central difference methods with different orders of accuracy.
package derivation

import (
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

// Derivator is the main struct that encapsulates a derivation strategy.
// It provides a high-level interface for numerical differentiation calculations.
type Derivator struct {
	// derivator holds the specific numerical differentiation strategy to be used
	derivator DerivationStrategy
}

// NewDerivator creates a new Derivator instance with the specified strategy and accuracy order.
//
// Parameters:
//   - strategy: The finite difference method to use ("Forward", "Backward", or "Central")
//   - order: The order of accuracy for the method (1, 2, 3, or 4 depending on the strategy)
//
// Returns:
//   - *Derivator: A configured derivator instance
//   - error: An error if the strategy/order combination is invalid
//
// Example:
//
//	derivator, err := NewDerivator("Central", 2)
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewDerivator(strategy string, order uint) (*Derivator, error) {
	derivator := &Derivator{}

	deriv, err := DerivacaoFactory(strategy, order)
	if err != nil {
		return nil, err
	}
	derivator.derivator = deriv

	return derivator, nil
}

// Calculate computes the numerical derivative of the given function at point x using step size h.
//
// Parameters:
//   - fn: The function to differentiate (must be a single-variable function)
//   - x: The point at which to calculate the derivative
//   - h: The step size for the finite difference approximation
//   - derivate: The order of the derivative to calculate (1 for first, 2 for second, 3 for third)
//
// Returns:
//   - float64: The calculated derivative value
//   - error: An error if the derivative order is invalid (only 1, 2, 3 are supported)
//
// Example:
//
//	// Calculate first derivative of f(x) = x² at x = 2 with h = 0.01
//	fn := func(x float64) float64 { return x * x }
//	result, err := derivator.Calculate(fn, 2.0, 0.01, 1)
func (c *Derivator) Calculate(
	fn func(float64) float64,
	x, h float64,
	derivate int,
) (float64, error) {
	var result float64
	switch derivate {
	case 1:
		result = c.derivator.CalculateFirst(fn, x, h)
	case 2:
		result = c.derivator.CalculateSecond(fn, x, h)
	case 3:
		result = c.derivator.CalculateThirty(fn, x, h)
	default:
		return 0, common.ErrInvalidDerivate
	}

	return result, nil
}
