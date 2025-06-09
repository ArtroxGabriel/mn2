package gausslegendre

import (
	"math"
	"strings" // Required for strings.Contains (though not used in this version of tests)
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	// "github.com/ArtroxGabriel/numeric-methods-cli/internal/integration" // Not strictly needed
)

const defaultTolerance = 1e-9 // General tolerance if not specified
const exactTolerance = 1e-14  // Tolerance for cases where Gauss-Legendre is exact

// --- Test Functions ---
func f_constant(x float64) float64   { return 5.0 }           // Integral: 5x
func f_linear(x float64) float64     { return 2*x + 3 }       // Integral: x^2 + 3x
func f_x_squared(x float64) float64  { return x * x }           // Integral: x^3/3
func f_x_cubed(x float64) float64    { return x * x * x }       // Integral: x^4/4
func f_x_fourth(x float64) float64   { return math.Pow(x, 4) }  // Integral: x^5/5
func f_x_fifth(x float64) float64    { return math.Pow(x, 5) }  // Integral: x^6/6
func f_x_sixth(x float64) float64    { return math.Pow(x, 6) }  // Integral: x^7/7
func f_x_seventh(x float64) float64  { return math.Pow(x, 7) } // Integral: x^8/8
func f_sin_x(x float64) float64      { return math.Sin(x) }     // Integral: -Cos(x)
func f_exp_x(x float64) float64      { return math.Exp(x) }     // Integral: Exp(x)

// --- Analytical Integral Helpers ---
func integral_f_constant(a, b float64) float64 { // 5
    return 5 * (b - a)
}
func integral_f_linear(a, b float64) float64 { // 2x+3
	val_b := b*b + 3*b
	val_a := a*a + 3*a
	return val_b - val_a
}
func integral_f_x_squared(a, b float64) float64 { // x^2
	return (math.Pow(b, 3) - math.Pow(a, 3)) / 3.0
}
func integral_f_x_cubed(a, b float64) float64 { // x^3
	return (math.Pow(b, 4) - math.Pow(a, 4)) / 4.0
}
func integral_f_x_fourth(a, b float64) float64 { // x^4
    return (math.Pow(b, 5) - math.Pow(a, 5)) / 5.0
}
func integral_f_x_fifth(a, b float64) float64 { // x^5
	return (math.Pow(b, 6) - math.Pow(a, 6)) / 6.0
}
func integral_f_x_sixth(a, b float64) float64 { // x^6
    return (math.Pow(b, 7) - math.Pow(a, 7)) / 7.0
}
func integral_f_x_seventh(a, b float64) float64 { // x^7
	return (math.Pow(b, 8) - math.Pow(a, 8)) / 8.0
}
func integral_f_sin_x(a, b float64) float64 { // sin(x)
	return -math.Cos(b) - (-math.Cos(a))
}
func integral_f_exp_x(a, b float64) float64 { // exp(x)
	return math.Exp(b) - math.Exp(a)
}

type gaussLegendreTestCase struct {
	name     string
	fn       common.MathFunc
	a, b     float64
	n_param  int // Value passed to Integrate method, typically ignored by strategy
	expected float64
	tol      float64 // Specific tolerance for this test case
}

func runGaussLegendreTest(t *testing.T, strategy common.IntegrationStrategy, tt gaussLegendreTestCase) {
	t.Helper()
	result, err := strategy.Integrate(tt.fn, tt.a, tt.b, tt.n_param)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	currentTol := tt.tol
	if currentTol == 0 { // Use default if not set
		currentTol = defaultTolerance
	}

	if math.Abs(result-tt.expected) > currentTol {
		t.Errorf("Expected %.15f, got %.15f, diff %.15f > tolerance %.1e",
			tt.expected, result, math.Abs(result-tt.expected), currentTol)
	}
}

func TestGaussLegendreOrder1(t *testing.T) {
	strategy := &GaussLegendreOrder1{} // n=1 point, exact for degree 2n-1 = 1 polynomials
	tests := []gaussLegendreTestCase{
		{"constant [0,1] (exact)", f_constant, 0, 1, 1, integral_f_constant(0,1), exactTolerance},
		{"linear [0,1] (exact)", f_linear, 0, 1, 1, integral_f_linear(0,1), exactTolerance},
		{"linear [-1,1] (exact)", f_linear, -1, 1, 1, integral_f_linear(-1,1), exactTolerance},
		{"x^2 [0,1] (not exact)", f_x_squared, 0, 1, 1, integral_f_x_squared(0,1), 1e-1},
		{"sin(x) [0,pi]", f_sin_x, 0, math.Pi, 1, integral_f_sin_x(0,math.Pi), 5e-1}, // Poor for sin with 1 point
		{"a > b (linear from 1 to 0)", f_linear, 1, 0, 1, integral_f_linear(1,0), exactTolerance},
		{"a == b (linear from 1 to 1)", f_linear, 1, 1, 1, 0, exactTolerance},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { runGaussLegendreTest(t, strategy, tt) })
	}
}

