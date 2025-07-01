package householderqr

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type QRResult struct {
	Lambda *mat.Dense
	X      *mat.Dense
}

func QRDecomp(A *mat.Dense) (*mat.Dense, *mat.Dense) {
	n, m := A.Dims()
	R := mat.DenseCopyOf(A)

	Q := NewIdentityMatrix(n, m)

	for j := range m {
		for i := n - 1; i > j; i-- {
			a := R.At(i-1, j)
			b := R.At(i, j)

			r := math.Sqrt(a*a + b*b)
			if r < 1e-15 {
				continue
			}
			cos := a / r
			sen := -b / r

			// Given's rotation matrix
			G := NewIdentityMatrix(n, n)
			G.Set(i-1, i-1, cos)
			G.Set(i, i, cos)
			G.Set(i-1, i, -sen)
			G.Set(i, i-1, sen)

			// R = G * R e Q = Q * G^T
			tempR := mat.DenseCopyOf(R)
			R.Mul(G, tempR)

			tempQ := mat.DenseCopyOf(Q)
			Q.Mul(tempQ, G.T())
		}
	}

	return Q, R
}

func QRMethod(T, H *mat.Dense, epsilon float64) QRResult {
	X := mat.DenseCopyOf(H)
	A := mat.DenseCopyOf(T)
	n, _ := A.Dims()

	error := 1.0

	for error > epsilon {
		// DecompQR of current A
		Q, R := QRDecomp(A)

		// A_k+1 = R_k * Q_k
		A.Mul(R, Q)

		// Acc the orthogonal transformations to obtain the eigenvectors.
		X.Mul(X, Q)

		error = 0.0
		for j := 0; j < n-1; j++ {
			error += math.Abs(A.At(j+1, j))
		}
	}

	return QRResult{Lambda: A, X: X}
}
