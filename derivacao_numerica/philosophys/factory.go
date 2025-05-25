package philosophys

import (
	"context"
	"log/slog"
)

// Derivacao defines the interface for numerical differentiation methods.
// All derivative calculation methods must implement this interface.
type Derivacao interface {
	// derivadaPrimeira computes the numerical derivative at point x using step size delta.
	// Returns the approximated derivative value.
	DerivadaPrimeira(x, delta float64) float64
	DerivadaSegunda(x, delta float64) float64
}

// DerivadaPrimeiraFactory creates a derivative calculation function based on the specified method.
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
//	derivative := DerivadaPrimeiraFactory(ctx, fn, 'c')
//	result := derivative(2.0, 0.001) // Computes f'(2.0) with h=0.001
func DerivadaPrimeiraFactory(
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
	)

	return func(x, delta float64) float64 {
		return philosophy.DerivadaPrimeira(x, delta)
	}
}

func DerivadaSegundaFactory(
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

	slog.InfoContext(ctx, "Selected second derivative philosophy",
		slog.String("philosophy", philosophyName),
	)

	return func(x, delta float64) float64 {
		return philosophy.DerivadaSegunda(x, delta)
	}
}
