package newtoncotes

import (
	"fmt"
	"math"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration"
)

// TrapezoidalRule implements Newton-Cotes integration of order 1.
type TrapezoidalRule struct{}

// Integrate computes the definite integral using the Trapezoidal Rule.
// n: number of subintervals. Must be >= 1.
func (s *TrapezoidalRule) Integrate(fn common.MathFunc, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number of subintervals n must be positive, got %d", n)
	}
	if a == b {
		return 0, nil
	}
    if b < a {
        // Standard practice: integral from b to a is - (integral from a to b)
        res, err := s.Integrate(fn, b, a, n)
        return -res, err
    }

	h := (b - a) / float64(n)
	sum := fn(a) + fn(b)

	for i := 1; i < n; i++ {
		sum += 2 * fn(a+float64(i)*h)
	}

	return (h / 2) * sum, nil
}

// SimpsonsOneThirdRule implements Newton-Cotes integration of order 2.
type SimpsonsOneThirdRule struct{}

// Integrate computes the definite integral using Simpson's 1/3 Rule.
// n: number of subintervals. Must be an even number >= 2.
func (s *SimpsonsOneThirdRule) Integrate(fn common.MathFunc, a, b float64, n int) (float64, error) {
	if n <= 0 || n%2 != 0 {
		return 0, fmt.Errorf("number of subintervals n must be a positive even integer, got %d", n)
	}
	if a == b {
		return 0, nil
	}
    if b < a {
        res, err := s.Integrate(fn, b, a, n)
        return -res, err
    }

	h := (b - a) / float64(n)
	sum := fn(a) + fn(b)

	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sum += 2 * fn(a+float64(i)*h)
		} else {
			sum += 4 * fn(a+float64(i)*h)
		}
	}

	return (h / 3) * sum, nil
}

// SimpsonsThreeEighthRule implements Newton-Cotes integration of order 3.
type SimpsonsThreeEighthRule struct{}

// Integrate computes the definite integral using Simpson's 3/8 Rule.
// n: number of subintervals. Must be a multiple of 3, >= 3.
func (s *SimpsonsThreeEighthRule) Integrate(fn common.MathFunc, a, b float64, n int) (float64, error) {
	if n <= 0 || n%3 != 0 {
		return 0, fmt.Errorf("number of subintervals n must be a positive integer multiple of 3, got %d", n)
	}
	if a == b {
		return 0, nil
	}
    if b < a {
        res, err := s.Integrate(fn, b, a, n)
        return -res, err
    }

	h := (b - a) / float64(n)
	sum := fn(a) + fn(b)

	for i := 1; i < n; i++ {
		if i%3 == 0 {
			sum += 2 * fn(a+float64(i)*h)
		} else {
			sum += 3 * fn(a+float64(i)*h)
		}
	}

	return (3 * h / 8) * sum, nil
}

// BoolesRule implements Newton-Cotes integration of order 4.
type BoolesRule struct{}

// Integrate computes the definite integral using Boole's Rule.
// n: number of subintervals. Must be a multiple of 4, >= 4.
// This implementation uses the composite Boole's rule.
func (s *BoolesRule) Integrate(fn common.MathFunc, a, b float64, n int) (float64, error) {
	if n <= 0 || n%4 != 0 {
		return 0, fmt.Errorf("number of subintervals n must be a positive integer multiple of 4, got %d", n)
	}
	if a == b {
		return 0, nil
	}
    if b < a {
        res, err := s.Integrate(fn, b, a, n)
        return -res, err
    }

	h := (b - a) / float64(n)
	sum := 0.0

    // Apply Boole's rule over each set of 4 subintervals
    for i := 0; i < n; i += 4 {
        x0 := a + float64(i)*h
        x1 := a + float64(i+1)*h
        x2 := a + float64(i+2)*h
        x3 := a + float64(i+3)*h
        x4 := a + float64(i+4)*h
        sum += (2 * h / 45) * (7*fn(x0) + 32*fn(x1) + 12*fn(x2) + 32*fn(x3) + 7*fn(x4))
    }
	return sum, nil
}

// Ensure strategies implement the interface
var _ integration.IntegrationStrategy = (*TrapezoidalRule)(nil)
var _ integration.IntegrationStrategy = (*SimpsonsOneThirdRule)(nil)
var _ integration.IntegrationStrategy = (*SimpsonsThreeEighthRule)(nil)
var _ integration.IntegrationStrategy = (*BoolesRule)(nil)
