package integration

import (
	"fmt"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration/strategies"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integrationcore"
)

// IntegrationFactory creates and returns the appropriate numerical integration strategy
// based on the specified method and order.
//
// This factory function implements the Factory design pattern to create different
// integration strategies.
//
// Parameters:
//   - strategyName: The name of the integration strategy (e.g., "NewtonCotesOrder1", "GaussLegendreOrder2")
//
// Returns:
//   - integrationcore.IntegrationStrategy: The configured strategy implementation
//   - error: An error if the strategyName is not supported
func IntegrationFactory(strategyName string) (integrationcore.IntegrationStrategy, error) {
	switch strategyName {
	// Newton-Cotes Methods
	case "NewtonCotesOrder1":
		return &strategies.NewtonCotesOrder1{}, nil
	case "NewtonCotesOrder2":
		return &strategies.NewtonCotesOrder2{}, nil
	case "NewtonCotesOrder3":
		return &strategies.NewtonCotesOrder3{}, nil
	case "NewtonCotesOrder4":
		return &strategies.NewtonCotesOrder4{}, nil
	// Gauss-Legendre Methods
	case "GaussLegendreOrder1":
		return &strategies.GaussLegendreOrder1{}, nil
	case "GaussLegendreOrder2":
		return &strategies.GaussLegendreOrder2{}, nil
	case "GaussLegendreOrder3":
		return &strategies.GaussLegendreOrder3{}, nil
	case "GaussLegendreOrder4":
		return &strategies.GaussLegendreOrder4{}, nil
	default:
		return nil, fmt.Errorf("invalid integration strategy name: %s", strategyName)
	}
}
