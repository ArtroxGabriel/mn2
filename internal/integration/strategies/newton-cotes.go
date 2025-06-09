package strategies

import (
	"fmt"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integrationcore"
)

// NewtonCotesOrder1 implements the Trapezoidal rule for numerical integration.
type NewtonCotesOrder1 struct{}

// Calculate computes the definite integral using the Trapezoidal rule.
// Formula: ∫[a,b] f(x) dx ≈ (h/2) * [f(a) + f(b) + 2 * Σ[i=1 to n-1] f(a+ih)]
// where h = (b-a)/n
func (s *NewtonCotesOrder1) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number of subintervals n must be positive")
	}
	if a == b {
		return 0, nil
	}
	if a > b { // Ensure a < b for correct calculation
		a, b = b, a
	}

	h := (b - a) / float64(n)
	sum := fn(a) + fn(b)

	for i := 1; i < n; i++ {
		sum += 2 * fn(a+float64(i)*h)
	}

	return (h / 2) * sum, nil
}

// NewtonCotesOrder2 implements Simpson's 1/3 rule for numerical integration.
type NewtonCotesOrder2 struct{}

// Calculate computes the definite integral using Simpson's 1/3 rule.
// Formula: ∫[a,b] f(x) dx ≈ (h/3) * [f(a) + f(b) + 4 * Σ[i=1,3,5...n-1] f(a+ih) + 2 * Σ[i=2,4,6...n-2] f(a+ih)]
// where h = (b-a)/n and n must be an even number.
func (s *NewtonCotesOrder2) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number of subintervals n must be positive")
	}
	if n%2 != 0 {
		return 0, fmt.Errorf("number of subintervals n must be even for Simpson's 1/3 rule")
	}
	if a == b {
		return 0, nil
	}
    if a > b { // Ensure a < b for correct calculation
		a, b = b, a
	}

	h := (b - a) / float64(n)
	sum := fn(a) + fn(b)

	for i := 1; i < n; i++ {
		if i%2 == 1 { // Odd indices
			sum += 4 * fn(a+float64(i)*h)
		} else { // Even indices
			sum += 2 * fn(a+float64(i)*h)
		}
	}

	return (h / 3) * sum, nil
}

// NewtonCotesOrder3 implements Simpson's 3/8 rule for numerical integration.
type NewtonCotesOrder3 struct{}

// Calculate computes the definite integral using Simpson's 3/8 rule.
// Formula: ∫[a,b] f(x) dx ≈ (3h/8) * [f(a) + f(b) + 3 * Σ[i=1,2,4,5...n-1] f(a+ih) + 2 * Σ[i=3,6,9...n-3] f(a+ih)]
// where h = (b-a)/n and n must be a multiple of 3.
func (s *NewtonCotesOrder3) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number of subintervals n must be positive")
	}
	if n%3 != 0 {
		return 0, fmt.Errorf("number of subintervals n must be a multiple of 3 for Simpson's 3/8 rule")
	}
	if a == b {
		return 0, nil
	}
    if a > b { // Ensure a < b for correct calculation
		a, b = b, a
	}

	h := (b - a) / float64(n)
	sum := fn(a) + fn(b)

	for i := 1; i < n; i++ {
		if i%3 == 0 { // Multiples of 3
			sum += 2 * fn(a+float64(i)*h)
		} else { // Other indices
			sum += 3 * fn(a+float64(i)*h)
		}
	}

	return (3 * h / 8) * sum, nil
}

// NewtonCotesOrder4 implements Boole's rule for numerical integration.
type NewtonCotesOrder4 struct{}

// Calculate computes the definite integral using Boole's rule.
// Formula: ∫[a,b] f(x) dx ≈ (2h/45) * [7f(a) + 32f(a+h) + 12f(a+2h) + 32f(a+3h) + 7f(b)] for a single interval (n=4)
// For composite Boole's rule, n must be a multiple of 4.
// Composite: (2h/45) * Σ (from k=0 to n/4 - 1) [7f(x_4k) + 32f(x_4k+1) + 12f(x_4k+2) + 32f(x_4k+3) + 7f(x_4k+4)]
func (s *NewtonCotesOrder4) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number of subintervals n must be positive")
	}
	if n%4 != 0 {
		return 0, fmt.Errorf("number of subintervals n must be a multiple of 4 for Boole's rule")
	}
	if a == b {
		return 0, nil
	}
    if a > b { // Ensure a < b for correct calculation
		a, b = b, a
	}

	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i += 4 {
		x0 := a + float64(i)*h
		x1 := a + float64(i+1)*h
		x2 := a + float64(i+2)*h
		x3 := a + float64(i+3)*h
		x4 := a + float64(i+4)*h
		sum += 7*fn(x0) + 32*fn(x1) + 12*fn(x2) + 32*fn(x3) + 7*fn(x4)
	}

	return (2 * h / 45) * sum, nil
}

// Compile-time interface compliance checks.
var (
	_ integrationcore.IntegrationStrategy = (*NewtonCotesOrder1)(nil)
	_ integrationcore.IntegrationStrategy = (*NewtonCotesOrder2)(nil)
	_ integrationcore.IntegrationStrategy = (*NewtonCotesOrder3)(nil)
	_ integrationcore.IntegrationStrategy = (*NewtonCotesOrder4)(nil)
)
