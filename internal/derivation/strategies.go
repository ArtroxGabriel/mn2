package derivation

import (
	"context"
	"math"
)

type DerivationStrategy interface {
	CalculateFirst(
		ctx context.Context,
		fn func(float64) float64,
		x, h float64,
	) float64
	CalculateSecond(
		ctx context.Context,
		fn func(float64) float64,
		x, h float64,
	) float64
	CalculateThirty(
		ctx context.Context,
		fn func(float64) float64,
		x, h float64,
	) float64
}

var (
	_ DerivationStrategy = (*ForwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*ForwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*BackwardSecondOrderStrategy)(nil)
	_ DerivationStrategy = (*BackwardFirstOrderStrategy)(nil)
	_ DerivationStrategy = (*CentralSecondOrderStrategy)(nil)
)

type ForwardSecondOrderStrategy struct{}

func (f *ForwardSecondOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (-3*fn(x) + 4*fn(x+h) - fn(x+2*h)) / (2 * h)
}

func (f *ForwardSecondOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (2*fn(x) - 5*fn(x+h) + 4*fn(x+2*h) - 3*fn(x+3*h)) / (h * h)
}

func (f *ForwardSecondOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (-5*fn(x) + 18*fn(x+h) - 24*fn(x+2*h) + 14*fn(x+3*h) - 3*fn(x+4*h)) / (2 * math.Pow(h, 3))
}

type ForwardFirstOrderStrategy struct{}

func (f *ForwardFirstOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x+h) - fn(x)) / h
}

func (f *ForwardFirstOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x+2*h) - 2*fn(x+h) + fn(x)) / (h * h)
}

func (f *ForwardFirstOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x+3*h) - 3*fn(x+2*h) + 3*fn(x+2*h) - fn(x)) / (math.Pow(h, 3))
}

type BackwardSecondOrderStrategy struct{}

func (f *BackwardSecondOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (3*fn(x) - 4*fn(x-h) + fn(x-2*h)) / (2 * h)
}

func (f *BackwardSecondOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (2*fn(x) - 5*fn(x-h) + 4*fn(x-2*h) - 3*fn(x-3*h)) / (h * h)
}

func (b *BackwardSecondOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (5*fn(x) - 18*fn(x-h) + 24*fn(x-2*h) - 14*fn(x-3*h) + 3*fn(x-4*h)) / (2 * math.Pow(h, 3))
}

type BackwardFirstOrderStrategy struct{}

func (b *BackwardFirstOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x) - fn(x-h)) / h
}

func (b *BackwardFirstOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x) - 2*fn(x-h) + fn(x-2*h)) / (h * h)
}

func (b *BackwardFirstOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x) - 3*fn(x-h) + 3*fn(x-2*h) - fn(x-3*h)) / (math.Pow(h, 3))
}

type CentralSecondOrderStrategy struct{}

func (c *CentralSecondOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x+h) - fn(x-h)) / (2 * h)
}

func (c *CentralSecondOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x+h) - 2*fn(x) + fn(x-h)) / (h * h)
}

func (c *CentralSecondOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x-2*h) - 2*fn(x-h) + 2*fn(x+h) - fn(x+2*h)) / (2 * math.Pow(h, 3))
}
