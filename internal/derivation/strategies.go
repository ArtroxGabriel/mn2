// Package derivation defines interfaces and strategy implementations for numerical differentiation.
// This file contains the strategy interface and compile-time interface compliance checks.
package derivation

import (
	firstorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/first-order"
	fourthorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/fourth-order"
	secondorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/second-order"
	thirdorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/third-order"
)

// DerivationStrategy defines the interface that all numerical differentiation strategies must implement.
// This interface follows the Strategy design pattern, allowing different finite difference methods
// to be used interchangeably.
//
// All methods use finite difference approximations to calculate derivatives:
// - CalculateFirst: Computes the first derivative (f'(x))
// - CalculateSecond: Computes the second derivative (f”(x))
// - CalculateThirty: Computes the third derivative (f”'(x))
//
// The accuracy and computational requirements vary depending on the specific strategy implementation.
type DerivationStrategy interface {
	// CalculateFirst computes the first derivative of the function at point x using step size h.
	//
	// Parameters:
	//   - fn: The function to differentiate
	//   - x: The point at which to calculate the derivative
	//   - h: The step size for finite difference approximation
	//
	// Returns:
	//   - float64: The approximated first derivative value f'(x)
	CalculateFirst(
		fn func(float64) float64,
		x, h float64,
	) float64

	// CalculateSecond computes the second derivative of the function at point x using step size h.
	//
	// Parameters:
	//   - fn: The function to differentiate
	//   - x: The point at which to calculate the derivative
	//   - h: The step size for finite difference approximation
	//
	// Returns:
	//   - float64: The approximated second derivative value f''(x)
	CalculateSecond(
		fn func(float64) float64,
		x, h float64,
	) float64

	// CalculateThirty computes the third derivative of the function at point x using step size h.
	//
	// Parameters:
	//   - fn: The function to differentiate
	//   - x: The point at which to calculate the derivative
	//   - h: The step size for finite difference approximation
	//
	// Returns:
	//   - float64: The approximated third derivative value f'''(x)
	CalculateThirty(
		fn func(float64) float64,
		x, h float64,
	) float64
}

// Compile-time interface compliance checks.
// These variable declarations ensure that all strategy implementations
// correctly implement the DerivationStrategy interface. If any implementation
// is missing required methods, the compilation will fail.
var (
	// Forward difference strategies
	_ DerivationStrategy = (*firstorder.ForwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*secondorder.ForwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*thirdorder.ForwardThirOrderStrategy)(nil)

	// Backward difference strategies
	_ DerivationStrategy = (*firstorder.BackwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*secondorder.BackwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*thirdorder.BackwardThirOrderStrategy)(nil)

	// Central difference strategies
	_ DerivationStrategy = (*secondorder.CentralSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*fourthorder.CentralFourthOrderStrategy)(nil)
)
