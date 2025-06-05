// Package secondorder implements second-order accurate finite difference methods.
package secondorder

// ForwardSecondOrderStrategy implements second-order forward finite difference approximations.
// This method achieves O(h²) accuracy using only forward points, making it suitable for
// boundary conditions where backward points are not available.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [-3f(x) + 4f(x+h) - f(x+2h)] / (2h)
//   - Second derivative: f”(x) ≈ [f(x+2h) - 2f(x+h) + f(x)] / h²
//   - Third derivative: f”'(x) ≈ [-f(x+3h) + 3f(x+2h) - 3f(x+h) + f(x)] / h³
//
// Characteristics:
//   - Accuracy: O(h²) for first derivative, O(h) for higher derivatives
//   - Stability: Good for smooth functions
//   - Requirements: Function evaluations only at x and forward points
//   - Use case: Left boundary conditions or when backward points unavailable
type ForwardSecondOrderStrategy struct{}

// CalculateFirst computes the first derivative using second-order forward difference.
// Formula: f'(x) ≈ [-3f(x) + 4f(x+h) - f(x+2h)] / (2h)
//
// This three-point forward formula achieves O(h²) accuracy by using a weighted
// combination of function values. The coefficients are derived from Taylor
// series expansion to eliminate the O(h) error term.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (f *ForwardSecondOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-3*fn(x) + 4*fn(x+h) - fn(x+2*h)) / (2 * h)
}

// CalculateSecond computes the second derivative using forward difference.
// Formula: f”(x) ≈ [f(x+2h) - 2f(x+h) + f(x)] / h²
//
// This is the same formula as the first-order forward second derivative,
// maintaining the standard three-point second derivative pattern.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (f *ForwardSecondOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x+2*h) - 2*fn(x+h) + fn(x)) / (h * h)
}

// CalculateThirty computes the third derivative using forward difference.
// Formula: f”'(x) ≈ [-f(x+3h) + 3f(x+2h) - 3f(x+h) + f(x)] / h³
//
// This four-point forward formula for the third derivative uses coefficients
// that follow the pattern of forward finite differences.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (f *ForwardSecondOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 3*fn(x+2*h) - 3*fn(x+h) + fn(x)) / (h * h * h)
}
