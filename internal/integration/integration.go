package integration

import "fmt"

// Integrator is the main struct that encapsulates an integration strategy.
// It provides a high-level interface for numerical integration calculations.
type Integrator struct {
	strategy IntegrationStrategy
}

// NewIntegrator creates a new Integrator instance with the specified method.
//
// Parameters:
//   - methodName: The integration method to use ("Gauss-Legendre", "Newton-Cotes").
//
// Returns:
//   - *Integrator: A configured integrator instance.
//   - error: An error if the method name is invalid or the factory fails.
func NewIntegrator(methodName string) (*Integrator, error) {
	strategy, err := IntegrationFactory(methodName)
	if err != nil {
		return nil, fmt.Errorf("failed to create integration strategy for method '%s': %w", methodName, err)
	}
	return &Integrator{strategy: strategy}, nil
}

// Calculate computes the definite integral of the given function fn from a to b,
// using the configured strategy, with n points/degree and specified tolerance.
//
// Parameters:
//   - fn: The function to integrate (must be a single-variable function).
//   - a: The lower limit of integration.
//   - b: The upper limit of integration.
//   - n: The number of points or degree of the polynomial, depending on the method.
//   - tol: The desired error tolerance.
//
// Returns:
//   - float64: The calculated integral value.
//   - error: An error if the calculation fails (e.g., invalid parameters, convergence issues).
func (itg *Integrator) Calculate(
	fn func(float64) float64,
	a, b float64,
	n int,
	tol float64,
) (float64, error) {
	if fn == nil {
		return 0, fmt.Errorf("function to integrate cannot be nil")
	}
	if itg.strategy == nil {
		return 0, fmt.Errorf("integrator has no strategy set") // Should not happen if NewIntegrator is used
	}

	result, err := itg.strategy.Calculate(fn, a, b, n, tol)
	if err != nil {
		return 0, fmt.Errorf("error during integration calculation with strategy: %w", err)
	}
	return result, nil
}
