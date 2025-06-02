package derivation

import (
	"context"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

type Derivator struct {
	derivator DerivationStrategy
}

func NewDerivator(ctx context.Context, strategy string, order int) (*Derivator, error) {
	derivator := &Derivator{}

	deriv, err := DerivacaoFactory(ctx, strategy, order)
	if err != nil {
		return nil, err
	}
	derivator.derivator = deriv

	return derivator, nil
}

func (c *Derivator) Calculate(
	ctx context.Context,
	fn func(float64) float64,
	x, h float64,
	derivate int,
) error {
	switch derivate {
	case 1:
		c.derivator.CalculateFirst(ctx, fn, x, h)
	case 2:
		c.derivator.CalculateSecond(ctx, fn, x, h)
	case 3:
		c.derivator.CalculateThirty(ctx, fn, x, h)
	default:
		return common.ErrInvalidDerivate
	}

	return nil
}
