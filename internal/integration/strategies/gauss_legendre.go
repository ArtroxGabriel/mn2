package strategies

import (
	"fmt"
	"math"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration"
)

// Compile-time check to ensure GaussLegendreStrategy implements the IntegrationStrategy interface.
var _ integration.IntegrationStrategy = (*GaussLegendreStrategy)(nil)

// GaussLegendreStrategy implements numerical integration using Gauss-Legendre quadrature.
type GaussLegendreStrategy struct{}

// legendrePolynomialRootsAndWeights returns the roots and weights for Gauss-Legendre quadrature
// for a given degree n.
// Note: This is a simplified placeholder. Real implementation requires a robust root-finding algorithm
// for Legendre polynomials (e.g., Newton-Raphson or using precomputed values for common N).
func legendrePolynomialRootsAndWeights(n int) ([]float64, []float64, error) {
	if n <= 0 {
		return nil, nil, fmt.Errorf("degree n must be positive for Gauss-Legendre, got %d", n)
	}
	// Placeholder values for common degrees.
	// For a production system, these would be calculated or looked up with higher precision.
	switch n {
	case 1: // Corresponds to 1 point, exact for polynomials of degree up to 2*1-1 = 1
		return []float64{0.0}, []float64{2.0}, nil
	case 2: // Corresponds to 2 points, exact for polynomials of degree up to 2*2-1 = 3
		return []float64{-1.0 / math.Sqrt(3.0), 1.0 / math.Sqrt(3.0)}, []float64{1.0, 1.0}, nil
	case 3: // Corresponds to 3 points, exact for polynomials of degree up to 2*3-1 = 5
		return []float64{-math.Sqrt(3.0 / 5.0), 0.0, math.Sqrt(3.0 / 5.0)}, []float64{5.0 / 9.0, 8.0 / 9.0, 5.0 / 9.0}, nil
	case 4:
		roots := []float64{
			-math.Sqrt((3.0/7.0) + (2.0/7.0)*math.Sqrt(6.0/5.0)),
			-math.Sqrt((3.0/7.0) - (2.0/7.0)*math.Sqrt(6.0/5.0)),
			math.Sqrt((3.0/7.0) - (2.0/7.0)*math.Sqrt(6.0/5.0)),
			math.Sqrt((3.0/7.0) + (2.0/7.0)*math.Sqrt(6.0/5.0)),
		}
		weights := []float64{
			(18.0 - math.Sqrt(30.0)) / 36.0,
			(18.0 + math.Sqrt(30.0)) / 36.0,
			(18.0 + math.Sqrt(30.0)) / 36.0,
			(18.0 - math.Sqrt(30.0)) / 36.0,
		}
		return roots, weights, nil
	case 5:
		roots := []float64{
			-math.Sqrt(5.0+2.0*math.Sqrt(10.0/7.0)) / 3.0,
			-math.Sqrt(5.0-2.0*math.Sqrt(10.0/7.0)) / 3.0,
			0.0,
			math.Sqrt(5.0-2.0*math.Sqrt(10.0/7.0)) / 3.0,
			math.Sqrt(5.0+2.0*math.Sqrt(10.0/7.0)) / 3.0,
		}
		weights := []float64{
			(322.0 - 13.0*math.Sqrt(70.0)) / 900.0,
			(322.0 + 13.0*math.Sqrt(70.0)) / 900.0,
			128.0 / 225.0,
			(322.0 + 13.0*math.Sqrt(70.0)) / 900.0,
			(322.0 - 13.0*math.Sqrt(70.0)) / 900.0,
		}
		// Correcting order for standard Gauss-Legendre from -1 to 1
		return []float64{roots[2], roots[1], roots[3], roots[0], roots[4]}, []float64{weights[2], weights[1], weights[3], weights[0], weights[4]}, nil


	default:
		// For n > 5, precomputed values or a numerical root finder would be necessary.
		// This is a simplification for this example.
		return nil, nil, fmt.Errorf("Gauss-Legendre for n=%d not implemented in this placeholder (supports 1-5)", n)
	}
}

