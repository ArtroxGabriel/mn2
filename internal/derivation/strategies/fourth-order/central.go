// Package fourthorder implements fourth-order accurate finite difference methods.
// These methods have O(h⁴) truncation error, providing the highest accuracy available
// in this implementation at the cost of more function evaluations.
package fourthorder

// CentralFourthOrderStrategy implements fourth-order central finite difference approximations.
// This method achieves O(h⁴) accuracy using symmetric function evaluations, providing
// the best balance of accuracy and stability for smooth functions.
//
// Mathematical formulations:
//   - First derivative: f'(x) ≈ [-f(x+2h) + 8f(x+h) - 8f(x-h) + f(x-2h)] / (12h)
//   - Second derivative: f”(x) ≈ [-f(x+2h) + 16f(x+h) - 30f(x) + 16f(x-h) - f(x-2h)] / (12h²)
//   - Third derivative: f”'(x) ≈ [-f(x+3h) + 8f(x+2h) - 13f(x+h) + 13f(x-h) - 8f(x-2h) + f(x-3h)] / (8h³)
//
// Characteristics:
//   - Accuracy: O(h⁴) for first and second derivatives, O(h²) for third derivative
//   - Stability: Excellent for smooth functions, best choice for high accuracy
//   - Requirements: Function evaluations at x±h, x±2h (and x±3h for third derivative)
//   - Use case: When highest accuracy is needed and function is sufficiently smooth
type CentralFourthOrderStrategy struct{}

// CalculateFirst computes the first derivative using fourth-order central difference.
// Formula: f'(x) ≈ [-f(x+2h) + 8f(x+h) - 8f(x-h) + f(x-2h)] / (12h)
//
// This five-point central formula achieves O(h⁴) accuracy, making it the most
// accurate first derivative approximation in this implementation. The symmetric
// nature cancels out even-order error terms, leaving only O(h⁴) truncation error.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size (should be chosen to balance truncation and rounding errors)
//
// Returns:
//   - float64: Approximated first derivative value
func (c *CentralFourthOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+2*h) + 8*fn(x+h) - 8*fn(x-h) + fn(x-2*h)) / (12 * h)
}

// CalculateSecond computes the second derivative using fourth-order central difference.
// Formula: f”(x) ≈ [-f(x+2h) + 16f(x+h) - 30f(x) + 16f(x-h) - f(x-2h)] / (12h²)
//
// This five-point central formula for the second derivative also achieves O(h⁴)
// accuracy. Note the large coefficient (-30) at f(x), which is characteristic
// of high-order central difference formulas for even derivatives.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated second derivative value
func (c *CentralFourthOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+2*h) + 16*fn(x+h) - 30*fn(x) + 16*fn(x-h) - fn(x-2*h)) / (12 * h * h)
}

// CalculateThirty computes the third derivative using fourth-order central difference.
// Formula: f”'(x) ≈ [-f(x+3h) + 8f(x+2h) - 13f(x+h) + 13f(x-h) - 8f(x-2h) + f(x-3h)] / (8h³)
//
// This six-point central formula achieves O(h²) accuracy for the third derivative.
// Note that f(x) does not appear due to the odd symmetry of the third derivative.
// The formula requires evaluation at six points: x±h, x±2h, x±3h.
//
// Parameters:
//   - fn: Function to differentiate
//   - x: Point of evaluation
//   - h: Step size
//
// Returns:
//   - float64: Approximated third derivative value
func (c *CentralFourthOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 8*fn(x+2*h) - 13*fn(x+h) + 13*fn(x-h) - 8*fn(x-2*h) + fn(x-3*h)) / (8 * h * h * h)
}
