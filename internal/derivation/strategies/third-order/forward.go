// Package thirdorder implements third-order accurate finite difference methods.
// These methods have O(h³) truncation error, providing higher accuracy at the cost
// of more function evaluations.
package thirdorder

// ForwardThirOrderStrategy implements third-order forward finite difference approximations.
// This method achieves O(h³) accuracy using only forward points, making it suitable for
// left boundary conditions where backward points are not available.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [-f(x+3h) + 6f(x+2h) - 18f(x+h) - 10f(x)] / (12h)
//   - Second derivative: f”(x) ≈ [-f(x+3h) + 4f(x+2h) - 5f(x+h) + 2f(x)] / h²
//   - Third derivative: f”'(x) ≈ [-3f(x+4h) + 14f(x+3h) - 24f(x+2h) + 18f(x+h) - 5f(x)] / (2h³)
//
// Characteristics:
//   - Accuracy: O(h³) for first derivative, O(h²) for second, O(h) for third
//   - Stability: Good for smooth functions, may be sensitive to noise
//   - Requirements: Function evaluations at x, x+h, x+2h, x+3h (and x+4h for third derivative)
//   - Use case: High-accuracy requirements at left boundaries
type ForwardThirOrderStrategy struct{}

// CalculateFirst computes the first derivative using third-order forward difference.
// Formula: f'(x) ≈ [-f(x+3h) + 6f(x+2h) - 18f(x+h) - 10f(x)] / (12h)
//
// This four-point forward formula achieves O(h³) accuracy by carefully choosing
// coefficients that eliminate lower-order error terms. The large coefficients
// may amplify rounding errors, so h should be chosen carefully.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (f *ForwardThirOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 6*fn(x+2*h) - 18*fn(x+h) - 10*fn(x)) / (12 * h)
}

// CalculateSecond computes the second derivative using third-order forward difference.
// Formula: f”(x) ≈ [-f(x+3h) + 4f(x+2h) - 5f(x+h) + 2f(x)] / h²
//
// This four-point formula provides O(h²) accuracy for the second derivative
// using only forward points.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (f *ForwardThirOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 4*fn(x+2*h) - 5*fn(x+h) + 2*fn(x)) / (h * h)
}

// CalculateThirty computes the third derivative using third-order forward difference.
// Formula: f”'(x) ≈ [-3f(x+4h) + 14f(x+3h) - 24f(x+2h) + 18f(x+h) - 5f(x)] / (2h³)
//
// This five-point forward formula for the third derivative requires evaluation
// at x, x+h, x+2h, x+3h, and x+4h.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (f *ForwardThirOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-3*fn(x+4*h) + 14*fn(x+3*h) - 24*fn(x+2*h) + 18*fn(x+h) - 5*fn(x)) / (2 * h * h * h)
}
