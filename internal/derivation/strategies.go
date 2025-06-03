package derivation

import (
	"context"

	firstorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/first-order"
	secondorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/second-order"
)

type DerivationStrategy interface {
	CalculateFirst(
		ctx context.Context,
		fn func(float64) float64,
		x, h float64,
	) float64
	CalculateSecond(
		ctx context.Context,
		fn func(float64) float64,
		x, h float64,
	) float64
	CalculateThirty(
		ctx context.Context,
		fn func(float64) float64,
		x, h float64,
	) float64
}

var (
	_ DerivationStrategy = (*secondorder.ForwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*firstorder.ForwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*secondorder.BackwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*firstorder.BackwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*secondorder.CentralSecondOrderStrategy)(nil)
)
