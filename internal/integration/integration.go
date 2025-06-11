package integration

import (
	"fmt"
)

// Integrator is the main struct that encapsulates an integration strategy.
// It provides a high-level interface for numerical integration calculations.
type Integrator struct {
	strategy IntegrationStrategy
}

// NewIntegrator creates a new Integrator instance with the specified strategy name.
//
// Parameters:
//   - strategyName: The name of the integration strategy to use (e.g., "NewtonCotesOrder1", "GaussLegendreOrder2")
//
// Returns:
//   - *Integrator: A configured integrator instance
//   - error: An error if the strategyName is invalid
func NewIntegrator(strategyEnum int) (*Integrator, error) {
	fmt.Println("selected strategy:", strategyEnum)
	var strategyName string
	switch strategyEnum {
	case 0:
		strategyName = "NewtonCotesOrder1"
	case 1:
		strategyName = "NewtonCotesOrder2"
	case 2:
		strategyName = "NewtonCotesOrder3"
	case 3:
		strategyName = "NewtonCotesOrder4"
	case 4:
		strategyName = "GaussLegendreOrder1"
	case 5:
		strategyName = "GaussLegendreOrder2"
	case 6:
		strategyName = "GaussLegendreOrder3"
	case 7:
		strategyName = "GaussLegendreOrder4"
	}

	strategy, err := IntegrationFactory(strategyName)
	if err != nil {
		return nil, fmt.Errorf("failed to create integrator: %w", err)
	}
	return &Integrator{strategy: strategy}, nil
}

// Calculate computes the definite integral of the given function from a to b
// using the configured strategy. For Newton-Cotes methods, n represents the number of subintervals.
// For Gauss-Legendre methods, n is typically not used for the basic implementation but is included
// for interface consistency (it could be used for composite Gauss-Legendre rules if implemented).
//
// Parameters:
//   - fn: The function to integrate (must be a single-variable function)
//   - a: The lower limit of integration
//   - b: The upper limit of integration
//   - n: The number of subintervals (for Newton-Cotes) or order/points (potentially for Gauss-Legendre extensions)
//
// Returns:
//   - float64: The calculated definite integral value
//   - error: An error if the calculation fails
func (it *Integrator) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if it.strategy == nil {
		return 0, fmt.Errorf("integrator strategy is not initialized")
	}
	return it.strategy.Calculate(fn, a, b, n)
}
