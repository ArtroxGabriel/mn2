package powermethods

import "gonum.org/v1/gonum/mat"

type PowerMethodResult struct {
	Eigenvalue  float64
	Eigenvector *mat.VecDense
}
