// Package integration implements numerical integration methods using various strategies.
// It provides a unified interface for calculating definite integrals.
package integration

import (
	"fmt" // Import fmt for error wrapping, if needed for more specific errors
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

// Integrator is the main struct that encapsulates an integration strategy.
// It provides a high-level interface for numerical integration calculations.
type Integrator struct {
	// strategy holds the specific numerical integration strategy to be used
	strategy IntegrationStrategy
}

// NewIntegrator creates a new Integrator instance with the specified method type and order.
//
// Parameters:
//   - methodType: The integration method ("NewtonCotes" or "GaussLegendre")
//   - order: The order of the method.
//     For NewtonCotes, this refers to the type of rule (1: Trapezoidal, 2: Simpson's 1/3, etc.).
//     For GaussLegendre, this refers to the number of points (1 to 4).
//
// Returns:
//   - *Integrator: A configured integrator instance
//   - error: An error if the methodType/order combination is invalid (from IntegrationFactory)
func NewIntegrator(methodType string, order int) (*Integrator, error) {
	strat, err := IntegrationFactory(methodType, order)
	if err != nil {
		// The error from IntegrationFactory already wraps common.ErrInvalidStrategy
		return nil, err
	}
	return &Integrator{strategy: strat}, nil
}

// Calculate computes the definite integral of the given function from a to b
// using n subintervals or points, based on the configured strategy.
//
// Parameters:
//   - fn: The function to integrate (common.MathFunc)
//   - a: The lower limit of integration
//   - b: The upper limit of integration
//   - n: The number of subintervals (for Newton-Cotes methods) or points (for Gauss-Legendre methods).
//        - For Newton-Cotes methods (Trapezoidal, Simpson's, Boole's): 'n' represents the
//          number of subintervals and must satisfy the specific requirements of the rule
//          (e.g., n >= 1 for Trapezoidal, n must be even for Simpson's 1/3, multiple of 3 for Simpson's 3/8, multiple of 4 for Boole's).
//        - For Gauss-Legendre methods: The strategy is already fixed to a certain number of points
//          (e.g., GaussLegendreOrder2 uses 2 points). The 'n' parameter passed here is
//          effectively ignored by the Gauss-Legendre strategy's Integrate method, as the
//          number of points is intrinsic to the strategy type chosen by 'order' in NewIntegrator.
//          It is part of the common IntegrationStrategy interface signature.
//
// Returns:
//   - float64: The calculated integral value
//   - error: An error if the integration fails (e.g., invalid 'n' for the chosen Newton-Cotes rule,
//            or other issues during calculation).
func (itg *Integrator) Calculate(fn common.MathFunc, a, b float64, n int) (float64, error) {
	if itg.strategy == nil {
		// This indicates NewIntegrator was not called or failed, and Calculate was somehow invoked on a zero Integrator struct.
		return 0, fmt.Errorf("%w: Integrator strategy not initialized. Call NewIntegrator first", common.ErrInvalidStrategy)
	}
	return itg.strategy.Integrate(fn, a, b, n)
}
