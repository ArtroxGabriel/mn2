package firstorder

type ForwardFirstOrderStrategy struct{}

func (f *ForwardFirstOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+h) - fn(x)) / h
}

func (f *ForwardFirstOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+2*h) - 2*fn(x+h) + fn(x)) / (h * h)
}

func (f *ForwardFirstOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x, h float64,
) float64 {
	return (fn(x+3*h) - 3*fn(x+2*h) + 3*fn(x+h) - fn(x)) / (h * h * h)
}
