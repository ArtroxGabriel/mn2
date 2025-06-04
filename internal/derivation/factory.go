package derivation

import (
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	firstorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/first-order"
	fourthorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/fourth-order"
	secondorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/second-order"
	thirdorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/third-order"
)

func DerivacaoFactory(strategy string, order uint) (DerivationStrategy, error) {
	var derivation DerivationStrategy
	switch {
	case strategy == "Forward" && order == 1:
		derivation = &firstorder.ForwardFirstOrderStrategy{}
	case strategy == "Forward" && order == 2:
		derivation = &secondorder.ForwardSecondOrderStrategy{}
	case strategy == "Forward" && order == 3:
		derivation = &thirdorder.ForwardThirOrderStrategy{}

	case strategy == "Backward" && order == 1:
		derivation = &firstorder.BackwardFirstOrderStrategy{}
	case strategy == "Backward" && order == 2:
		derivation = &secondorder.BackwardSecondOrderStrategy{}
	case strategy == "Backward" && order == 3:
		derivation = &thirdorder.BackwardThirOrderStrategy{}

	case strategy == "Central" && order == 2:
		derivation = &secondorder.CentralSecondOrderStrategy{}
	case strategy == "Central" && order == 4:
		derivation = &fourthorder.CentralFourthOrderStrategy{}
	default:
		return nil, common.ErrInvalidStrategy
	}

	return derivation, nil
}
