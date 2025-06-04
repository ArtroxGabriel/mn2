package fourthorder

type CentralFourthOrderStrategy struct{}

func (c *CentralFourthOrderStrategy) CalculateFirst(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+2*h) + 8*fn(x+h) - 8*fn(x-h) + fn(x-2*h)) / (12 * h)
}

func (c *CentralFourthOrderStrategy) CalculateSecond(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+2*h) + 16*fn(x+h) - 30*fn(x) + 16*fn(x-h) - fn(x-2*h)) / (12 * h * h)
}

func (c *CentralFourthOrderStrategy) CalculateThirty(
	fn func(float64) float64,
	x float64,
	h float64,
) float64 {
	return (-fn(x+3*h) + 8*fn(x+2*h) - 13*fn(x+h) + 13*fn(x-h) - 8*fn(x-2*h) + fn(x-3*h)) / (8 * h * h * h)
}