// Calculate computes the definite integral of the function f(x) from a to b
// using n-point Gauss-Legendre quadrature.
// It uses an adaptive approach by doubling the number of intervals until the desired tolerance is met.
func (s *GaussLegendreStrategy) Calculate(
	fn func(float64) float64,
	a, b float64,
	n int, // Degree of Legendre polynomial, corresponds to n points
	tol float64,
) (float64, error) {
	if a == b {
		return 0.0, nil
	}
	if n <= 0 {
		return 0, fmt.Errorf("number of points n must be positive for Gauss-Legendre, got %d", n)
	}
	if tol <= 0 {
		return 0, fmt.Errorf("tolerance must be positive, got %f", tol)
	}

	// Get Legendre roots (xi) and weights (wi) for the interval [-1, 1]
	roots, weights, err := legendrePolynomialRootsAndWeights(n)
	if err != nil {
		return 0, fmt.Errorf("failed to get Legendre roots/weights for n=%d: %w", n, err)
	}

	// Transformation factors for the interval [a, b]
	// t = ( (b-a)/2 )*x + ( (a+b)/2 )
	// dt = ( (b-a)/2 )*dx
	factor1 := (b - a) / 2.0
	factor2 := (a + b) / 2.0

	// Adaptive quadrature
	var currentIntegral, previousIntegral float64
	numIntervals := 1 // Start with 1 interval covering [a,b]

	// Initial calculation with numIntervals
	currentIntegral, err = s.calculateGaussLegendreForIntervals(fn, a, b, n, roots, weights, numIntervals, factor1, factor2)
	if err != nil {
		return 0, err // Should not happen if roots/weights are valid
	}

	maxIterations := 20 // Safeguard against infinite loops
	for iteration := 0; iteration < maxIterations; iteration++ {
		previousIntegral = currentIntegral
		numIntervals *= 2 // Double the number of subintervals

		currentIntegral, err = s.calculateGaussLegendreForIntervals(fn, a, b, n, roots, weights, numIntervals, factor1, factor2)
		if err != nil {
			return 0, err
		}

		// Check for convergence
		// Using relative error, or absolute error if previousIntegral is close to zero
		if math.Abs(previousIntegral) < 1e-9 { // Avoid division by zero or near-zero
			if math.Abs(currentIntegral-previousIntegral) < tol {
				return currentIntegral, nil
			}
		} else if math.Abs((currentIntegral-previousIntegral)/previousIntegral) < tol {
			return currentIntegral, nil
		}
	}

	return currentIntegral, fmt.Errorf("adaptive Gauss-Legendre failed to converge within %d iterations for tolerance %g (last integral: %f, prev: %f, intervals: %d)", maxIterations, tol, currentIntegral, previousIntegral, numIntervals)
}

// calculateGaussLegendreForIntervals performs the core quadrature sum over a specified number of subintervals.
func (s *GaussLegendreStrategy) calculateGaussLegendreForIntervals(
	fn func(float64) float64,
	globalA, globalB float64,
	n int, // n is the number of Gauss points per subinterval
	roots []float64, // Roots for [-1,1]
	weights []float64, // Weights for [-1,1]
	numIntervals int,
	// factor1 and factor2 are precomputed for the original interval [a,b]
	// these are NOT used directly here, we need to recompute for each subinterval.
	_ float64, // factor1, unused
	_ float64, // factor2, unused
) (float64, error) {
	totalIntegral := 0.0
	subIntervalWidth := (globalB - globalA) / float64(numIntervals)

	for i := 0; i < numIntervals; i++ {
		// Determine the limits for the current subinterval [subA, subB]
		subA := globalA + float64(i)*subIntervalWidth
		subB := subA + subIntervalWidth

		// Transformation factors for the current subinterval [subA, subB]
		subFactor1 := (subB - subA) / 2.0 // (b_i - a_i) / 2
		subFactor2 := (subA + subB) / 2.0 // (a_i + b_i) / 2

		subIntegral := 0.0
		for j := 0; j < n; j++ {
			// Transform root from [-1, 1] to [subA, subB]
			transformedRoot := subFactor1*roots[j] + subFactor2
			subIntegral += weights[j] * fn(transformedRoot)
		}
		totalIntegral += subFactor1 * subIntegral // Don't forget the (b-a)/2 factor for the sum
	}
	return totalIntegral, nil
}
