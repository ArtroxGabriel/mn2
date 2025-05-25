package philosophys

import (
	"context"
	"log/slog"
)

// Backward implements the backward difference method for numerical differentiation.
//
// Mathematical Formula: f'(x) ≈ [f(x) - f(x-h)] / h
//
// This method approximates the first derivative using function values at x and x-h.
// It has a truncation error of order O(h), same as the forward difference method.
//
// Advantages:
// - Simple implementation
// - Only requires function evaluation at x and x-h
// - Suitable when forward evaluation is not possible
// - Useful for boundary conditions
//
// Disadvantages:
// - Lower accuracy compared to central difference
// - Can be less stable for small h values
type Backward struct {
	fn func(float64) float64 // Function to differentiate
}

// DerivadaPrimeira computes the backward difference approximation of the derivative.
// Parameters:
//   - x: Point at which to compute the derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated derivative value f'(x)
func (f Backward) DerivadaPrimeira(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating backward difference",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "backward"),
		slog.String("accuracy", "O(h)"))
	return (f.fn(x) - f.fn(x-delta)) / delta
}

func (f Backward) DerivadaSegunda(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating backward difference for second derivative",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "backward"),
		slog.String("accuracy", "O(h)"))

	return (f.fn(x) - 2*f.fn(x-delta) + f.fn(x-2*delta)) / (delta * delta)
}
