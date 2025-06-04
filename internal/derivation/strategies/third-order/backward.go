package thirdorder

type BackwardThirOrderStrategy struct{}

func (b *BackwardThirOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (10*fn(x) - 18*fn(x-h) + 6*fn(x-2*h) - fn(x-3*h)) / (12 * h)
}

func (b *BackwardThirOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (2*fn(x) - 5*fn(x-h) + 4*fn(x-2*h) - fn(x-3*h)) / (h * h)
}

func (b *BackwardThirOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (5*fn(x) - 18*fn(x-h) + 24*fn(x-2*h) - 14*fn(x-3*h) + 3*fn(x-4*h)) / (2 * h * h * h)
}
