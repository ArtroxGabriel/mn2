package powermethods

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func PotenciaInversa(
	a *mat.Dense,
	x0 *mat.VecDense,
	tolerance float64,
	maxIterations int,
) (*PowerMethodResult, error) {
	r, c := a.Dims()
	matrixInverse := mat.NewDense(r, c, nil)
	if err := matrixInverse.Inverse(a); err != nil {
		return nil, fmt.Errorf("could not compute inverse matrix: %v", err)
	}

	regularResult, err := PotenciaRegular(matrixInverse, x0, tolerance, maxIterations)
	if err != nil {
		return nil, fmt.Errorf("power iteration on inverse matrix failed: %v", err)
	}

	// ajust eigenValue to find the smallest eigenvalue of A
	eigenvalue := 1.0 / regularResult.Eigenvalue

	return &PowerMethodResult{
		Eigenvalue:  eigenvalue,
		Eigenvector: regularResult.Eigenvector,
	}, nil
}
