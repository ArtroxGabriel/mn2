package functions

import (
	"math"
	"strings"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

var PredefinedFunctions = []common.FunctionDefinition{
	{ID: "polynomial", Name: "P(x) = 2x^2 +3x + 1", Func: polynomialFunction},
	{ID: "exponential", Name: "exp^x = e^(πx) + 1", Func: exponentialFunction},
	{ID: "trigonometric", Name: "T(x) = sin(x) + cos(x)", Func: trigonometricFunction},
	{ID: "hyperbolic", Name: "H(x) = sinh(x) + cosh(x)", Func: hyperbolicFunction},
	{ID: "logarithmic", Name: "L(x) = ln(x) + log10(x)", Func: logarithmicFunction},
	{ID: "compound", Name: "C(x) = P(x) + exp^x + T(x) + H(x)", Func: compoundFunction},
}

var (
	polynomialFunction common.MathFunc = func(x float64) float64 {
		return 2*x*x + 3*x + 1
	}

	exponentialFunction common.MathFunc = func(x float64) float64 {
		return math.Exp(math.Pi*x) + 1
	}

	trigonometricFunction common.MathFunc = func(x float64) float64 {
		return math.Sin(x) + math.Cos(x)
	}

	hyperbolicFunction common.MathFunc = func(x float64) float64 {
		return math.Sinh(x) + math.Cosh(x)
	}

	logarithmicFunction common.MathFunc = func(x float64) float64 {
		if x <= 0 {
			panic(common.ErrZeroValue)
		}
		return math.Log(x) + math.Log10(x)
	}

	compoundFunction common.MathFunc = func(x float64) float64 {
		return polynomialFunction(x) + exponentialFunction(x) + trigonometricFunction(x) +
			hyperbolicFunction(x)
	}
)

func GetPredefinedFunc(id string) common.MathFunc {
	for _, def := range PredefinedFunctions {
		if def.ID == id {
			return def.Func
		}
	}
	return nil // Ou uma função padrão/erro
}

func GetFuncName(id string) string {
	if strings.TrimSpace(id) == "" {
		return "Nenhuma"
	}

	for _, def := range PredefinedFunctions {
		if def.ID == id {
			return def.Name
		}
	}

	return "Função Personalizada"
}

func GetFunctionDefinitions() []common.FunctionDefinition {
	return PredefinedFunctions
}
