package thirdorder

type ForwardThirOrderStrategy struct{}

func (f *ForwardThirOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 6*fn(x+2*h) - 18*fn(x+h) - 10*fn(x)) / (12 * h)
}

func (f *ForwardThirOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 4*fn(x+2*h) - 5*fn(x+h) + 2*fn(x)) / (h * h)
}

func (f *ForwardThirOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-3*fn(x+4*h) + 14*fn(x+3*h) - 24*fn(x+2*h) + 18*fn(x+h) - 5*fn(x)) / (2 * h * h * h)
}
