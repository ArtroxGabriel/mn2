// Package secondorder implements second-order accurate finite difference methods.
package secondorder

// BackwardSecondOrderStrategy implements second-order backward finite difference approximations.
// This method achieves O(h²) accuracy using only backward points, making it suitable for
// boundary conditions where forward points are not available.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [3f(x) - 4f(x-h) + f(x-2h)] / (2h)
//   - Second derivative: f”(x) ≈ [f(x) - 2f(x-h) + f(x-2h)] / h²
//   - Third derivative: f”'(x) ≈ [f(x) - 3f(x-h) + 3f(x-2h) - f(x-3h)] / h³
//
// Characteristics:
//   - Accuracy: O(h²) for first derivative, O(h) for higher derivatives
//   - Stability: Good for smooth functions
//   - Requirements: Function evaluations only at x and backward points
//   - Use case: Right boundary conditions or when forward points unavailable
type BackwardSecondOrderStrategy struct{}

// CalculateFirst computes the first derivative using second-order backward difference.
// Formula: f'(x) ≈ [3f(x) - 4f(x-h) + f(x-2h)] / (2h)
//
// This three-point backward formula achieves O(h²) accuracy by using a weighted
// combination of function values. The coefficients are the mirror image of the
// forward second-order formula, maintaining the same accuracy.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (f *BackwardSecondOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (3*fn(x) - 4*fn(x-h) + fn(x-2*h)) / (2 * h)
}

// CalculateSecond computes the second derivative using backward difference.
// Formula: f”(x) ≈ [f(x) - 2f(x-h) + f(x-2h)] / h²
//
// This is the backward equivalent of the standard three-point second derivative
// formula, using only backward points for evaluation.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (f *BackwardSecondOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - 2*fn(x-h) + fn(x-2*h)) / (h * h)
}

// CalculateThirty computes the third derivative using backward difference.
// Formula: f”'(x) ≈ [f(x) - 3f(x-h) + 3f(x-2h) - f(x-3h)] / h³
//
// This four-point backward formula for the third derivative uses coefficients
// that follow the pattern of backward finite differences.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (b *BackwardSecondOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - 3*fn(x-h) + 3*fn(x-2*h) - fn(x-3*h)) / (h * h * h)
}
