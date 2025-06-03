package derivation

import (
	"context"
	"log/slog"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	firstorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/first-order"
	secondorder "github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation/strategies/second-order"
)

func DerivacaoFactory(ctx context.Context, strategy string, order int) (DerivationStrategy, error) {
	slog.DebugContext(ctx, "Creating derivation strategy",
		slog.String("strategy", strategy),
		slog.Int("order", order),
	)

	var derivation DerivationStrategy
	switch {
	case strategy == "forward" && order == 1:
		derivation = &firstorder.ForwardFirstOrderStrategy{}
	case strategy == "forward" && order == 2:
		derivation = &secondorder.ForwardSecondOrderStrategy{}
	case strategy == "backward" && order == 1:
		derivation = &firstorder.BackwardFirstOrderStrategy{}
	case strategy == "backward" && order == 2:
		derivation = &secondorder.BackwardSecondOrderStrategy{}
	case strategy == "central" && order == 2:
		derivation = &secondorder.CentralSecondOrderStrategy{}
	default:
		slog.ErrorContext(ctx, "Invalid derivation strategy")
		return nil, common.ErrInvalidStrategy
	}

	slog.InfoContext(ctx, "Derivation created succesfully")
	return derivation, nil
}
