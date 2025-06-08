package integration

import (
	"fmt"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration/strategies"
)

// IntegrationFactory creates and returns the appropriate numerical integration strategy
// based on the specified method name.
//
// This factory function implements the Factory design pattern to create different
// integration strategies:
//   - "Gauss-Legendre": Uses Gauss-Legendre quadrature.
//   - "Newton-Cotes": Uses Newton-Cotes formulas (Trapezoidal, Simpson's).
//
// Parameters:
//   - methodName: The name of the integration method ("Gauss-Legendre", "Newton-Cotes").
//
// Returns:
//   - IntegrationStrategy: The configured strategy implementation.
//   - error: An error if the method name is not supported.
func IntegrationFactory(methodName string) (IntegrationStrategy, error) {
	switch methodName {
	case "Gauss-Legendre":
		return &strategies.GaussLegendreStrategy{}, nil
	case "Newton-Cotes":
		return &strategies.NewtonCotesStrategy{}, nil
	default:
		return nil, fmt.Errorf("invalid integration method: %s. Supported methods are 'Gauss-Legendre', 'Newton-Cotes'", methodName)
	}
}
