package householderqr

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type HouseholderResult struct {
	T *mat.Dense
	H *mat.Dense
}

func NewIdentityMatrix(r, c int) *mat.Dense {
	h := mat.NewDense(r, c, nil)
	for i := range r {
		h.Set(i, i, 1.0)
	}
	return h
}

func householderMatrix(A *mat.Dense, i int) *mat.Dense {
	n, _ := A.Dims()
	I := NewIdentityMatrix(n, n)

	w := mat.NewVecDense(n, nil)
	w_linha := mat.NewVecDense(n, nil)

	col := A.ColView(i)
	for j := i + 1; j < n; j++ {
		w.SetVec(j, col.AtVec(j))
	}

	Lw := mat.Norm(w, 2)

	if i+1 < n {
		w_linha.SetVec(i+1, Lw)
	}

	N := mat.NewVecDense(n, nil)
	N.SubVec(w, w_linha)

	normN := mat.Norm(N, 2)
	n_vec := mat.NewVecDense(n, nil)
	if normN != 0 {
		n_vec.ScaleVec(1/normN, N)
	}

	n_vec_t := n_vec.T()
	outerProd := mat.NewDense(n, n, nil)
	outerProd.Mul(n_vec, n_vec_t)

	hMatrix := mat.NewDense(n, n, nil)
	hMatrix.Scale(-2, outerProd)
	hMatrix.Add(I, hMatrix)

	return hMatrix
}

func HouseholderMethod(A *mat.Dense) HouseholderResult {
	n, _ := A.Dims()
	H := NewIdentityMatrix(n, n)

	A_old := mat.DenseCopyOf(A)
	A_new := mat.NewDense(n, n, nil)

	for i := 0; i < n-2; i++ {
		Hi := householderMatrix(A_old, i)

		temp := mat.NewDense(n, n, nil)
		temp.Mul(A_old, Hi)
		A_new.Mul(Hi.T(), temp)

		A_old.Copy(A_new)

		H.Mul(H, Hi)
	}

	return HouseholderResult{T: A_old, H: H}
}

func PrintMatrix(m mat.Matrix) {
	fa := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
