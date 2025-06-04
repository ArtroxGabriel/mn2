package derivation

import (
	firstorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/first-order"
	fourthorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/fourth-order"
	secondorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/second-order"
	thirdorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/third-order"
)

type DerivationStrategy interface {
	CalculateFirst(
		fn func(float64) float64,
		x, h float64,
	) float64
	CalculateSecond(
		fn func(float64) float64,
		x, h float64,
	) float64
	CalculateThirty(
		fn func(float64) float64,
		x, h float64,
	) float64
}

var (
	_ DerivationStrategy = (*secondorder.ForwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*firstorder.ForwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*thirdorder.ForwardThirOrderStrategy)(nil)

	_ DerivationStrategy = (*secondorder.BackwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*firstorder.BackwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*thirdorder.BackwardThirOrderStrategy)(nil)

	_ DerivationStrategy = (*secondorder.CentralSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*fourthorder.CentralFourthOrderStrategy)(nil)
)
