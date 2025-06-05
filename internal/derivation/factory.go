// Package derivation provides finite difference numerical differentiation strategies.
// This file defines the factory function for creating appropriate derivation strategies
// based on the specified method and accuracy order.
package derivation

import (
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	firstorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/first-order"
	fourthorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/fourth-order"
	secondorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/second-order"
	thirdorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/third-order"
)

// DerivacaoFactory creates and returns the appropriate numerical differentiation strategy
// based on the specified method and accuracy order.
//
// This factory function implements the Factory design pattern to create different
// finite difference strategies:
//
// Forward Difference Methods:
//   - First-order forward difference (O(h) accuracy)
//   - Second-order forward difference (O(h²) accuracy)
//   - Third-order forward difference (O(h³) accuracy)
//
// Backward Difference Methods:
//   - First-order backward difference (O(h) accuracy)
//   - Second-order backward difference (O(h²) accuracy)
//   - Third-order backward difference (O(h³) accuracy)
//
// Central Difference Methods:
//   - Second-order central difference (O(h²) accuracy)
//   - Fourth-order central difference (O(h⁴) accuracy)
//
// Parameters:
//   - strategy: The finite difference method ("Forward", "Backward", or "Central")
//   - order: The order of accuracy for the method
//
// Returns:
//   - DerivationStrategy: The configured strategy implementation
//   - error: An error if the strategy/order combination is not supported
//
// Example:
//
//	strategy, err := DerivacaoFactory("Central", 2)
//	if err != nil {
//	    log.Fatal(err)
//	}
func DerivacaoFactory(strategy string, order uint) (DerivationStrategy, error) {
	var derivation DerivationStrategy
	switch {
	// Forward difference strategies
	case strategy == "Forward" && order == 1:
		derivation = &firstorder.ForwardFirstOrderStrategy{}
	case strategy == "Forward" && order == 2:
		derivation = &secondorder.ForwardSecondOrderStrategy{}
	case strategy == "Forward" && order == 3:
		derivation = &thirdorder.ForwardThirOrderStrategy{}

	// Backward difference strategies
	case strategy == "Backward" && order == 1:
		derivation = &firstorder.BackwardFirstOrderStrategy{}
	case strategy == "Backward" && order == 2:
		derivation = &secondorder.BackwardSecondOrderStrategy{}
	case strategy == "Backward" && order == 3:
		derivation = &thirdorder.BackwardThirOrderStrategy{}

	// Central difference strategies
	case strategy == "Central" && order == 2:
		derivation = &secondorder.CentralSecondOrderStrategy{}
	case strategy == "Central" && order == 4:
		derivation = &fourthorder.CentralFourthOrderStrategy{}
	default:
		return nil, common.ErrInvalidStrategy
	}

	return derivation, nil
}
