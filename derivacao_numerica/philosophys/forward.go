package philosophys

import (
	"context"
	"log/slog"
)

// Forward implements the forward difference method for numerical differentiation.
//
// Mathematical Formula: f'(x) ≈ [f(x+h) - f(x)] / h
//
// This method approximates the first derivative using function values at x and x+h.
// It has a truncation error of order O(h), meaning the error decreases linearly
// with the step size h.
//
// Advantages:
// - Simple implementation
// - Only requires function evaluation at x and x+h
// - Suitable when backward evaluation is not possible
//
// Disadvantages:
// - Lower accuracy compared to central difference
// - Can be less stable for small h values
type Forward struct {
	fn func(float64) float64 // Function to differentiate
}

// DerivadaPrimeira computes the forward difference approximation of the derivative.
// Parameters:
//   - x: Point at which to compute the derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated derivative value f'(x)
func (f Forward) DerivadaPrimeira(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating forward difference",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "forward"),
		slog.String("accuracy", "O(h)"))
	return (f.fn(x+delta) - f.fn(x)) / delta
}

// DerivadaSegunda computes the forward difference approximation of the second derivative.
//
// Mathematical Formula: f”(x) ≈ [f(x+2h) - 2f(x+h) + f(x)] / h²
//
// This method approximates the second derivative using function values at x, x+h, and x+2h.
// It has a truncation error of order O(h), which is lower accuracy compared to central difference.
//
// The formula uses three forward points to construct a second-order difference approximation.
// This approach is particularly useful when backward evaluations are not available or
// when working near the left boundary of a domain.
//
// Advantages:
// - Only requires forward function evaluations
// - Suitable for boundary conditions at the left edge
// - Simple implementation and computation
// - Useful when f(x-h) is not available or meaningful
//
// Disadvantages:
// - Lower accuracy O(h) compared to central difference O(h²)
// - May have larger truncation errors
// - Can be less stable for small step sizes
// - Requires evaluation at x+2h which extends the computational domain
//
// Parameters:
//   - x: Point at which to compute the second derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated second derivative value f”(x)
func (f Forward) DerivadaSegunda(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating forward difference for second derivative",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "forward"),
		slog.String("accuracy", "O(h)"),
		slog.String("derivative_order", "second"))

	return (f.fn(x+2*delta) - 2*f.fn(x+delta) + f.fn(x)) / (delta * delta)
}
