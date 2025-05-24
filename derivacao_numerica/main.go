package main

import (
	"context"
	"log/slog"
	"math"
	"os"
)

type Derivada interface {
	calculate(float64, float64) float64
}

var fn func(float64) float64

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	ctx := context.Background()

	fn = func(x float64) float64 {
		slog.DebugContext(ctx, "Calculando f(x)", slog.Float64("x", x))
		return math.Sqrt(math.Exp(3*x) + 4*x*x)
	}

	x := 2.0
	dx := 0.5
	limiteErro := 1e-6

	derivadaPrimeiraFn := derivadaPrimeiraFactory(
		ctx,
		fn,
		'f',
	)

	resultadoAtual := derivadaPrimeiraFn(x, dx)
	slog.InfoContext(ctx, "Initial calculation", slog.Float64("dx", dx), slog.Float64("f'(x)", resultadoAtual))

	var resultadoAnterior float64

	for i := 0; ; i++ {
		slog.DebugContext(ctx, "Iteration start", slog.Int("iteration", i), slog.Float64("current_dx", dx), slog.Float64("current_f'(x)", resultadoAtual))
		resultadoAnterior = resultadoAtual

		dx /= 2

		resultadoAtual = derivadaPrimeiraFn(x, dx)

		var erroRelativo float64
		if math.Abs(resultadoAtual) < 1e-15 {
			erroRelativo = math.Inf(1)
			if math.Abs(resultadoAtual-resultadoAnterior) < 1e-15 {
				erroRelativo = 0
			}
		} else {
			erroRelativo = math.Abs((resultadoAtual - resultadoAnterior) / resultadoAtual)
		}
		slog.InfoContext(ctx, "Iteration result", slog.Float64("dx", dx), slog.Float64("f'(x)", resultadoAtual), slog.Float64("error", erroRelativo))

		if erroRelativo < limiteErro {
			slog.InfoContext(ctx, "Convergence reached")
			break
		}

		if dx < 1e-10 {
			slog.WarnContext(ctx, "dx too small, potential convergence or precision issue")
			break
		}
	}

	slog.InfoContext(ctx, "Final result", slog.Float64("final_dx", dx), slog.Float64("final_f'(x)", resultadoAtual))
}
