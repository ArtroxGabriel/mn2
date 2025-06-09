package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/functions"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration" // Added import

	tea "github.com/charmbracelet/bubbletea"
)

var integrationMethodDisplayToStrategyName = map[string]string{
	"Regra do Trapézio (Newton-Cotes O1)": "NewtonCotesOrder1",
	"Simpson 1/3 (Newton-Cotes O2)":     "NewtonCotesOrder2",
	"Simpson 3/8 (Newton-Cotes O3)":     "NewtonCotesOrder3",
	"Regra de Boole (Newton-Cotes O4)":  "NewtonCotesOrder4",
	"Gauss-Legendre O1":                 "GaussLegendreOrder1",
	"Gauss-Legendre O2":                 "GaussLegendreOrder2",
	"Gauss-Legendre O3":                 "GaussLegendreOrder3",
	"Gauss-Legendre O4":                 "GaussLegendreOrder4",
}

var _ tea.Model = (*MainModel)(nil)

type ErrorOrderOption struct {
	Display string
	Value   uint
}

type DerivativeOrderOption struct {
	Display string
	Value   int
}

var allPossibleErrorOrderOptions = []ErrorOrderOption{
	{Display: "Ordem 1 (O(h))", Value: 1},
	{Display: "Ordem 2 (O(h^2))", Value: 2},
	{Display: "Ordem 3 (O(h^3))", Value: 3},
	{Display: "Ordem 4 (O(h^4))", Value: 4},
}

type MainModel struct {
	state           common.State
	mainMenuChoices []string
	cursor          int

	derivationMenuChoices        []string
	derivationCursor             int
	selectedDerivationPhilosophy string
	selectedDerivationErrorOrder uint
	selectedDerivationOrder      int
	selectedFunctionDef          common.FunctionDefinition
	currentX                     string
	currentH                     string
	result                       string
	err                          error

	focus int

	philosophyOptions        []string
	currentErrorOrderOptions []ErrorOrderOption
	derivativeOrderOptions   []DerivativeOrderOption
	functionDefinitions      []common.FunctionDefinition
	selectionCursor          int

	// Integration fields
	integrationMenuChoices        []string
	integrationCursor             int
	selectedIntegrationMethod     string
	availableIntegrationMethods   []string
	currentA                      string
	currentB                      string
	currentN                      string
	selectedIntegrationFunctionDef common.FunctionDefinition
	integrationFocus              int

	previousState common.State // Added to store the state before navigating to a sub-menu like function selection
}

func NewMainModel() *MainModel {
	philosophyOpts := []string{"Forward", "Backward", "Central"}
	derivativeOrderOpts := []DerivativeOrderOption{
		{Display: "Primeira Derivada", Value: 1},
		{Display: "Segunda Derivada", Value: 2},
		{Display: "Terceira Derivada", Value: 3},
	}
	funcDefs := functions.GetFunctionDefinitions()

	m := &MainModel{
		state:                 common.StateMainMenu,
		mainMenuChoices:       []string{"Derivação Numérica", "Integração Numérica", "Sair"},
		derivationMenuChoices: []string{"Filosofia", "Ordem do Erro", "Ordem da Derivada", "Função", "Ponto x", "Passo h", "Calcular", "Voltar"},

		philosophyOptions:      philosophyOpts,
		derivativeOrderOptions: derivativeOrderOpts,
		functionDefinitions:    funcDefs,

		selectedDerivationPhilosophy: philosophyOpts[0],
		selectedDerivationOrder:      derivativeOrderOpts[0].Value,
		currentX:                     "1.0",
		currentH:                     "0.1",
	}

	if len(funcDefs) > 0 {
		m.selectedFunctionDef = funcDefs[0]
	}

	m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy)

	// Initialize integration fields
	m.integrationMenuChoices = []string{"Método", "Função", "Limite Inferior (a)", "Limite Superior (b)", "Num de Subintervalos/Ordem (n)", "Calcular", "Voltar"}
	m.availableIntegrationMethods = []string{
		"Regra do Trapézio (Newton-Cotes O1)", // Maps to NewtonCotesOrder1
		"Simpson 1/3 (Newton-Cotes O2)",     // Maps to NewtonCotesOrder2
		"Simpson 3/8 (Newton-Cotes O3)",     // Maps to NewtonCotesOrder3
		"Regra de Boole (Newton-Cotes O4)",  // Maps to NewtonCotesOrder4
		"Gauss-Legendre O1",                 // Maps to GaussLegendreOrder1
		"Gauss-Legendre O2",                 // Maps to GaussLegendreOrder2
		"Gauss-Legendre O3",                 // Maps to GaussLegendreOrder3
		"Gauss-Legendre O4",                 // Maps to GaussLegendreOrder4
	}
	if len(m.availableIntegrationMethods) > 0 {
		m.selectedIntegrationMethod = m.availableIntegrationMethods[0] // Default selection
	}
	if len(funcDefs) > 0 { // Reuse function definitions from derivation for now
		m.selectedIntegrationFunctionDef = funcDefs[0]
	}
	m.currentA = "0.0"
	m.currentB = "1.0"
	m.currentN = "10"
	m.integrationFocus = common.FocusNone // Assuming common.FocusNone is 0 or defined

	return m
}

