package main

import (
	"fmt"
	householderqr "potencia/householder-qr"
	powermethods "potencia/power-methods"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{
		2, 3,
		5, 4,
	}
	n := 2
	tolerance := 1e-5
	A := mat.NewDense(n, n, data)

	fmt.Println("Matriz Original (A):")
	householderqr.PrintMatrix(A)

	fmt.Println("\n--- Método Householder QR ---")
	resultHouseholder := householderqr.HouseholderMethod(A)

	fmt.Println("\nMatriz Tridiagonal (T):")
	householderqr.PrintMatrix(resultHouseholder.T)

	fmt.Println("\nAplicando o método QR iterativo...")
	resultQR := householderqr.QRMethod(resultHouseholder.T, resultHouseholder.H, tolerance)

	fmt.Println("\nMatriz com Autovalores na Diagonal (Lambda):")
	householderqr.PrintMatrix(resultQR.Lambda)

	fmt.Println("\nMatriz de Autovetores (X):")
	householderqr.PrintMatrix(resultQR.X)

	fmt.Println("\nAutovalores (diagonal de Lambda):")
	diagonalView := resultQR.Lambda.DiagView()
	for i := 0; i < diagonalView.Diag(); i++ {
		fmt.Printf("λ%d = %.8f\n", i+1, diagonalView.At(i, i))
	}

	fmt.Println("\n--- Método da Potência Regular ---")
	x0 := mat.NewVecDense(n, nil)
	for i := range n {
		x0.SetVec(i, 1)
	}

	resultPower, err := powermethods.PotenciaRegular(A, x0, tolerance, 100)
	if err != nil {
		fmt.Println("Erro no método da potência regular:", err)
	} else {
		fmt.Printf("Autovalor dominante: %.8f\n", resultPower.Eigenvalue)
		fmt.Println("Autovetor correspondente:")
		householderqr.PrintMatrix(resultPower.Eigenvector)
	}
}
