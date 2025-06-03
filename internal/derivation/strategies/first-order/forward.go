package firstorder

import "context"

type ForwardFirstOrderStrategy struct{}

func (f *ForwardFirstOrderStrategy) CalculateFirst(
	ctx context.Context,
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+h) - fn(x)) / h
}

func (f *ForwardFirstOrderStrategy) CalculateSecond(
	ctx context.Context,
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+2*h) - 2*fn(x+h) + fn(x)) / (h * h)
}

func (f *ForwardFirstOrderStrategy) CalculateThirty(
	ctx context.Context,
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+3*h) - 3*fn(x+2*h) + 3*fn(x+h) - fn(x)) / (h * h * h)
}