func (m *MainModel) updateAvailableErrorOrders(philosophy string) {
	var availableOrders []ErrorOrderOption
	switch philosophy {
	case "Central":
		for _, opt := range allPossibleErrorOrderOptions {
			if opt.Value == 2 || opt.Value == 4 {
				availableOrders = append(availableOrders, opt)
			}
		}
	case "Forward", "Backward":
		for _, opt := range allPossibleErrorOrderOptions {
			if opt.Value == 1 || opt.Value == 2 || opt.Value == 3 {
				availableOrders = append(availableOrders, opt)
			}
		}
	default:
		availableOrders = allPossibleErrorOrderOptions
	}
	m.currentErrorOrderOptions = availableOrders

	isValidSelection := false
	for _, opt := range m.currentErrorOrderOptions {
		if opt.Value == m.selectedDerivationErrorOrder {
			isValidSelection = true
			break
		}
	}

	if !isValidSelection && len(m.currentErrorOrderOptions) > 0 {
		m.selectedDerivationErrorOrder = m.currentErrorOrderOptions[0].Value
	} else if len(m.currentErrorOrderOptions) == 0 {
		m.selectedDerivationErrorOrder = 0
	}
}

func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.state == common.StateDerivationMenu && m.focus != common.FocusNone {
			switch msg.String() {
			case "enter":
				m.focus = common.FocusNone
				return m, nil
			case "backspace":
				if m.focus == common.FocusX && len(m.currentX) > 0 {
					m.currentX = m.currentX[:len(m.currentX)-1]
				} else if m.focus == common.FocusH && len(m.currentH) > 0 {
					m.currentH = m.currentH[:len(m.currentH)-1]
				}
				return m, nil
			default:
				char := msg.String()
				currentStr := ""
				if m.focus == common.FocusX {
					currentStr = m.currentX
				} else if m.focus == common.FocusH {
					currentStr = m.currentH
				}

				if (char >= "0" && char <= "9") || (char == "." && !strings.Contains(currentStr, ".")) || (char == "-" && len(currentStr) == 0) {
					if m.focus == common.FocusX {
						m.currentX += char
					} else if m.focus == common.FocusH {
						m.currentH += char
					}
				}
				return m, nil
			}
		}

		// Handle input for integration fields (A, B, N)
		if m.state == common.StateIntegrationMenu && m.integrationFocus != common.FocusNone {
			switch msg.String() {
			case "enter":
				m.integrationFocus = common.FocusNone
				return m, nil
			case "backspace":
				if m.integrationFocus == common.FocusIntegrationA && len(m.currentA) > 0 {
					m.currentA = m.currentA[:len(m.currentA)-1]
				} else if m.integrationFocus == common.FocusIntegrationB && len(m.currentB) > 0 {
					m.currentB = m.currentB[:len(m.currentB)-1]
				} else if m.integrationFocus == common.FocusIntegrationN && len(m.currentN) > 0 {
					m.currentN = m.currentN[:len(m.currentN)-1]
				}
				return m, nil
			default:
				char := msg.String()
				var targetStr *string
				isNField := false

				if m.integrationFocus == common.FocusIntegrationA {
					targetStr = &m.currentA
				} else if m.integrationFocus == common.FocusIntegrationB {
					targetStr = &m.currentB
				} else if m.integrationFocus == common.FocusIntegrationN {
					targetStr = &m.currentN
					isNField = true
				}

				if targetStr != nil {
					if isNField { // N field (subintervals/order) should only accept positive integers
						if char >= "0" && char <= "9" {
							// Prevent leading zeros unless it's the only digit
							if *targetStr == "0" && char != "0" {
								*targetStr = char
							} else if *targetStr != "0" || char != "0" { // Allow '0' if it's not already '0'
								*targetStr += char
							}
						}
					} else { // A and B fields (limits) can be float
						if (char >= "0" && char <= "9") || (char == "." && !strings.Contains(*targetStr, ".")) || (char == "-" && len(*targetStr) == 0) {
							*targetStr += char
						}
					}
				}
				return m, nil
			}
		}

		switch m.state {
		case common.StateMainMenu:
			return m.updateMainMenu(msg)
		case common.StateDerivationMenu:
			return m.updateDerivationMenu(msg)
		case common.StateSelectPhilosophy:
			return m.updateSelectPhilosophy(msg)
		case common.StateSelectErrorOrder:
			return m.updateSelectErrorOrder(msg)
		case common.StateSelectDerivativeOrder:
			return m.updateSelectDerivativeOrder(msg)
		case common.StateSelectFunction:
			return m.updateSelectFunction(msg)
		case common.StateResult:
			return m.updateResultScreen(msg)
		case common.StateIntegrationMenu:
			return m.updateIntegrationMenu(msg)
		// Note: StateSelectIntegrationMethod will need its own update function if it's a separate screen
		// For now, assuming selection happens and returns to StateIntegrationMenu or handled by a generic selector
		}
	}
	return m, nil
}

