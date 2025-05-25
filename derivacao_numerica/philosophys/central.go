package philosophys

import (
	"context"
	"log/slog"
)

// Central implements the central difference method for numerical differentiation.
//
// Mathematical Formula: f'(x) ≈ [f(x+h) - f(x-h)] / (2h)
//
// This method approximates the first derivative using function values at x+h and x-h.
// It has a truncation error of order O(h²), making it significantly more accurate
// than forward or backward differences for the same step size.
//
// Advantages:
// - Higher accuracy (O(h²) vs O(h))
// - More stable numerical behavior
// - Symmetric around the point of interest
// - Better convergence properties
//
// Disadvantages:
// - Requires function evaluation at both x+h and x-h
// - May not be suitable for boundary conditions
// - Slightly more computational cost
type Central struct {
	fn func(float64) float64 // Function to differentiate
}

// DerivadaPrimeira computes the central difference approximation of the derivative.
// Parameters:
//   - x: Point at which to compute the derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated derivative value f'(x)
func (f Central) DerivadaPrimeira(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating central difference",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "central"),
		slog.String("accuracy", "O(h²)"))
	return (f.fn(x+delta) - f.fn(x-delta)) / (2 * delta)
}

// DerivadaSegunda computes the central difference approximation of the second derivative.
//
// Mathematical Formula: f”(x) ≈ [f(x+h) - 2f(x) + f(x-h)] / h²
//
// This method approximates the second derivative using function values at x-h, x, and x+h.
// It has a truncation error of order O(h²), providing good accuracy for the second derivative.
//
// The formula is derived from the Taylor series expansion and represents the discrete
// analog of the continuous second derivative.
//
// Advantages:
// - Symmetric formulation provides balanced approximation
// - O(h²) accuracy, better than forward/backward methods
// - Stable numerical behavior for reasonable step sizes
// - Natural extension of central difference philosophy
//
// Disadvantages:
// - Requires three function evaluations (x-h, x, x+h)
// - May amplify numerical errors for very small h
// - Not suitable for boundary points without modification
//
// Parameters:
//   - x: Point at which to compute the second derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated second derivative value f”(x)
func (f Central) DerivadaSegunda(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating central difference for second derivative",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "central"),
		slog.String("accuracy", "O(h²)"),
		slog.String("derivative_order", "second"))

	return (f.fn(x+delta) - 2*f.fn(x) + f.fn(x-delta)) / (delta * delta)
}
