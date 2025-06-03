package firstorder

import "context"

type BackwardFirstOrderStrategy struct{}

func (b *BackwardFirstOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x) - fn(x-h)) / h
}

func (b *BackwardFirstOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x) - 2*fn(x-h) + fn(x-2*h)) / (h * h)
}

func (b *BackwardFirstOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (fn(x) - 3*fn(x-h) + 3*fn(x-2*h) - fn(x-3*h)) / (h * h * h)
}
