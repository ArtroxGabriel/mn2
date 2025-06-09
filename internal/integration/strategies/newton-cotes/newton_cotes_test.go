package newtoncotes

import (
	"math"
	"strings"
	"testing"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	// "github.com/ArtroxGabriel/numeric-methods-cli/internal/integration" // Not strictly needed here
)

const tolerance = 1e-9 // General tolerance for comparisons

// Define some test functions
func f_x_squared(x float64) float64 { return x * x }      // Integral: x^3/3
func f_x_cubed(x float64) float64 { return x * x * x }    // Integral: x^4/4
func f_constant(x float64) float64 { return 5.0 }          // Integral: 5x
func f_sin_x(x float64) float64 { return math.Sin(x) }    // Integral: -Cos(x)
func f_exp_x(x float64) float64 { return math.Exp(x) }    // Integral: Exp(x)

// Helper to calculate analytical integral for polynomials like ax^p
func analyticalIntegralPoly(coeff, power, a, b float64) float64 {
	if power == -1 { // Special case for 1/x -> ln|x|
		if a*b <= 0 { // Integral is improper or undefined across 0
			panic("Analytical integral of 1/x across 0 is not handled here")
		}
		return coeff * (math.Log(math.Abs(b)) - math.Log(math.Abs(a)))
	}
	return coeff * (math.Pow(b, power+1) - math.Pow(a, power+1)) / (power + 1)
}

func analyticalIntegralSin(a, b float64) float64 {
	return -math.Cos(b) - (-math.Cos(a))
}

func analyticalIntegralExp(a, b float64) float64 {
    return math.Exp(b) - math.Exp(a)
}


func TestTrapezoidalRule(t *testing.T) {
	strategy := &TrapezoidalRule{}

	tests := []struct {
		name           string
		fn             common.MathFunc
		a, b           float64
		n              int
		expected       float64
		customTol      float64 // 0 means use default tolerance
		expectError    bool
		errorSubstring string
	}{
		{"x^2 from 0 to 1, n=10", f_x_squared, 0, 1, 10, analyticalIntegralPoly(1, 2, 0, 1), 1e-2, false, ""}, // Lower accuracy
		{"x^2 from 0 to 1, n=100", f_x_squared, 0, 1, 100, analyticalIntegralPoly(1, 2, 0, 1), 1e-4, false, ""},// Better with more n
		{"x^3 from 1 to 2, n=50", f_x_cubed, 1, 2, 50, analyticalIntegralPoly(1, 3, 1, 2), 1e-3, false, ""},
		{"constant from -1 to 1, n=2", f_constant, -1, 1, 2, analyticalIntegralPoly(5, 0, -1, 1), 0, false, ""}, // Exact
		{"sin(x) from 0 to pi, n=100", f_sin_x, 0, math.Pi, 100, analyticalIntegralSin(0, math.Pi), 1e-3, false, ""},
        {"exp(x) from 0 to 1, n=100", f_exp_x, 0, 1, 100, analyticalIntegralExp(0, 1), 1e-3, false, ""},
		{"a > b (x^2 from 1 to 0), n=10", f_x_squared, 1, 0, 10, -analyticalIntegralPoly(1, 2, 0, 1), 1e-2, false, ""},
		{"a == b (x^2 from 1 to 1), n=10", f_x_squared, 1, 1, 10, 0, 0, false, ""},
		{"invalid n=0", f_x_squared, 0, 1, 0, 0, 0, true, "n must be positive"},
		{"invalid n=-1", f_x_squared, 0, 1, -1, 0, 0, true, "n must be positive"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := strategy.Integrate(tt.fn, tt.a, tt.b, tt.n)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorSubstring)
				} else if tt.errorSubstring != "" && !strings.Contains(err.Error(), tt.errorSubstring) {
					t.Errorf("Expected error string '%s' not found in '%s'", tt.errorSubstring, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				currentTolerance := tolerance
				if tt.customTol > 0 {
					currentTolerance = tt.customTol
				}
				if math.Abs(result-tt.expected) > currentTolerance {
					t.Errorf("Expected %v, got %v, diff %v > tolerance %v", tt.expected, result, math.Abs(result-tt.expected), currentTolerance)
				}
			}
		})
	}
}