func TestGaussLegendreOrder2(t *testing.T) {
	strategy := &GaussLegendreOrder2{} // n=2 points, exact for degree 2n-1 = 3 polynomials
	tests := []gaussLegendreTestCase{
		{"x^2 [-1,1] (exact)", f_x_squared, -1, 1, 2, integral_f_x_squared(-1,1), exactTolerance},
		{"x^3 [0,1] (exact)", f_x_cubed, 0, 1, 2, integral_f_x_cubed(0,1), exactTolerance},
		{"x^3 [-1,1] (exact)", f_x_cubed, -1, 1, 2, integral_f_x_cubed(-1,1), exactTolerance},
		{"x^fourth [0,1] (not exact)", f_x_fourth, 0, 1, 2, integral_f_x_fourth(0,1), 1e-2},
		{"sin(x) [0,pi]", f_sin_x, 0, math.Pi, 2, integral_f_sin_x(0,math.Pi), 1e-2},
		{"exp(x) [0,1]", f_exp_x, 0, 1, 2, integral_f_exp_x(0,1), 1e-3},
		{"a > b (x^3 from 2 to 1)", f_x_cubed, 2, 1, 2, integral_f_x_cubed(2,1), exactTolerance},
		{"a == b (x^3 from 1 to 1)", f_x_cubed, 1, 1, 2, 0, exactTolerance},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { runGaussLegendreTest(t, strategy, tt) })
	}
}

func TestGaussLegendreOrder3(t *testing.T) {
	strategy := &GaussLegendreOrder3{} // n=3 points, exact for degree 2n-1 = 5 polynomials
	tests := []gaussLegendreTestCase{
		{"x^fourth [-1,1] (exact)", f_x_fourth, -1, 1, 3, integral_f_x_fourth(-1,1), exactTolerance},
		{"x^fifth [0,1] (exact)", f_x_fifth, 0, 1, 3, integral_f_x_fifth(0,1), exactTolerance},
		{"x^fifth [-1,1] (exact)", f_x_fifth, -1, 1, 3, integral_f_x_fifth(-1,1), exactTolerance},
        {"x^sixth [0,1] (not exact)", f_x_sixth, 0, 1, 3, integral_f_x_sixth(0,1), 1e-3},
		{"sin(x) [0,pi]", f_sin_x, 0, math.Pi, 3, integral_f_sin_x(0,math.Pi), 1e-4},
		{"exp(x) [0,1]", f_exp_x, 0, 1, 3, integral_f_exp_x(0,1), 1e-5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { runGaussLegendreTest(t, strategy, tt) })
	}
}

func TestGaussLegendreOrder4(t *testing.T) {
	strategy := &GaussLegendreOrder4{} // n=4 points, exact for degree 2n-1 = 7 polynomials
	tests := []gaussLegendreTestCase{
		{"x^sixth [-1,1] (exact)", f_x_sixth, -1, 1, 4, integral_f_x_sixth(-1,1), exactTolerance},
		{"x^seventh [0,1] (exact)", f_x_seventh, 0, 1, 4, integral_f_x_seventh(0,1), exactTolerance},
		{"x^seventh [-1,1] (exact)", f_x_seventh, -1, 1, 4, integral_f_x_seventh(-1,1), exactTolerance},
		{"sin(x) [0,pi]", f_sin_x, 0, math.Pi, 4, integral_f_sin_x(0,math.Pi), 1e-6},
		{"exp(x) [0,1]", f_exp_x, 0, 1, 4, integral_f_exp_x(0,1), 1e-7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { runGaussLegendreTest(t, strategy, tt) })
	}
}

func TestGaussLegendre_GeneralCases(t *testing.T) {
    // Test generalGaussLegendre directly for error cases or unsupported orders if needed
    // For example, trying to get an unsupported order from generalGaussLegendre:
    _, err := generalGaussLegendre(f_linear, 0, 1, 0) // Order 0 is not supported
    if err == nil {
        t.Errorf("Expected error for unsupported order 0, got nil")
    } else {
        expectedErrorMsg := "Gauss-Legendre quadrature is not defined for order 0"
        if !strings.Contains(err.Error(), expectedErrorMsg) {
            t.Errorf("Expected error message containing '%s', got '%s'", expectedErrorMsg, err.Error())
        }
    }

    _, err = generalGaussLegendre(f_linear, 0, 1, 5) // Order 5 is not in our maps currently
    if err == nil {
        t.Errorf("Expected error for unsupported order 5, got nil")
    } else {
        expectedErrorMsg := "Gauss-Legendre quadrature is not defined for order 5"
        if !strings.Contains(err.Error(), expectedErrorMsg) {
            t.Errorf("Expected error message containing '%s', got '%s'", expectedErrorMsg, err.Error())
        }
    }
}