func (m *MainModel) updateMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.mainMenuChoices)-1 {
			m.cursor++
		}
	case "enter":
		selectedItem := m.mainMenuChoices[m.cursor]
		switch selectedItem {
		case "Derivação Numérica":
			m.state = common.StateDerivationMenu
			m.derivationCursor = 0
			m.err = nil
			m.result = ""
		case "Integração Numérica":
			m.state = common.StateIntegrationMenu
			m.resetIntegrationInputs() // Call the reset method
			m.err = nil
		case "Sair":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MainModel) updateDerivationMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q":
		m.state = common.StateMainMenu
		m.cursor = 0
		m.resetDerivationInputs()
		return m, nil
	case "up", "k":
		if m.derivationCursor > 0 {
			m.derivationCursor--
		}
	case "down", "j":
		if m.derivationCursor < len(m.derivationMenuChoices)-1 {
			m.derivationCursor++
		}
	case "enter":
		choice := m.derivationMenuChoices[m.derivationCursor]
		switch choice {
		case "Filosofia":
			m.state = common.StateSelectPhilosophy
			m.selectionCursor = 0
			for i, phil := range m.philosophyOptions {
				if phil == m.selectedDerivationPhilosophy {
					m.selectionCursor = i
					break
				}
			}
		case "Ordem do Erro":
			m.state = common.StateSelectErrorOrder
			m.selectionCursor = 0
			for i, eo := range m.currentErrorOrderOptions {
				if eo.Value == m.selectedDerivationErrorOrder {
					m.selectionCursor = i
					break
				}
			}
		case "Ordem da Derivada":
			m.state = common.StateSelectDerivativeOrder
			m.selectionCursor = 0
			for i, do := range m.derivativeOrderOptions {
				if do.Value == m.selectedDerivationOrder {
					m.selectionCursor = i
					break
				}
			}
		case "Função":
			m.previousState = m.state // Store current state (StateDerivationMenu)
			m.state = common.StateSelectFunction
			m.selectionCursor = 0
			// Set selection cursor to currently selected derivation function
			for i, fd := range m.functionDefinitions {
				if fd.ID == m.selectedFunctionDef.ID {
					m.selectionCursor = i
					break
				}
			}
		case "Ponto x":
			m.focus = common.FocusX
		case "Passo h":
			m.focus = common.FocusH
		case "Calcular":
			m.performDerivation()
			m.state = common.StateResult
		case "Voltar":
			m.state = common.StateMainMenu
			m.cursor = 0
			m.resetDerivationInputs()
		}
	}
	return m, nil
}

