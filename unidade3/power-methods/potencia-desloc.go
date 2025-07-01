package powermethods

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func PotenciaDeslocamento(
	A *mat.Dense,
	v0 *mat.VecDense,
	tolerance float64,
	mu float64,
	maxIterations int,
) (*PowerMethodResult, error) {

	r, c := A.Dims()
	if r == 0 || r != c {
		return nil, fmt.Errorf("the matrix A must be square and non-empty")
	}

	A_hat := mat.NewDense(r, c, nil)
	A_hat.Copy(A)

	for i := range r {
		currentVal := A_hat.At(i, i)
		A_hat.Set(i, i, currentVal-mu)
	}

	inverseResult, invErr := PotenciaInversa(A_hat, v0, tolerance, maxIterations)
	if invErr != nil {
		return nil, fmt.Errorf("the shifted matrix A_hat failed in the inverse power method: %v", invErr)
	}

	// ajust eigenvalue
	lambda_i := inverseResult.Eigenvalue + mu

	x_i := inverseResult.Eigenvector

	return &PowerMethodResult{
		Eigenvalue:  lambda_i,
		Eigenvector: x_i,
	}, nil

}
