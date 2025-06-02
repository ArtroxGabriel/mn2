package functions

import (
	"math"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

var (
	PolynomialFunction common.MathFunc = func(x float64) float64 {
		return 2*x*x + 3*x + 1
	}

	ExponentialFunction common.MathFunc = func(x float64) float64 {
		return math.Exp(math.Pi*x) + 1
	}

	TrigonometricFunction common.MathFunc = func(x float64) float64 {
		return math.Sin(x) + math.Cos(x)
	}

	HyperbolicFunction common.MathFunc = func(x float64) float64 {
		return math.Sinh(x) + math.Cosh(x)
	}

	LogarithmicFunction common.MathFunc = func(x float64) float64 {
		if x <= 0 {
			panic(common.ErrZeroValue)
		}
		return math.Log(x) + math.Log10(x)
	}

	CompoundFunction common.MathFunc = func(x float64) float64 {
		return PolynomialFunction(x) + ExponentialFunction(x) + TrigonometricFunction(x) +
			HyperbolicFunction(x)
	}
)
