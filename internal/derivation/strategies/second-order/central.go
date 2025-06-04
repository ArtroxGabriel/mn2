package secondorder

type CentralSecondOrderStrategy struct{}

func (c *CentralSecondOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x+h) - fn(x-h)) / (2 * h)
}

func (c *CentralSecondOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x+h) - 2*fn(x) + fn(x-h)) / (h * h)
}

func (c *CentralSecondOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (fn(x-2*h) - 2*fn(x-h) + 2*fn(x+h) - fn(x+2*h)) / (2 * h * h * h)
}
