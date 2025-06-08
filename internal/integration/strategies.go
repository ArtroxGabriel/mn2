package integration

// IntegrationStrategy defines the interface that all numerical integration strategies must implement.
// This interface follows the Strategy design pattern, allowing different methods
// to be used interchangeably.
type IntegrationStrategy interface {
	// Calculate computes the definite integral of the function f(x) from a to b
	// using n points or a polynomial of degree n, aiming for a specified tolerance.
	//
	// Parameters:
	//   - fn: The function to integrate.
	//   - a: The lower limit of integration.
	//   - b: The upper limit of integration.
	//   - n: The number of points or degree of the polynomial, depending on the method.
	//   - tol: The desired error tolerance.
	//
	// Returns:
	//   - float64: The approximated value of the integral.
	//   - error: An error if the calculation fails or parameters are invalid.
	Calculate(
		fn func(float64) float64,
		a, b float64,
		n int,
		tol float64,
	) (float64, error)
}