func TestSimpsonsOneThirdRule(t *testing.T) {
	strategy := &SimpsonsOneThirdRule{}
	tests := []struct {
		name        string
		fn          common.MathFunc
		a, b        float64
		n           int
		expected    float64
		customTol   float64
		expectError bool
		errorSubstring string
	}{
		{"x^2 from 0 to 1, n=10", f_x_squared, 0, 1, 10, analyticalIntegralPoly(1, 2, 0, 1), 0, false, ""}, // Exact for quadratics
		{"x^3 from 1 to 2, n=10", f_x_cubed, 1, 2, 10, analyticalIntegralPoly(1, 3, 1, 2), 0, false, ""},   // Exact for cubics
		{"sin(x) from 0 to pi, n=100", f_sin_x, 0, math.Pi, 100, analyticalIntegralSin(0, math.Pi), 1e-7, false, ""}, // Very accurate
        {"exp(x) from 0 to 1, n=100", f_exp_x, 0, 1, 100, analyticalIntegralExp(0, 1), 1e-7, false, ""},
		{"a > b (x^2 from 1 to 0), n=10", f_x_squared, 1, 0, 10, -analyticalIntegralPoly(1, 2, 0, 1), 0, false, ""},
		{"a == b (x^2 from 1 to 1), n=10", f_x_squared, 1, 1, 10, 0, 0, false, ""},
		{"invalid n=1 (odd)", f_x_squared, 0, 1, 1, 0, 0, true, "n must be a positive even integer"},
		{"invalid n=0", f_x_squared, 0, 1, 0, 0, 0, true, "n must be a positive even integer"},
		{"invalid n=3 (odd)", f_x_squared, 0, 1, 3, 0, 0, true, "n must be a positive even integer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := strategy.Integrate(tt.fn, tt.a, tt.b, tt.n)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorSubstring)
				} else if tt.errorSubstring != "" && !strings.Contains(err.Error(), tt.errorSubstring) {
					t.Errorf("Expected error string '%s' not found in '%s'", tt.errorSubstring, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
                currentTolerance := tolerance
				if tt.customTol > 0 {
					currentTolerance = tt.customTol
				}
				if math.Abs(result-tt.expected) > currentTolerance {
					t.Errorf("Expected %v, got %v, diff %v > tolerance %v", tt.expected, result, math.Abs(result-tt.expected), currentTolerance)
				}
			}
		})
	}
}

