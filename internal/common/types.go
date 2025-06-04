package common

type MathFunc func(float64) float64

type State uint64

const (
	StateMainMenu State = iota
	StateDerivationMenu
	StateSelectPhilosophy
	StateSelectErrorOrder
	StateSelectDerivativeOrder
	StateSelectFunction
	StateResult
	StateIntegrationMenu // Mantendo como placeholder
	// Adicione outros estados conforme necessário
)

type FunctionDefinition struct {
	Name string
	Func MathFunc
	ID   string // Um ID único, se necessário para GetPredefinedFunc
}

// Constantes para foco em campos de input
const (
	FocusNone = 0
	FocusX    = 1
	FocusH    = 2
)
