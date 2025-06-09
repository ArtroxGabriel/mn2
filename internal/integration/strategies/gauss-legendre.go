package strategies

import (
	"math"
)

// GaussLegendreOrder1 implements 1-point Gauss-Legendre quadrature.
type GaussLegendreOrder1 struct{}

// Calculate computes the definite integral using 1-point Gauss-Legendre quadrature.
// Formula: ∫[-1,1] f(x) dx ≈ w1*f(x1)
// For a general interval [a,b]: ∫[a,b] f(t) dt ≈ ((b-a)/2) * Σ[i=1 to n] wi*f( ((b-a)/2)*xi + ((a+b)/2) )
func (s *GaussLegendreOrder1) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	// n is not used for non-composite Gauss-Legendre, but included for interface compliance.
	// We can choose to interpret 'n' as the order for Gauss-Legendre if desired,
	// but here we stick to specific order structs.
	if a == b {
		return 0, nil
	}
	// Transform the interval [a,b] to [-1,1] for Gauss-Legendre
	// t = ((b-a)/2)*x + ((a+b)/2)
	// dt = ((b-a)/2)*dx
	// ∫[a,b] f(t) dt = ∫[-1,1] f(((b-a)/2)*x + ((a+b)/2)) * ((b-a)/2) dx

	// 1-point formula: x1 = 0, w1 = 2
	x1 := 0.0
	w1 := 2.0

	term1 := w1 * fn(((b-a)/2)*x1+((a+b)/2))

	return ((b - a) / 2) * term1, nil
}

// GaussLegendreOrder2 implements 2-point Gauss-Legendre quadrature.
type GaussLegendreOrder2 struct{}

// Calculate computes the definite integral using 2-point Gauss-Legendre quadrature.
func (s *GaussLegendreOrder2) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if a == b {
		return 0, nil
	}
	// 2-point formula: x1 = -1/sqrt(3), x2 = 1/sqrt(3); w1 = 1, w2 = 1
	x_coords := []float64{-1 / math.Sqrt(3), 1 / math.Sqrt(3)}
	weights := []float64{1.0, 1.0}

	sum := 0.0
	for i := range x_coords {
		sum += weights[i] * fn(((b-a)/2)*x_coords[i]+((a+b)/2))
	}

	return ((b - a) / 2) * sum, nil
}

// GaussLegendreOrder3 implements 3-point Gauss-Legendre quadrature.
type GaussLegendreOrder3 struct{}

// Calculate computes the definite integral using 3-point Gauss-Legendre quadrature.
func (s *GaussLegendreOrder3) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if a == b {
		return 0, nil
	}
	// 3-point formula:
	// x1 = 0, x2 = -sqrt(3/5), x3 = sqrt(3/5)
	// w1 = 8/9, w2 = 5/9, w3 = 5/9
	x_coords := []float64{-math.Sqrt(3.0 / 5.0), 0.0, math.Sqrt(3.0 / 5.0)}
	weights := []float64{5.0 / 9.0, 8.0 / 9.0, 5.0 / 9.0}

	sum := 0.0
	for i := range x_coords {
		sum += weights[i] * fn(((b-a)/2)*x_coords[i]+((a+b)/2))
	}

	return ((b - a) / 2) * sum, nil
}

// GaussLegendreOrder4 implements 4-point Gauss-Legendre quadrature.
type GaussLegendreOrder4 struct{}

// Calculate computes the definite integral using 4-point Gauss-Legendre quadrature.
func (s *GaussLegendreOrder4) Calculate(fn func(float64) float64, a, b float64, n int) (float64, error) {
	if a == b {
		return 0, nil
	}
	// 4-point formula:
	// x1,2 = ±sqrt((3 - 2*sqrt(6/5))/7)
	// x3,4 = ±sqrt((3 + 2*sqrt(6/5))/7)
	// w1,2 = (18 + sqrt(30))/36
	// w3,4 = (18 - sqrt(30))/36
	x_coords := []float64{
		-math.Sqrt((3 + 2*math.Sqrt(6.0/5.0)) / 7.0),
		-math.Sqrt((3 - 2*math.Sqrt(6.0/5.0)) / 7.0),
		math.Sqrt((3 - 2*math.Sqrt(6.0/5.0)) / 7.0),
		math.Sqrt((3 + 2*math.Sqrt(6.0/5.0)) / 7.0),
	}
	weights := []float64{
		(18 - math.Sqrt(30)) / 36.0,
		(18 + math.Sqrt(30)) / 36.0,
		(18 + math.Sqrt(30)) / 36.0,
		(18 - math.Sqrt(30)) / 36.0,
	}

	sum := 0.0
	for i := range x_coords {
		sum += weights[i] * fn(((b-a)/2)*x_coords[i]+((a+b)/2))
	}

	return ((b - a) / 2) * sum, nil
}
