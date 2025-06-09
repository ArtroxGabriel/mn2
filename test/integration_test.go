package test

import (
	"math"
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration"
)

const tolerance = 1e-6 // Tolerance for comparing float values

// Simple polynomial function for testing: f(x) = x^2
// Integral from a to b is (b^3/3) - (a^3/3)
func polyFunc(x float64) float64 {
	return x * x
}
func polyFuncIntegral(a, b float64) float64 {
	return (math.Pow(b, 3) / 3.0) - (math.Pow(a, 3) / 3.0)
}

// More complex function: f(x) = sin(x)
// Integral from a to b is -cos(b) - (-cos(a)) = cos(a) - cos(b)
func sinFunc(x float64) float64 {
	return math.Sin(x)
}
func sinFuncIntegral(a, b float64) float64 {
	return math.Cos(a) - math.Cos(b)
}

// Function f(x) = e^x
// Integral from a to b is e^b - e^a
func expFunc(x float64) float64 {
	return math.Exp(x)
}
func expFuncIntegral(a, b float64) float64 {
	return math.Exp(b) - math.Exp(a)
}


type integrationTestCase struct {
	name         string
	strategyName string
	fn           func(float64) float64
	a, b         float64
	n            int // For Newton-Cotes, number of subintervals
	expected     float64
	customTol    float64 // Custom tolerance if needed, 0 for default
	expectError  bool
}

func runIntegrationTest(t *testing.T, tc integrationTestCase) {
	t.Helper()
	integrator, err := integration.NewIntegrator(tc.strategyName)
	if err != nil {
		if tc.expectError {
			return // Expected error during integrator creation
		}
		t.Fatalf("NewIntegrator(%s) failed: %v", tc.strategyName, err)
	}

	result, err := integrator.Calculate(tc.fn, tc.a, tc.b, tc.n)

	if tc.expectError {
		if err == nil {
			t.Errorf("Test %s: Expected an error, but got nil", tc.name)
		}
		return // Expected error, no need to check result
	}
	if err != nil {
		t.Fatalf("Test %s: Calculate failed: %v", tc.name, err)
	}

	tol := tolerance
	if tc.customTol != 0 {
		tol = tc.customTol
	}

	if math.Abs(result-tc.expected) > tol {
		t.Errorf("Test %s: Expected %v, got %v. Difference: %v", tc.name, tc.expected, result, math.Abs(result-tc.expected))
	}
}

func TestNewtonCotesIntegration(t *testing.T) {
	testCases := []integrationTestCase{
		// NewtonCotesOrder1 (Trapezoidal)
		{name: "NC1_Poly_0_1_N10", strategyName: "NewtonCotesOrder1", fn: polyFunc, a: 0, b: 1, n: 10, expected: polyFuncIntegral(0, 1), customTol: 2e-3},
		{name: "NC1_Poly_0_1_N100", strategyName: "NewtonCotesOrder1", fn: polyFunc, a: 0, b: 1, n: 100, expected: polyFuncIntegral(0, 1), customTol: 2e-5},
		{name: "NC1_Sin_0_Pi_N20", strategyName: "NewtonCotesOrder1", fn: sinFunc, a: 0, b: math.Pi, n: 20, expected: sinFuncIntegral(0, math.Pi), customTol: 5e-3},
		{name: "NC1_Exp_0_1_N50", strategyName: "NewtonCotesOrder1", fn: expFunc, a: 0, b: 1, n: 50, expected: expFuncIntegral(0,1), customTol: 6e-5},
		{name: "NC1_Poly_1_0_N10_ReverseBounds", strategyName: "NewtonCotesOrder1", fn: polyFunc, a: 1, b: 0, n: 10, expected: polyFuncIntegral(0, 1), customTol: 2e-3}, // Integral from 0 to 1
		{name: "NC1_ZeroInterval", strategyName: "NewtonCotesOrder1", fn: polyFunc, a: 1, b: 1, n: 10, expected: 0},
		{name: "NC1_Error_N0", strategyName: "NewtonCotesOrder1", fn: polyFunc, a: 0, b: 1, n: 0, expectError: true},


		// NewtonCotesOrder2 (Simpson's 1/3)
		{name: "NC2_Poly_0_1_N10", strategyName: "NewtonCotesOrder2", fn: polyFunc, a: 0, b: 1, n: 10, expected: polyFuncIntegral(0, 1)},
		{name: "NC2_Poly_0_1_N100", strategyName: "NewtonCotesOrder2", fn: polyFunc, a: 0, b: 1, n: 100, expected: polyFuncIntegral(0, 1)},
		{name: "NC2_Sin_0_Pi_N20", strategyName: "NewtonCotesOrder2", fn: sinFunc, a: 0, b: math.Pi, n: 20, expected: sinFuncIntegral(0, math.Pi), customTol: 7e-6},
		{name: "NC2_Exp_0_1_N50", strategyName: "NewtonCotesOrder2", fn: expFunc, a: 0, b: 1, n: 50, expected: expFuncIntegral(0,1), customTol: 2e-9},
		{name: "NC2_Error_N_Odd", strategyName: "NewtonCotesOrder2", fn: polyFunc, a: 0, b: 1, n: 9, expectError: true},

		// NewtonCotesOrder3 (Simpson's 3/8)
		{name: "NC3_Poly_0_1_N12", strategyName: "NewtonCotesOrder3", fn: polyFunc, a: 0, b: 1, n: 12, expected: polyFuncIntegral(0, 1)},
		{name: "NC3_Poly_0_1_N99", strategyName: "NewtonCotesOrder3", fn: polyFunc, a: 0, b: 1, n: 99, expected: polyFuncIntegral(0, 1)},
		{name: "NC3_Sin_0_Pi_N21", strategyName: "NewtonCotesOrder3", fn: sinFunc, a: 0, b: math.Pi, n: 21, expected: sinFuncIntegral(0, math.Pi), customTol: 1.3e-5},
		{name: "NC3_Exp_0_1_N51", strategyName: "NewtonCotesOrder3", fn: expFunc, a: 0, b: 1, n: 51, expected: expFuncIntegral(0,1), customTol: 4e-9},
		{name: "NC3_Error_N_NotMult3", strategyName: "NewtonCotesOrder3", fn: polyFunc, a: 0, b: 1, n: 10, expectError: true},

		// NewtonCotesOrder4 (Boole's)
		{name: "NC4_Poly_0_1_N12", strategyName: "NewtonCotesOrder4", fn: polyFunc, a: 0, b: 1, n: 12, expected: polyFuncIntegral(0, 1), customTol: 1e-9},
		{name: "NC4_Poly_0_1_N100", strategyName: "NewtonCotesOrder4", fn: polyFunc, a: 0, b: 1, n: 100, expected: polyFuncIntegral(0, 1), customTol: 1e-10},
		{name: "NC4_Sin_0_Pi_N20", strategyName: "NewtonCotesOrder4", fn: sinFunc, a: 0, b: math.Pi, n: 20, expected: sinFuncIntegral(0, math.Pi), customTol: 7e-8},
		{name: "NC4_Exp_0_1_N52", strategyName: "NewtonCotesOrder4", fn: expFunc, a: 0, b: 1, n: 52, expected: expFuncIntegral(0,1), customTol: 1e-11},
		{name: "NC4_Error_N_NotMult4", strategyName: "NewtonCotesOrder4", fn: polyFunc, a: 0, b: 1, n: 10, expectError: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runIntegrationTest(t, tc)
		})
	}
}