func TestSimpsonsThreeEighthRule(t *testing.T) {
	strategy := &SimpsonsThreeEighthRule{}
	tests := []struct {
		name        string
		fn          common.MathFunc
		a, b        float64
		n           int
		expected    float64
        customTol   float64
		expectError bool
		errorSubstring string
	}{
		{"x^2 from 0 to 1, n=9", f_x_squared, 0, 1, 9, analyticalIntegralPoly(1, 2, 0, 1), 0, false, ""}, // Exact for quadratics
		{"x^3 from 1 to 2, n=9", f_x_cubed, 1, 2, 9, analyticalIntegralPoly(1, 3, 1, 2), 0, false, ""},   // Exact for cubics
		{"sin(x) from 0 to pi, n=99", f_sin_x, 0, math.Pi, 99, analyticalIntegralSin(0, math.Pi), 1e-7, false, ""}, // Very accurate
        {"exp(x) from 0 to 1, n=99", f_exp_x, 0, 1, 99, analyticalIntegralExp(0, 1), 1e-7, false, ""},
		{"a > b (x^2 from 1 to 0), n=9", f_x_squared, 1, 0, 9, -analyticalIntegralPoly(1, 2, 0, 1), 0, false, ""},
		{"a == b (x^2 from 1 to 1), n=9", f_x_squared, 1, 1, 9, 0, 0, false, ""},
		{"invalid n=2 (not mult of 3)", f_x_squared, 0, 1, 2, 0, 0, true, "n must be a positive integer multiple of 3"},
		{"invalid n=0", f_x_squared, 0, 1, 0, 0, 0, true, "n must be a positive integer multiple of 3"},
		{"invalid n=4 (not mult of 3)", f_x_squared, 0, 1, 4, 0, 0, true, "n must be a positive integer multiple of 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := strategy.Integrate(tt.fn, tt.a, tt.b, tt.n)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorSubstring)
				} else if tt.errorSubstring != "" && !strings.Contains(err.Error(), tt.errorSubstring) {
					t.Errorf("Expected error string '%s' not found in '%s'", tt.errorSubstring, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
                currentTolerance := tolerance
				if tt.customTol > 0 {
					currentTolerance = tt.customTol
				}
				if math.Abs(result-tt.expected) > currentTolerance {
					t.Errorf("Expected %v, got %v, diff %v > tolerance %v", tt.expected, result, math.Abs(result-tt.expected), currentTolerance)
				}
			}
		})
	}
}

func TestBoolesRule(t *testing.T) {
	strategy := &BoolesRule{}
	tests := []struct {
		name        string
		fn          common.MathFunc
		a, b        float64
		n           int
		expected    float64
        customTol   float64
		expectError bool
		errorSubstring string
	}{
		// Boole's rule is exact for polynomials up to degree 5.
		{"x^2 from 0 to 1, n=4", f_x_squared, 0, 1, 4, analyticalIntegralPoly(1, 2, 0, 1), 0, false, ""},
		{"x^3 from 1 to 2, n=4", f_x_cubed, 1, 2, 4, analyticalIntegralPoly(1, 3, 1, 2), 0, false, ""},
		{"x^3 from 1 to 2, n=8 (composite)", f_x_cubed, 1, 2, 8, analyticalIntegralPoly(1, 3, 1, 2), 0, false, ""},
		{"sin(x) from 0 to pi, n=100", f_sin_x, 0, math.Pi, 100, analyticalIntegralSin(0, math.Pi), 1e-9, false, ""}, // Highly accurate
        {"exp(x) from 0 to 1, n=100", f_exp_x, 0, 1, 100, analyticalIntegralExp(0, 1), 1e-9, false, ""},
		{"a > b (x^2 from 1 to 0), n=4", f_x_squared, 1, 0, 4, -analyticalIntegralPoly(1, 2, 0, 1), 0, false, ""},
		{"a == b (x^2 from 1 to 1), n=4", f_x_squared, 1, 1, 4, 0, 0, false, ""},
		{"invalid n=2 (not mult of 4)", f_x_squared, 0, 1, 2, 0, 0, true, "n must be a positive integer multiple of 4"},
		{"invalid n=0", f_x_squared, 0, 1, 0, 0, 0, true, "n must be a positive integer multiple of 4"},
		{"invalid n=5 (not mult of 4)", f_x_squared, 0, 1, 5, 0, 0, true, "n must be a positive integer multiple of 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := strategy.Integrate(tt.fn, tt.a, tt.b, tt.n)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorSubstring)
				} else if tt.errorSubstring != "" && !strings.Contains(err.Error(), tt.errorSubstring) {
					t.Errorf("Expected error string '%s' not found in '%s'", tt.errorSubstring, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
                currentTolerance := tolerance
				if tt.customTol > 0 {
					currentTolerance = tt.customTol
				}
				if math.Abs(result-tt.expected) > currentTolerance {
					t.Errorf("Expected %v, got %v, diff %v > tolerance %v", tt.expected, result, math.Abs(result-tt.expected), currentTolerance)
				}
			}
		})
	}
}
