// Package secondorder implements second-order accurate finite difference methods.
// These methods have O(h²) truncation error, providing better accuracy than first-order methods.
package secondorder

// CentralSecondOrderStrategy implements second-order central finite difference approximations.
// This method uses symmetric function evaluations around the point of interest for better accuracy.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [f(x+h) - f(x-h)] / (2h)
//   - Second derivative: f”(x) ≈ [f(x+h) - 2f(x) + f(x-h)] / h²
//   - Third derivative: f”'(x) ≈ [f(x-2h) - 2f(x-h) + 2f(x+h) - f(x+2h)] / (2h³)
//
// Characteristics:
//   - Accuracy: O(h²) for first and second derivatives
//   - Stability: Excellent for smooth functions
//   - Requirements: Function evaluations at both forward and backward points
//   - Use case: Best choice when function can be evaluated at x±h points
type CentralSecondOrderStrategy struct{}

// CalculateFirst computes the first derivative using second-order central difference.
// Formula: f'(x) ≈ [f(x+h) - f(x-h)] / (2h)
//
// This is the most commonly used finite difference approximation due to its
// superior accuracy (O(h²)) compared to forward/backward methods. It uses
// symmetric points around x, which cancels out the first-order error terms.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (c *CentralSecondOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x+h) - fn(x-h)) / (2 * h)
}

// CalculateSecond computes the second derivative using central difference.
// Formula: f”(x) ≈ [f(x+h) - 2f(x) + f(x-h)] / h²
//
// This is the standard three-point formula for second derivatives.
// It has O(h²) accuracy and is widely used in numerical analysis.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (c *CentralSecondOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x+h) - 2*fn(x) + fn(x-h)) / (h * h)
}

// CalculateThirty computes the third derivative using central difference.
// Formula: f”'(x) ≈ [f(x-2h) - 2f(x-h) + 2f(x+h) - f(x+2h)] / (2h³)
//
// This approximation uses four symmetric function evaluations around x.
// Note that f(x) does not appear in this formula due to the odd nature
// of the third derivative.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (c *CentralSecondOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x-2*h) - 2*fn(x-h) + 2*fn(x+h) - fn(x+2*h)) / (2 * h * h * h)
}
