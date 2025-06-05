// Package thirdorder implements third-order accurate finite difference methods.
package thirdorder

// BackwardThirOrderStrategy implements third-order backward finite difference approximations.
// This method achieves O(h³) accuracy using only backward points, making it suitable for
// right boundary conditions where forward points are not available.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [10f(x) - 18f(x-h) + 6f(x-2h) - f(x-3h)] / (12h)
//   - Second derivative: f”(x) ≈ [2f(x) - 5f(x-h) + 4f(x-2h) - f(x-3h)] / h²
//   - Third derivative: f”'(x) ≈ [5f(x) - 18f(x-h) + 24f(x-2h) - 14f(x-3h) + 3f(x-4h)] / (2h³)
//
// Characteristics:
//   - Accuracy: O(h³) for first derivative, O(h²) for second, O(h) for third
//   - Stability: Good for smooth functions, may be sensitive to noise
//   - Requirements: Function evaluations at x, x-h, x-2h, x-3h (and x-4h for third derivative)
//   - Use case: High-accuracy requirements at right boundaries
type BackwardThirOrderStrategy struct{}

// CalculateFirst computes the first derivative using third-order backward difference.
// Formula: f'(x) ≈ [10f(x) - 18f(x-h) + 6f(x-2h) - f(x-3h)] / (12h)
//
// This four-point backward formula achieves O(h³) accuracy and is the mirror
// image of the forward third-order formula. The coefficients are carefully
// chosen to eliminate lower-order error terms.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (b *BackwardThirOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (10*fn(x) - 18*fn(x-h) + 6*fn(x-2*h) - fn(x-3*h)) / (12 * h)
}

// CalculateSecond computes the second derivative using third-order backward difference.
// Formula: f”(x) ≈ [2f(x) - 5f(x-h) + 4f(x-2h) - f(x-3h)] / h²
//
// This four-point backward formula provides O(h²) accuracy for the second
// derivative using only backward points.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (b *BackwardThirOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (2*fn(x) - 5*fn(x-h) + 4*fn(x-2*h) - fn(x-3*h)) / (h * h)
}

// CalculateThirty computes the third derivative using third-order backward difference.
// Formula: f”'(x) ≈ [5f(x) - 18f(x-h) + 24f(x-2h) - 14f(x-3h) + 3f(x-4h)] / (2h³)
//
// This five-point backward formula for the third derivative requires evaluation
// at x, x-h, x-2h, x-3h, and x-4h.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (b *BackwardThirOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (5*fn(x) - 18*fn(x-h) + 24*fn(x-2*h) - 14*fn(x-3*h) + 3*fn(x-4*h)) / (2 * h * h * h)
}
