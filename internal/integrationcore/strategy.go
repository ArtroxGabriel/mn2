package integrationcore

// IntegrationStrategy defines the interface that all numerical integration strategies must implement.
// This interface follows the Strategy design pattern, allowing different integration methods
// to be used interchangeably.
type IntegrationStrategy interface {
	// Calculate computes the definite integral of the function fn from a to b
	// using n subintervals (if applicable to the method).
	//
	// Parameters:
	//   - fn: The function to integrate (must be a single-variable function)
	//   - a: The lower limit of integration
	//   - b: The upper limit of integration
	//   - n: The number of subintervals (for composite rules)
	//
	// Returns:
	//   - float64: The approximated definite integral value
	//   - error: An error if the calculation fails (e.g., invalid input)
	Calculate(fn func(float64) float64, a, b float64, n int) (float64, error)
}
