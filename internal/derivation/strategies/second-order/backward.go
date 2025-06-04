package secondorder

type BackwardSecondOrderStrategy struct{}

func (f *BackwardSecondOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (3*fn(x) - 4*fn(x-h) + fn(x-2*h)) / (2 * h)
}

func (f *BackwardSecondOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - 2*fn(x-h) + fn(x-2*h)) / (h * h)
}

func (b *BackwardSecondOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x) - 3*fn(x-h) + 3*fn(x-2*h) - fn(x-3*h)) / (h * h * h)
}
