package strategies

import (
	"fmt"
	"math"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration"
)

// Compile-time check to ensure NewtonCotesStrategy implements the IntegrationStrategy interface.
var _ integration.IntegrationStrategy = (*NewtonCotesStrategy)(nil)

// NewtonCotesStrategy implements numerical integration using Newton-Cotes formulas.
type NewtonCotesStrategy struct{}

// Calculate computes the definite integral of the function f(x) from a to b
// using Newton-Cotes formulas. The specific formula depends on n:
// n=1: Trapezoidal Rule (composite)
// n=2: Simpson's 1/3 Rule (composite)
// n=3: Simpson's 3/8 Rule (composite)
// Other values of n are not supported for closed Newton-Cotes common formulas.
// It uses an adaptive (composite) approach by increasing the number of subintervals
// until the desired tolerance is met.
func (s *NewtonCotesStrategy) Calculate(
	fn func(float64) float64,
	a, b float64,
	n int, // Degree for Newton-Cotes (1: Trapezoidal, 2: Simpson's 1/3, 3: Simpson's 3/8)
	tol float64,
) (float64, error) {
	if a == b {
		return 0.0, nil
	}
	if tol <= 0 {
		return 0, fmt.Errorf("tolerance must be positive, got %f", tol)
	}

	if n < 1 || n > 3 {
		return 0, fmt.Errorf("Newton-Cotes degree n must be 1 (Trapezoidal), 2 (Simpson's 1/3), or 3 (Simpson's 3/8), got %d", n)
	}

	// Adaptive (composite) quadrature
	var currentIntegral, previousIntegral float64
	// Initial number of subintervals. Must be compatible with the rule.
	// For Simpson's 1/3, it must be even. For Simpson's 3/8, multiple of 3.
	numSubIntervals := 1
	if n == 2 {
		numSubIntervals = 2 // Start with at least 2 for Simpson's 1/3
	} else if n == 3 {
		numSubIntervals = 3 // Start with at least 3 for Simpson's 3/8
	}


	// Initial calculation
	currentIntegral = s.calculateCompositeNewtonCotes(fn, a, b, n, numSubIntervals)

	maxIterations := 20 // Safeguard against infinite loops
	for iteration := 0; iteration < maxIterations; iteration++ {
		previousIntegral = currentIntegral

		// Increase number of subintervals
		// For n=1 (Trapezoidal), any doubling is fine.
		// For n=2 (Simpson's 1/3), numSubIntervals must be even.
		// For n=3 (Simpson's 3/8), numSubIntervals must be a multiple of 3.
		if n == 1 {
			numSubIntervals *= 2
		} else if n == 2 {
			numSubIntervals *= 2 // Already even, stays even
		} else { // n == 3
			// Try to double, then adjust to be a multiple of 3
			numSubIntervals *= 2
			if numSubIntervals % 3 != 0 {
				numSubIntervals = ((numSubIntervals / 3) + 1) * 3
			}
		}
		// Ensure a minimum reasonable number of intervals for higher order methods
		if n == 2 && numSubIntervals < 2 { numSubIntervals = 2 }
		if n == 3 && numSubIntervals < 3 { numSubIntervals = 3 }


		currentIntegral = s.calculateCompositeNewtonCotes(fn, a, b, n, numSubIntervals)

		// Check for convergence
		if math.Abs(previousIntegral) < 1e-9 { // Avoid division by zero or near-zero
			if math.Abs(currentIntegral-previousIntegral) < tol {
				return currentIntegral, nil
			}
		} else if math.Abs((currentIntegral-previousIntegral)/previousIntegral) < tol {
			return currentIntegral, nil
		}

		// Safety break for very large number of subintervals if tolerance is too small
		if numSubIntervals > 10000000 {
			break;
		}
	}

	return currentIntegral, fmt.Errorf("adaptive Newton-Cotes (n=%d) failed to converge within %d iterations for tolerance %g (last integral: %f, prev: %f, subintervals: %d)", n, maxIterations, tol, currentIntegral, previousIntegral, numSubIntervals)
}

// calculateCompositeNewtonCotes performs the composite Newton-Cotes sum.
func (s *NewtonCotesStrategy) calculateCompositeNewtonCotes(
	fn func(float64) float64,
	a, b float64,
	n_rule int, // 1 for Trapezoidal, 2 for Simpson's 1/3, 3 for Simpson's 3/8
	numSubIntervals int, // Total number of subintervals for the composite rule
) float64 {
	h := (b - a) / float64(numSubIntervals)
	sum := 0.0

	switch n_rule {
	case 1: // Composite Trapezoidal Rule
		sum = fn(a) + fn(b)
		for i := 1; i < numSubIntervals; i++ {
			sum += 2 * fn(a+float64(i)*h)
		}
		return (h / 2.0) * sum

	case 2: // Composite Simpson's 1/3 Rule
		// numSubIntervals must be even.
		if numSubIntervals%2 != 0 {
			// This should be handled by the calling adaptive logic, but as a safeguard:
			// Fallback or error, here we'll proceed but it might be inaccurate.
			// A better approach would be to ensure numSubIntervals is always appropriate.
			// For simplicity in this example, we assume it's correctly managed.
			// Or, one could adjust numSubIntervals here, e.g., numSubIntervals++
		}
		sum = fn(a) + fn(b)
		for i := 1; i < numSubIntervals; i++ {
			if i%2 == 1 { // Odd indices
				sum += 4 * fn(a+float64(i)*h)
			} else { // Even indices
				sum += 2 * fn(a+float64(i)*h)
			}
		}
		return (h / 3.0) * sum

	case 3: // Composite Simpson's 3/8 Rule
		// numSubIntervals must be a multiple of 3.
		if numSubIntervals%3 != 0 {
			// Safeguard similar to Simpson's 1/3.
		}
		sum = fn(a) + fn(b)
		for i := 1; i < numSubIntervals; i++ {
			if i%3 == 0 {
				sum += 2 * fn(a+float64(i)*h)
			} else {
				sum += 3 * fn(a+float64(i)*h)
			}
		}
		return (3.0 * h / 8.0) * sum
	}
	return 0 // Should not be reached if n_rule is validated
}
