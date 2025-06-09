package integration

import (
	"fmt"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common" // For common.ErrInvalidStrategy
	gausslegendre "github.com/ArtroxGabriel/numeric-methods-cli/internal/integration/strategies/gauss-legendre"
	newtoncotes "github.com/ArtroxGabriel/numeric-methods-cli/internal/integration/strategies/newton-cotes"
)

// IntegrationFactory creates and returns the appropriate numerical integration strategy
// based on the specified method type and order.
//
// Parameters:
//   - methodType: The integration method ("NewtonCotes" or "GaussLegendre")
//   - order: The order of the method.
//     For NewtonCotes:
//       - 1: TrapezoidalRule (n=subintervals, but order=1 implies the rule itself)
//       - 2: SimpsonsOneThirdRule (order=2 implies the rule)
//       - 3: SimpsonsThreeEighthRule (order=3 implies the rule)
//       - 4: BoolesRule (order=4 implies the rule)
//     For GaussLegendre:
//       - 1: GaussLegendreOrder1 (n=points, order=1 means 1 point)
//       - 2: GaussLegendreOrder2 (order=2 means 2 points)
//       - 3: GaussLegendreOrder3 (order=3 means 3 points)
//       - 4: GaussLegendreOrder4 (order=4 means 4 points)
//
// Returns:
//   - IntegrationStrategy: The configured strategy implementation
//   - error: An error if the methodType/order combination is not supported
func IntegrationFactory(methodType string, order int) (IntegrationStrategy, error) {
	var strategy IntegrationStrategy
	var specificError error // Renamed to avoid confusion when wrapping

	switch methodType {
	case "NewtonCotes":
		switch order {
		case 1:
			strategy = &newtoncotes.TrapezoidalRule{}
		case 2:
			strategy = &newtoncotes.SimpsonsOneThirdRule{}
		case 3:
			strategy = &newtoncotes.SimpsonsThreeEighthRule{}
		case 4:
			strategy = &newtoncotes.BoolesRule{}
		default:
			specificError = fmt.Errorf("unsupported order %d for Newton-Cotes method. Supported orders are 1, 2, 3, 4", order)
		}
	case "GaussLegendre":
		switch order {
		case 1:
			strategy = &gausslegendre.GaussLegendreOrder1{}
		case 2:
			strategy = &gausslegendre.GaussLegendreOrder2{}
		case 3:
			strategy = &gausslegendre.GaussLegendreOrder3{}
		case 4:
			strategy = &gausslegendre.GaussLegendreOrder4{}
		default:
			specificError = fmt.Errorf("unsupported order %d for Gauss-Legendre method. Supported orders are 1, 2, 3, 4", order)
		}
	default:
		specificError = fmt.Errorf("unknown integration method type: '%s'. Supported types are 'NewtonCotes', 'GaussLegendre'", methodType)
	}

	if specificError != nil {
		return nil, fmt.Errorf("%w: %v", common.ErrInvalidStrategy, specificError)
	}

	// This safeguard checks if a strategy was unexpectedly not initialized.
	// This might happen if methodType is valid, order is valid, but no assignment occurs.
	if strategy == nil {
		// This case indicates a logic flaw in the factory itself (e.g., a new valid case added to documentation but not code).
		internalError := fmt.Errorf("failed to create strategy for method '%s', order %d due to an unexpected internal issue (strategy is nil after checks)", methodType, order)
		return nil, fmt.Errorf("%w: %v", common.ErrInvalidStrategy, internalError)
	}

	return strategy, nil
}
