package powermethods

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func PotenciaRegular(
	a *mat.Dense,
	x0 *mat.VecDense,
	tolerance float64,
	maxIterations int,
) (*PowerMethodResult, error) {
	r, c := a.Dims()

	if r != c {
		return nil, fmt.Errorf("matrix must be square")
	}

	if r == 0 {
		return nil, fmt.Errorf("matrix cannot be empty")
	}

	if c != x0.Len() {
		return nil, fmt.Errorf("The dimensions of the matrix and the vector are incompatible")
	}

	x := mat.NewVecDense(x0.Len(), nil)
	x.CopyVec(x0)

	y := mat.NewVecDense(x0.Len(), nil)

	var oldEigenvalue float64
	var eigenvalue float64

	for i := range maxIterations {
		y.MulVec(a, x)

		eigenvalue = mat.Norm(y, math.Inf(0))

		if eigenvalue == 0 {
			return nil, fmt.Errorf("the eigenvalue is zero, it's not possible to normalize the eigenvector")
		}

		x.ScaleVec(1/eigenvalue, y)

		if i > 0 && math.Abs((eigenvalue-oldEigenvalue)/eigenvalue) < tolerance {
			return &PowerMethodResult{
				Eigenvalue:  eigenvalue,
				Eigenvector: x,
			}, nil
		}

		oldEigenvalue = eigenvalue
	}

	return nil, fmt.Errorf("power iteration did not converge within %d iterations", maxIterations)
}
