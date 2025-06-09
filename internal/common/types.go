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
	StateIntegrationMenu
	StateSelectIntegrationMethod
	StateSelectIntegrationLimitA // May not be needed if direct input is used
	StateSelectIntegrationLimitB // May not be needed if direct input is used
	StateSelectIntegrationN      // May not be needed if direct input is used
	// Adicione outros estados conforme necessário
)

type FunctionDefinition struct {
	Name string
	Func MathFunc
	ID   string // Um ID único, se necessário para GetPredefinedFunc
}

// Constantes para foco em campos de input
const (
	FocusNone int = iota
	FocusX
	FocusH
	FocusIntegrationA
	FocusIntegrationB
	FocusIntegrationN
)
