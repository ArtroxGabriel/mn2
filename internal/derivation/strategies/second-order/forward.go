package secondorder

import "context"

type ForwardSecondOrderStrategy struct{}

func (f *ForwardSecondOrderStrategy) CalculateFirst(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (-3*fn(x) + 4*fn(x+h) - fn(x+2*h)) / (2 * h)
}

func (f *ForwardSecondOrderStrategy) CalculateSecond(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (2*fn(x) - 5*fn(x+h) + 4*fn(x+2*h) - 3*fn(x+3*h)) / (h * h)
}

func (f *ForwardSecondOrderStrategy) CalculateThirty(ctx context.Context, fn func(float64) float64, x float64, h float64) float64 {
	return (-5*fn(x) + 18*fn(x+h) - 24*fn(x+2*h) + 14*fn(x+3*h) - 3*fn(x+4*h)) / (2 * h * h * h)
}