func (m *MainModel) updateSelectPhilosophy(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		if m.previousState != 0 {
			m.state = m.previousState
		} else {
			// Fallback if previousState somehow not set, though it should be.
			// Defaulting to StateDerivationMenu might be contextually wrong if called from integration.
			// However, StateSelectFunction was originally only for derivation.
			// This logic implies StateSelectFunction should ideally know its caller without previousState if it were more isolated.
			m.state = common.StateDerivationMenu // Or StateMainMenu for a more generic fallback
		}
		m.previousState = 0 // Reset previousState
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {
			m.selectionCursor--
		}
	case "down", "j":
		if m.selectionCursor < len(m.philosophyOptions)-1 {
			m.selectionCursor++
		}
	case "enter":
		selectedPhilosophy := m.philosophyOptions[m.selectionCursor]
		if m.selectedDerivationPhilosophy != selectedPhilosophy {
			m.selectedDerivationPhilosophy = selectedPhilosophy
			m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy)
		}
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

func (m *MainModel) updateSelectErrorOrder(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if len(m.currentErrorOrderOptions) == 0 {
		if msg.String() == "enter" || msg.String() == "q" || msg.String() == "ctrl+c" {
			m.state = common.StateDerivationMenu
			return m, nil
		}
		return m, nil
	}

	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateDerivationMenu
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {
			m.selectionCursor--
		}
	case "down", "j":
		if m.selectionCursor < len(m.currentErrorOrderOptions)-1 {
			m.selectionCursor++
		}
	case "enter":
		if m.selectionCursor >= 0 && m.selectionCursor < len(m.currentErrorOrderOptions) {
			m.selectedDerivationErrorOrder = m.currentErrorOrderOptions[m.selectionCursor].Value
		}
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

func (m *MainModel) updateSelectDerivativeOrder(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateDerivationMenu
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {
			m.selectionCursor--
		}
	case "down", "j":
		if m.selectionCursor < len(m.derivativeOrderOptions)-1 {
			m.selectionCursor++
		}
	case "enter":
		m.selectedDerivationOrder = m.derivativeOrderOptions[m.selectionCursor].Value
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

func (m *MainModel) updateSelectFunction(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateDerivationMenu
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {
			m.selectionCursor--
		}
	case "down", "j":
		if m.selectionCursor < len(m.functionDefinitions)-1 {
			m.selectionCursor++
		}
	case "enter":
		if m.selectionCursor < 0 || m.selectionCursor >= len(m.functionDefinitions) {
			// Should not happen if navigation is correct, but as a safeguard:
			if m.previousState != 0 {
				m.state = m.previousState
			} else {
				m.state = common.StateMainMenu // Fallback
			}
			m.previousState = 0
			return m, nil
		}

		selectedFunc := m.functionDefinitions[m.selectionCursor]

		if m.previousState == common.StateIntegrationMenu {
			m.selectedIntegrationFunctionDef = selectedFunc
		} else {
			// Default to derivation or handle other previous states if any
			m.selectedFunctionDef = selectedFunc
		}

		if m.previousState != 0 { // Check if previousState was set
			m.state = m.previousState
		} else {
			m.state = common.StateMainMenu // Fallback, though previousState should always be set
		}
		m.previousState = 0 // Reset previousState
	}
	return m, nil
}

func (m *MainModel) updateResultScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "enter":
		m.state = common.StateMainMenu
		if m.err != nil && (strings.Contains(m.err.Error(), "derivador") || strings.Contains(m.err.Error(), "combinação inválida")) {
			m.state = common.StateDerivationMenu
		}
		m.result = ""
		m.err = nil
		m.cursor = 0
	}
	return m, nil
}

