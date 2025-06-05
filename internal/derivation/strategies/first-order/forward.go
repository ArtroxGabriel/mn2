// Package firstorder implements first-order accurate finite difference methods.
// These methods have O(h) truncation error, meaning the error decreases linearly with step size.
package firstorder

// ForwardFirstOrderStrategy implements first-order forward finite difference approximations.
// This method uses function evaluations at x and x+h to approximate derivatives.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [f(x+h) - f(x)] / h
//   - Second derivative: f”(x) ≈ [f(x+2h) - 2f(x+h) + f(x)] / h²
//   - Third derivative: f”'(x) ≈ [f(x+3h) - 3f(x+2h) + 3f(x+h) - f(x)] / h³
//
// Characteristics:
//   - Accuracy: O(h) for first derivative, O(h) for higher derivatives
//   - Stability: Good for smooth functions
//   - Requirements: Function evaluations only at x and forward points
//   - Use case: When backward points are not available or near boundaries
type ForwardFirstOrderStrategy struct{}

// CalculateFirst computes the first derivative using first-order forward difference.
// Formula: f'(x) ≈ [f(x+h) - f(x)] / h
//
// This is the simplest forward difference approximation with O(h) truncation error.
// It requires only two function evaluations and is suitable when only forward
// points are available (e.g., at the left boundary of a domain).
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be small but not too small to avoid numerical errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (f *ForwardFirstOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+h) - fn(x)) / h
}

// CalculateSecond computes the second derivative using forward difference.
// Formula: f”(x) ≈ [f(x+2h) - 2f(x+h) + f(x)] / h²
//
// This approximation uses three function evaluations at x, x+h, and x+2h.
// The accuracy is O(h) due to the forward difference nature.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (f *ForwardFirstOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+2*h) - 2*fn(x+h) + fn(x)) / (h * h)
}

// CalculateThirty computes the third derivative using forward difference.
// Formula: f”'(x) ≈ [f(x+3h) - 3f(x+2h) + 3f(x+h) - f(x)] / h³
//
// This approximation uses four function evaluations at x, x+h, x+2h, and x+3h.
// The coefficients follow the pattern of finite differences.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (f *ForwardFirstOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+3*h) - 3*fn(x+2*h) + 3*fn(x+h) - fn(x)) / (h * h * h)
}
