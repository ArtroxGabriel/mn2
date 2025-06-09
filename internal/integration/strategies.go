package integration

import (
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	gausslegendre "github.com/ArtroxGabriel/numeric-methods-cli/internal/integration/strategies/gauss-legendre"
	newtoncotes "github.com/ArtroxGabriel/numeric-methods-cli/internal/integration/strategies/newton-cotes"
)

// IntegrationStrategy defines the interface that all numerical integration strategies must implement.
// This interface follows the Strategy design pattern, allowing different integration methods
// to be used interchangeably.
type IntegrationStrategy interface {
	// Integrate computes the definite integral of the function fn from a to b
	// using n subintervals or points.
	//
	// Parameters:
	//   - fn: The function to integrate (common.MathFunc)
	//   - a: The lower limit of integration
	//   - b: The upper limit of integration
	//   - n: The number of subintervals (for Newton-Cotes) or points (for Gauss-Legendre)
	//
	// Returns:
	//   - float64: The approximated integral value
	//   - error: An error if the integration cannot be performed (e.g., invalid parameters)
	Integrate(fn common.MathFunc, a, b float64, n int) (float64, error)
}

// Compile-time interface compliance checks.
// These variable declarations ensure that all strategy implementations
// correctly implement the IntegrationStrategy interface. If any implementation
// is missing required methods, the compilation will fail.
var (
	// Newton-Cotes strategies
	_ IntegrationStrategy = (*newtoncotes.TrapezoidalRule)(nil)
	_ IntegrationStrategy = (*newtoncotes.SimpsonsOneThirdRule)(nil)
	_ IntegrationStrategy = (*newtoncotes.SimpsonsThreeEighthRule)(nil)
	_ IntegrationStrategy = (*newtoncotes.BoolesRule)(nil)

	// Gauss-Legendre strategies
	_ IntegrationStrategy = (*gausslegendre.GaussLegendreOrder1)(nil)
	_ IntegrationStrategy = (*gausslegendre.GaussLegendreOrder2)(nil)
	_ IntegrationStrategy = (*gausslegendre.GaussLegendreOrder3)(nil)
	_ IntegrationStrategy = (*gausslegendre.GaussLegendreOrder4)(nil)
)
