package main

import (
	"context"
	"log/slog"
)

type Derivacao interface {
	calculate(float64, float64) float64
}

// Forward implements the forward difference method for numerical differentiation.
// This method approximates the first derivative and has a truncation error of order O(delta).
type Forward struct {
	fn func(float64) float64
}

func (f Forward) calculate(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating forward difference", slog.Float64("x", x), slog.Float64("delta", delta))
	return (f.fn(x+delta) - f.fn(x)) / delta
}

// Backward implements the backward difference method for numerical differentiation.
// This method approximates the first derivative and has a truncation error of order O(delta).
type Backward struct {
	fn func(float64) float64
}

func (f Backward) calculate(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating backward difference", slog.Float64("x", x), slog.Float64("delta", delta))
	return (f.fn(x) - f.fn(x-delta)) / delta
}

// Central implements the central difference method for numerical differentiation.
// This method approximates the first derivative and has a truncation error of order O(delta^2).
type Central struct {
	fn func(float64) float64
}

func (f Central) calculate(x, delta float64) float64 {
	slog.DebugContext(context.Background(), "Calculating central difference", slog.Float64("x", x), slog.Float64("delta", delta))
	return (f.fn(x+delta) - f.fn(x-delta)) / (2 * delta)
}

func derivadaPrimeiraFactory(
	ctx context.Context,
	fn func(float64) float64,
	philosophyOption rune,
) func(float64, float64) float64 {
	var philosophy Derivacao
	var philosophyName string

	switch philosophyOption {
	case 'f':
		philosophy = &Forward{fn: fn}
		philosophyName = "Forward"
	case 'b':
		philosophy = &Backward{fn: fn}
		philosophyName = "Backward"
	default:
		philosophy = &Central{fn: fn}
		philosophyName = "Central"
	}
	slog.InfoContext(ctx, "Selected derivative philosophy", slog.String("philosophy", philosophyName))

	return func(x, delta float64) float64 {
		return philosophy.calculate(x, delta)
	}
}
