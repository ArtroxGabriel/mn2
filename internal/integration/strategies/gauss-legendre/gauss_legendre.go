package gausslegendre

import (
	"fmt"
	"math"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration"
)

// Abscissas and weights for Gauss-Legendre quadrature
// Source: Numerical Recipes and Wikipedia
var (
	gaussLegendreAbscissas = map[int][]float64{
		1: {0.0},
		2: {-1.0 / math.Sqrt(3), 1.0 / math.Sqrt(3)}, // approx +/- 0.5773502692
		3: {-math.Sqrt(3.0 / 5.0), 0.0, math.Sqrt(3.0 / 5.0)}, // approx +/- 0.7745966692, 0
		4: { // In increasing order for clarity with weights
			-math.Sqrt((3.0/7.0) + (2.0/7.0)*math.Sqrt(6.0/5.0)), // approx -0.8611363116
			-math.Sqrt((3.0/7.0) - (2.0/7.0)*math.Sqrt(6.0/5.0)), // approx -0.3399810436
			math.Sqrt((3.0/7.0) - (2.0/7.0)*math.Sqrt(6.0/5.0)),  // approx  0.3399810436
			math.Sqrt((3.0/7.0) + (2.0/7.0)*math.Sqrt(6.0/5.0)),  // approx  0.8611363116
		},
		// Add order 5 for completeness if ever needed, from reliable source
		// 5: {
		// 	-math.Sqrt( (5.0 + 2.0*math.Sqrt(10.0/7.0)) / 9.0 ),
		// 	-math.Sqrt( (5.0 - 2.0*math.Sqrt(10.0/7.0)) / 9.0 ),
		// 	0.0,
		// 	math.Sqrt( (5.0 - 2.0*math.Sqrt(10.0/7.0)) / 9.0 ),
		// 	math.Sqrt( (5.0 + 2.0*math.Sqrt(10.0/7.0)) / 9.0 ),
		// },
	}
	gaussLegendreWeights = map[int][]float64{
		1: {2.0},
		2: {1.0, 1.0},
		3: {5.0 / 9.0, 8.0 / 9.0, 5.0 / 9.0}, // Corresponds to -sqrt(3/5), 0, +sqrt(3/5)
		4: { // Corresponds to the ordered abscissas for n=4
			(18.0 - math.Sqrt(30.0)) / 36.0, // approx 0.3478548451 for -0.8611...
			(18.0 + math.Sqrt(30.0)) / 36.0, // approx 0.6521451549 for -0.3399...
			(18.0 + math.Sqrt(30.0)) / 36.0, // approx 0.6521451549 for  0.3399...
			(18.0 - math.Sqrt(30.0)) / 36.0, // approx 0.3478548451 for  0.8611...
		},
		// Weights for n=5
		// 5: {
		// 	(322.0 - 13.0*math.Sqrt(70.0))/900.0,
		// 	(322.0 + 13.0*math.Sqrt(70.0))/900.0,
		// 	128.0/225.0,
		// 	(322.0 + 13.0*math.Sqrt(70.0))/900.0,
		// 	(322.0 - 13.0*math.Sqrt(70.0))/900.0,
		// },
	}
)

// generalGaussLegendre performs the core integration logic for a given order.
// The 'order' parameter specifies the number of points for the quadrature.
func generalGaussLegendre(fn common.MathFunc, a, b float64, order int) (float64, error) {
	abscissas, okA := gaussLegendreAbscissas[order]
	weights, okW := gaussLegendreWeights[order]

	if !okA || !okW {
		// This check is important as the maps only contain specific orders.
		return 0, fmt.Errorf("Gauss-Legendre quadrature is not defined for order %d. Supported orders: 1, 2, 3, 4", order)
	}

	if a == b {
		return 0, nil
	}

	// Standard transformation for Gauss-Legendre:
	// Integral from a to b of f(x)dx = (b-a)/2 * Sum[w_i * f( (b-a)/2 * x_i + (a+b)/2 )]
	// where x_i are abscissas for interval [-1, 1] and w_i are corresponding weights.
	term1 := (b - a) / 2.0
	term2 := (a + b) / 2.0

	sum := 0.0
	for i := 0; i < order; i++ {
		sum += weights[i] * fn(term1*abscissas[i]+term2)
	}

	return term1 * sum, nil
}

// GaussLegendreOrder1 strategy implements Gauss-Legendre quadrature with 1 point.
type GaussLegendreOrder1 struct{}

// Integrate computes the definite integral using 1-point Gauss-Legendre quadrature.
// The 'n' parameter from the interface is expected to be 1 for this strategy.
func (s *GaussLegendreOrder1) Integrate(fn common.MathFunc, a, b float64, n_points int) (float64, error) {
	// n_points is the number of points, which for this type is fixed at 1.
	// We could add validation: if n_points != 1 && n_points != 0 { error }
	// However, the factory should instantiate the correct type based on desired points.
	// So, we use the fixed order of the struct.
	return generalGaussLegendre(fn, a, b, 1)
}

// GaussLegendreOrder2 strategy implements Gauss-Legendre quadrature with 2 points.
type GaussLegendreOrder2 struct{}

// Integrate computes the definite integral using 2-point Gauss-Legendre quadrature.
func (s *GaussLegendreOrder2) Integrate(fn common.MathFunc, a, b float64, n_points int) (float64, error) {
	return generalGaussLegendre(fn, a, b, 2)
}

// GaussLegendreOrder3 strategy implements Gauss-Legendre quadrature with 3 points.
type GaussLegendreOrder3 struct{}

// Integrate computes the definite integral using 3-point Gauss-Legendre quadrature.
func (s *GaussLegendreOrder3) Integrate(fn common.MathFunc, a, b float64, n_points int) (float64, error) {
	return generalGaussLegendre(fn, a, b, 3)
}

// GaussLegendreOrder4 strategy implements Gauss-Legendre quadrature with 4 points.
type GaussLegendreOrder4 struct{}

// Integrate computes the definite integral using 4-point Gauss-Legendre quadrature.
func (s *GaussLegendreOrder4) Integrate(fn common.MathFunc, a, b float64, n_points int) (float64, error) {
	return generalGaussLegendre(fn, a, b, 4)
}

// Ensure strategies implement the IntegrationStrategy interface
var _ integration.IntegrationStrategy = (*GaussLegendreOrder1)(nil)
var _ integration.IntegrationStrategy = (*GaussLegendreOrder2)(nil)
var _ integration.IntegrationStrategy = (*GaussLegendreOrder3)(nil)
var _ integration.IntegrationStrategy = (*GaussLegendreOrder4)(nil)