func TestGaussLegendreIntegration(t *testing.T) {
	// For Gauss-Legendre, 'n' in Calculate is not strictly used by the basic strategy,
	// as the number of points is fixed by the strategy's order.
	// We pass n=1 (or any positive int) as a placeholder.
	testCases := []integrationTestCase{
		// GaussLegendreOrder1
		{name: "GL1_Poly_0_1", strategyName: "GaussLegendreOrder1", fn: polyFunc, a: 0, b: 1, n: 1, expected: polyFuncIntegral(0, 1), customTol: 1e-1}, // Low accuracy expected
		{name: "GL1_Sin_0_Pi", strategyName: "GaussLegendreOrder1", fn: sinFunc, a: 0, b: math.Pi, n: 1, expected: sinFuncIntegral(0, math.Pi), customTol: 1.2},
        {name: "GL1_Exp_0_1", strategyName: "GaussLegendreOrder1", fn: expFunc, a:0, b:1, n:1, expected: expFuncIntegral(0,1), customTol: 7e-2},

		// GaussLegendreOrder2
		{name: "GL2_Poly_0_1", strategyName: "GaussLegendreOrder2", fn: polyFunc, a: 0, b: 1, n: 1, expected: polyFuncIntegral(0, 1)}, // Exact for cubics
		{name: "GL2_Sin_0_Pi", strategyName: "GaussLegendreOrder2", fn: sinFunc, a: 0, b: math.Pi, n: 1, expected: sinFuncIntegral(0, math.Pi), customTol: 6.5e-2},
        {name: "GL2_Exp_0_1", strategyName: "GaussLegendreOrder2", fn: expFunc, a:0, b:1, n:1, expected: expFuncIntegral(0,1), customTol: 4e-4},

		// GaussLegendreOrder3
		{name: "GL3_Poly_0_1", strategyName: "GaussLegendreOrder3", fn: polyFunc, a: 0, b: 1, n: 1, expected: polyFuncIntegral(0, 1)}, // Exact for quintics
		{name: "GL3_Sin_0_Pi", strategyName: "GaussLegendreOrder3", fn: sinFunc, a: 0, b: math.Pi, n: 1, expected: sinFuncIntegral(0, math.Pi), customTol: 1.4e-3},
        {name: "GL3_Exp_0_1", strategyName: "GaussLegendreOrder3", fn: expFunc, a:0, b:1, n:1, expected: expFuncIntegral(0,1), customTol: 1e-6},

		// GaussLegendreOrder4
		{name: "GL4_Poly_0_1", strategyName: "GaussLegendreOrder4", fn: polyFunc, a: 0, b: 1, n: 1, expected: polyFuncIntegral(0, 1)}, // Exact for 7th degree polynomials
		{name: "GL4_Sin_0_Pi", strategyName: "GaussLegendreOrder4", fn: sinFunc, a: 0, b: math.Pi, n: 1, expected: sinFuncIntegral(0, math.Pi), customTol: 1.6e-5},
        {name: "GL4_Exp_0_1", strategyName: "GaussLegendreOrder4", fn: expFunc, a:0, b:1, n:1, expected: expFuncIntegral(0,1), customTol: 1e-9},
		{name: "GL4_Poly_ZeroInterval", strategyName: "GaussLegendreOrder4", fn: polyFunc, a: 1, b: 1, n: 1, expected: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runIntegrationTest(t, tc)
		})
	}
}
