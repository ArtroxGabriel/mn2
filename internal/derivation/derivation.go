package derivation

import (
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

type Derivator struct {
	derivator DerivationStrategy
}

func NewDerivator(strategy string, order uint) (*Derivator, error) {
	derivator := &Derivator{}

	deriv, err := DerivacaoFactory(strategy, order)
	if err != nil {
		return nil, err
	}
	derivator.derivator = deriv

	return derivator, nil
}

func (c *Derivator) Calculate(
	fn func(float64) float64,
	x, h float64,
	derivate int,
) (float64, error) {
	var result float64
	switch derivate {
	case 1:
		result = c.derivator.CalculateFirst(fn, x, h)
	case 2:
		result = c.derivator.CalculateSecond(fn, x, h)
	case 3:
		result = c.derivator.CalculateThirty(fn, x, h)
	default:
		return 0, common.ErrInvalidDerivate
	}

	return result, nil
}
