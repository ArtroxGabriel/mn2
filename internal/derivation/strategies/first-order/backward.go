// Package firstorder implements first-order accurate finite difference methods.
package firstorder

// BackwardFirstOrderStrategy implements first-order backward finite difference approximations.
// This method uses function evaluations at x and x-h to approximate derivatives.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [f(x) - f(x-h)] / h
//   - Second derivative: f”(x) ≈ [f(x) - 2f(x-h) + f(x-2h)] / h²
//   - Third derivative: f”'(x) ≈ [f(x) - 3f(x-h) + 3f(x-2h) - f(x-3h)] / h³
//
// Characteristics:
//   - Accuracy: O(h) for first derivative, O(h) for higher derivatives
//   - Stability: Good for smooth functions
//   - Requirements: Function evaluations only at x and backward points
//   - Use case: When forward points are not available or near boundaries
type BackwardFirstOrderStrategy struct{}

// CalculateFirst computes the first derivative using first-order backward difference.
// Formula: f'(x) ≈ [f(x) - f(x-h)] / h
//
// This is the backward equivalent of forward difference with O(h) truncation error.
// It requires only two function evaluations and is suitable when only backward
// points are available (e.g., at the right boundary of a domain).
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (b *BackwardFirstOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - fn(x-h)) / h
}

// CalculateSecond computes the second derivative using backward difference.
// Formula: f”(x) ≈ [f(x) - 2f(x-h) + f(x-2h)] / h²
//
// This approximation uses three function evaluations at x, x-h, and x-2h.
// The accuracy is O(h) due to the backward difference nature.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (b *BackwardFirstOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - 2*fn(x-h) + fn(x-2*h)) / (h * h)
}

// CalculateThirty computes the third derivative using backward difference.
// Formula: f”'(x) ≈ [f(x) - 3*fn(x-h) + 3*fn(x-2h) - f(x-3h)] / h³
//
// This approximation uses four function evaluations at x, x-h, x-2h, and x-3h.
// The coefficients follow the pattern of backward finite differences.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (b *BackwardFirstOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - 3*fn(x-h) + 3*fn(x-2*h) - fn(x-3*h)) / (h * h * h)
}
