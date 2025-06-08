package common

import "fmt"

// State represents the current view/context of the application.
type State uint64 // Changed to uint64 to match existing type

const (
	StateMainMenu State = iota
	StateDerivationMenu
	StateSelectPhilosophy
	StateSelectErrorOrder
	StateSelectDerivativeOrder
	StateSelectFunction
	StateResult
	// States for Integration
	StateIntegrationMenu         // Was placeholder, now formally part of sequence
	StateSelectIntegrationMethod
	StateSelectIntegrationDegree // For selecting N (degree/points)
)

// Focus indicates which input field is currently active.
// Using iota for Focus as well for consistency and ease of adding new ones.
type Focus int

const (
	FocusNone Focus = iota // Starts at 0
	FocusX                 // 1
	FocusH                 // 2
	// New focus states for Integration
	FocusLowerLimit        // 3 - Input for 'a'
	FocusUpperLimit        // 4 - Input for 'b'
	FocusTolerance         // 5 - Input for 'tol'
	FocusNPoints           // 6 - Input for 'n' (degree/points for integration)
)

// FunctionDefinition stores a predefined function's details.
type FunctionDefinition struct {
	ID   string
	Name string
	// Func func(float64) float64 // Changed from MathFunc to actual signature for clarity
	Func func(float64) float64
}

// Global errors
var (
	ErrInvalidStrategy = fmt.Errorf("estratégia de derivação inválida ou não suportada para a ordem de erro especificada")
	ErrInvalidDerivate = fmt.Errorf("ordem de derivada não suportada pela estratégia atual")
	ErrZeroValue       = fmt.Errorf("o valor de h (passo) ou divisor não pode ser zero")
)