func (m *MainModel) View() string {
	return View(m)
}

func (m *MainModel) performDerivation() {
	m.err = nil
	m.result = ""

	x, errX := strconv.ParseFloat(strings.TrimSpace(m.currentX), 64)
	if errX != nil {
		m.err = fmt.Errorf("valor inválido para Ponto x '%s': %w. Use '.' como separador decimal e opcionalmente '-' no início.", m.currentX, errX)
		return
	}
	h, errH := strconv.ParseFloat(strings.TrimSpace(m.currentH), 64)
	if errH != nil {
		m.err = fmt.Errorf("valor inválido para Passo h '%s': %w. Use '.' como separador decimal.", m.currentH, errH)
		return
	}

	if h == 0 {
		m.err = common.ErrZeroValue
		return
	}

	if m.selectedFunctionDef.Func == nil {
		m.err = fmt.Errorf("nenhuma função selecionada")
		return
	}

	_, factoryErrTest := derivation.DerivacaoFactory(m.selectedDerivationPhilosophy, m.selectedDerivationErrorOrder)
	if factoryErrTest != nil {
		m.err = fmt.Errorf("erro interno ou combinação inválida não tratada: filosofia '%s', ordem de erro O(h^%d). Detalhe: %w", m.selectedDerivationPhilosophy, m.selectedDerivationErrorOrder, factoryErrTest)
		return
	}

	derivator, err := derivation.NewDerivator(
		m.selectedDerivationPhilosophy,
		m.selectedDerivationErrorOrder,
	)
	if err != nil {
		m.err = fmt.Errorf("falha ao criar derivador: %w", err)
		return
	}

	calculatedVal, calcErr := derivator.Calculate(m.selectedFunctionDef.Func, x, h, m.selectedDerivationOrder)
	if calcErr != nil {
		if calcErr == common.ErrInvalidDerivate {
			m.err = fmt.Errorf("a ordem da derivada selecionada (%da) não é suportada pela estratégia '%s' com ordem de erro O(h^%d). Detalhe: %w", m.selectedDerivationOrder, m.selectedDerivationPhilosophy, m.selectedDerivationErrorOrder, calcErr)
		} else {
			m.err = fmt.Errorf("erro no cálculo: %w", calcErr)
		}
		return
	}
	m.result = fmt.Sprintf(
		"f%s(%.4f) ≈ %.6f",
		strings.Repeat("'", m.selectedDerivationOrder),
		x,
		calculatedVal,
	)
}

func (m *MainModel) resetDerivationInputs() {
	if len(m.philosophyOptions) > 0 {
		m.selectedDerivationPhilosophy = m.philosophyOptions[0]
	}
	m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy)

	if len(m.derivativeOrderOptions) > 0 {
		m.selectedDerivationOrder = m.derivativeOrderOptions[0].Value
	}
	if len(m.functionDefinitions) > 0 {
		m.selectedFunctionDef = m.functionDefinitions[0]
	} else {
		m.selectedFunctionDef = common.FunctionDefinition{}
	}
	m.currentX = "1.0"
	m.currentH = "0.1"
	m.focus = common.FocusNone
	m.err = nil
	m.result = ""
	m.derivationCursor = 0
}

func (m *MainModel) resetIntegrationInputs() {
	if len(m.availableIntegrationMethods) > 0 {
		m.selectedIntegrationMethod = m.availableIntegrationMethods[0]
	}
	if len(m.functionDefinitions) > 0 {
		// For now, let's assume we want to use the same function list and selected function
		// as derivation. This might change based on specific integration requirements.
		m.selectedIntegrationFunctionDef = m.functionDefinitions[0]
	}
	m.currentA = "0.0"
	m.currentB = "1.0"
	m.currentN = "10" // Default N value, can be adjusted
	m.integrationFocus = common.FocusNone
	m.err = nil
	m.result = ""
	m.integrationCursor = 0
}

