package main

import (
	"context"
	"log/slog"
)

// Derivacao defines the interface for numerical differentiation methods.
// All derivative calculation methods must implement this interface.
type Derivacao interface {
	// derivadaPrimeira computes the numerical derivative at point x using step size delta.
	// Returns the approximated derivative value.
	derivadaPrimeira(x, delta float64) float64
}

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

// derivadaPrimeira computes the forward difference approximation of the derivative.
// Parameters:
//   - x: Point at which to compute the derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated derivative value f'(x)
func (f Forward) derivadaPrimeira(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating forward difference",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "forward"),
		slog.String("accuracy", "O(h)"))
	return (f.fn(x+delta) - f.fn(x)) / delta
}

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

// derivadaPrimeira computes the backward difference approximation of the derivative.
// Parameters:
//   - x: Point at which to compute the derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated derivative value f'(x)
func (f Backward) derivadaPrimeira(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating backward difference",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "backward"),
		slog.String("accuracy", "O(h)"))
	return (f.fn(x) - f.fn(x-delta)) / delta
}

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

// derivadaPrimeira computes the central difference approximation of the derivative.
// Parameters:
//   - x: Point at which to compute the derivative
//   - delta: Step size (h) for the finite difference
//
// Returns: Approximated derivative value f'(x)
func (f Central) derivadaPrimeira(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating central difference",
		slog.Float64("x", x),
		slog.Float64("delta", delta),
		slog.String("method", "central"),
		slog.String("accuracy", "O(h²)"))
	return (f.fn(x+delta) - f.fn(x-delta)) / (2 * delta)
}

// derivadaPrimeiraFactory creates a derivative calculation function based on the specified method.
//
// This factory function implements the Strategy design pattern, allowing the selection
// of different numerical differentiation methods at runtime. It provides a unified
// interface for computing first derivatives using various finite difference schemes.
//
// Parameters:
//   - ctx: Context for logging and potential cancellation
//   - fn: The function to differentiate (f: ℝ → ℝ)
//   - philosophyOption: Character specifying the differentiation method:
//   - 'f' or 'F': Forward difference (O(h) accuracy)
//   - 'b' or 'B': Backward difference (O(h) accuracy)
//   - 'c' or 'C': Central difference (O(h²) accuracy)
//   - default: Central difference (recommended for best accuracy)
//
// Returns: A function that computes f'(x) given x and step size delta
//
// Example usage:
//
//	fn := func(x float64) float64 { return x*x }
//	derivative := derivadaPrimeiraFactory(ctx, fn, 'c')
//	result := derivative(2.0, 0.001) // Computes f'(2.0) with h=0.001
func derivadaPrimeiraFactory(
	ctx context.Context,
	fn func(float64) float64,
	philosophyOption rune,
) func(float64, float64) float64 {
	var philosophy Derivacao
	var philosophyName string

	switch philosophyOption {
	case 'f', 'F':
		philosophy = &Forward{fn: fn}
		philosophyName = "Forward"
	case 'b', 'B':
		philosophy = &Backward{fn: fn}
		philosophyName = "Backward"
	case 'c', 'C':
		philosophy = &Central{fn: fn}
		philosophyName = "Central"
	default:
		philosophy = &Central{fn: fn}
		philosophyName = "Central (default)"
	}

	slog.InfoContext(ctx, "Selected derivative philosophy",
		slog.String("philosophy", philosophyName),
		slog.String("option", string(philosophyOption)))

	return func(x, delta float64) float64 {
		return philosophy.derivadaPrimeira(x, delta)
	}
}
