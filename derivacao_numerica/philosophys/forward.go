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

func (f Forward) DerivadaSegunda(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating forward difference for second derivative",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "forward"),
		slog.String("accuracy", "O(h)"))

	return (f.fn(x+2*delta) - 2*f.fn(x+delta) + f.fn(x)) / (delta * delta)
}