func (m *MainModel) updateIntegrationMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q":
		m.state = common.StateMainMenu
		m.cursor = 0 // Reset main menu cursor
		m.resetIntegrationInputs()
		return m, nil
	case "up", "k":
		if m.integrationCursor > 0 {
			m.integrationCursor--
		}
	case "down", "j":
		if m.integrationCursor < len(m.integrationMenuChoices)-1 {
			m.integrationCursor++
		}
	case "enter":
		choice := m.integrationMenuChoices[m.integrationCursor]
		switch choice {
		case "Método":
			m.state = common.StateSelectIntegrationMethod // This state will need its own view and update logic
			m.selectionCursor = 0
			for i, method := range m.availableIntegrationMethods {
				if method == m.selectedIntegrationMethod {
					m.selectionCursor = i
					break
				}
			}
		case "Função":
			m.previousState = m.state // Store current state (StateIntegrationMenu)
			m.state = common.StateSelectFunction
			m.selectionCursor = 0
			// Set selection cursor to currently selected integration function
			for i, fd := range m.functionDefinitions {
				if fd.ID == m.selectedIntegrationFunctionDef.ID {
					m.selectionCursor = i
					break
				}
			}
		case "Limite Inferior (a)":
			m.integrationFocus = common.FocusIntegrationA
		case "Limite Superior (b)":
			m.integrationFocus = common.FocusIntegrationB
		case "Num de Subintervalos/Ordem (n)":
			m.integrationFocus = common.FocusIntegrationN
		case "Calcular":
			m.performIntegration()
		case "Voltar":
			m.state = common.StateMainMenu
			m.cursor = 0 // Reset main menu cursor
			m.resetIntegrationInputs()
		}
	}
	return m, nil
}

func (m *MainModel) performIntegration() {
	m.err = nil
	m.result = ""
	m.state = common.StateResult // Set state to result regardless of outcome for now

	a, errA := strconv.ParseFloat(strings.TrimSpace(m.currentA), 64)
	if errA != nil {
		m.err = fmt.Errorf("valor inválido para Limite Inferior (a) '%s': %w. Use '.' como separador decimal.", m.currentA, errA)
		return
	}

	b, errB := strconv.ParseFloat(strings.TrimSpace(m.currentB), 64)
	if errB != nil {
		m.err = fmt.Errorf("valor inválido para Limite Superior (b) '%s': %w. Use '.' como separador decimal.", m.currentB, errB)
		return
	}

	n, errN := strconv.Atoi(strings.TrimSpace(m.currentN))
	if errN != nil {
		m.err = fmt.Errorf("valor inválido para Num de Subintervalos/Ordem (n) '%s': %w. Deve ser um inteiro.", m.currentN, errN)
		return
	}

	if a >= b {
		m.err = fmt.Errorf("o limite inferior 'a' (%.2f) deve ser menor que o limite superior 'b' (%.2f)", a, b)
		return
	}

	if n <= 0 {
		m.err = fmt.Errorf("o número de subintervalos/ordem 'n' (%d) deve ser um inteiro positivo", n)
		return
	}

	if m.selectedIntegrationFunctionDef.Func == nil {
		m.err = fmt.Errorf("nenhuma função selecionada para integração")
		return
	}

	strategyName, ok := integrationMethodDisplayToStrategyName[m.selectedIntegrationMethod]
	if !ok {
		m.err = fmt.Errorf("método de integração selecionado ('%s') não tem um nome de estratégia interna definido", m.selectedIntegrationMethod)
		return
	}

	integrator, err := integration.NewIntegrator(strategyName)
	if err != nil {
		m.err = fmt.Errorf("falha ao criar integrador para estratégia '%s': %w", strategyName, err)
		return
	}

	calculatedVal, calcErr := integrator.Calculate(m.selectedIntegrationFunctionDef.Func, a, b, n)
	if calcErr != nil {
		m.err = fmt.Errorf("erro no cálculo da integral com %s (n=%d): %w", m.selectedIntegrationMethod, n, calcErr)
		return
	}

	m.result = fmt.Sprintf("∫[%.2f, %.2f] %s dx ≈ %.6f", a, b, m.selectedIntegrationFunctionDef.Name, calculatedVal)
	// m.state is already set to common.StateResult at the beginning of the function
}
