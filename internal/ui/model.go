package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/functions"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/integration" // Added

	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = (*MainModel)(nil)

type ErrorOrderOption struct {
	Display string
	Value   uint
}

type DerivativeOrderOption struct {
	Display string
	Value   int
}

// New type for integration method options
type IntegrationMethodOption struct {
	Display string
	Value   string // e.g., "Gauss-Legendre", "Newton-Cotes"
}

// New type for integration degree options
type IntegrationDegreeOption struct {
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
	cursor          int // For main menu

	// Derivation fields
	derivationMenuChoices        []string
	derivationCursor             int
	selectedDerivationPhilosophy string
	selectedDerivationErrorOrder uint
	selectedDerivationOrder      int
	currentX                     string
	currentH                     string

	// Integration fields - NEW
	integrationMenuChoices    []string
	integrationCursor         int
	selectedIntegrationMethod string
	// selectedIntegrationDegree is used when selecting from a list of degrees.
	// currentNPoints is the string representation, used for direct input and for passing to calculation.
	selectedIntegrationDegree int    // Stores the value from selection, primarily for display consistency
	currentNPoints            string // Stores the string for 'n' (degree/points), used for input and parsing
	currentLowerLimit         string // "a"
	currentUpperLimit         string // "b"
	currentTolerance          string // "tol"

	// Common fields for selection and results
	selectedFunctionDef common.FunctionDefinition
	result              string
	err                 error
	focus               int // Re-used for different input fields based on current menu and item
	selectionCursor     int // Re-used for different selection lists

	// Options for various menus
	philosophyOptions        []string
	currentErrorOrderOptions []ErrorOrderOption // Dynamically updated for derivation
	derivativeOrderOptions   []DerivativeOrderOption
	functionDefinitions      []common.FunctionDefinition

	// Integration options - NEW
	integrationMethodOptions []IntegrationMethodOption
	integrationDegreeOptions []IntegrationDegreeOption // Static for now, could be dynamic
	previousState            common.State // To help navigate back from shared states like StateSelectFunction
}

func NewMainModel() *MainModel {
	philosophyOpts := []string{"Forward", "Backward", "Central"}
	derivativeOrderOpts := []DerivativeOrderOption{
		{Display: "Primeira Derivada", Value: 1},
		{Display: "Segunda Derivada", Value: 2},
		{Display: "Terceira Derivada", Value: 3},
	}
	funcDefs := functions.GetFunctionDefinitions()

	integrationMethodOpts := []IntegrationMethodOption{
		{Display: "Gauss-Legendre", Value: "Gauss-Legendre"},
		{Display: "Newton-Cotes", Value: "Newton-Cotes"},
	}
	// Static list for now. Newton-Cotes (1,2,3), Gauss-Legendre (1-5 example).
	// Actual validation happens in strategy or performIntegration.
	integrationDegreeOpts := []IntegrationDegreeOption{
		{Display: "1", Value: 1}, {Display: "2", Value: 2}, {Display: "3", Value: 3},
		{Display: "4", Value: 4}, {Display: "5", Value: 5},
		// {Display: "6", Value: 6}, // Example if we add more later
	}

	m := &MainModel{
		state:                 common.StateMainMenu,
		previousState:         common.StateMainMenu,
		mainMenuChoices:       []string{"Derivação Numérica", "Integração Numérica", "Sair"},
		derivationMenuChoices: []string{"Filosofia", "Ordem do Erro", "Ordem da Derivada", "Função", "Ponto x", "Passo h", "Calcular", "Voltar"},

		philosophyOptions:      philosophyOpts,
		derivativeOrderOptions: derivativeOrderOpts,
		functionDefinitions:    funcDefs,

		selectedDerivationPhilosophy: philosophyOpts[0],
		selectedDerivationOrder:      derivativeOrderOpts[0].Value,
		currentX:                     "1.0",
		currentH:                     "0.1",

		integrationMenuChoices:    []string{"Método", "Grau/Pontos (n)", "Tolerância", "Limite Inferior (a)", "Limite Superior (b)", "Função", "Calcular", "Voltar"},
		integrationMethodOptions:  integrationMethodOpts,
		integrationDegreeOptions:  integrationDegreeOpts,
		selectedIntegrationMethod: integrationMethodOpts[0].Value,
		selectedIntegrationDegree: integrationDegreeOpts[0].Value, // Default selected degree
		currentNPoints:            strconv.Itoa(integrationDegreeOpts[0].Value), // Sync with default selected degree
		currentLowerLimit:         "0.0",
		currentUpperLimit:         "1.0",
		currentTolerance:          "0.0001",
	}

	if len(funcDefs) > 0 {
		m.selectedFunctionDef = funcDefs[0]
	}

	m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy)
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
	default: // Should not happen
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
		m.selectedDerivationErrorOrder = 0 // Or some indicator of no valid option
	}
}

func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.focus != common.FocusNone && (m.state == common.StateDerivationMenu || m.state == common.StateIntegrationMenu) {
			var currentText *string
			// Determine which field is being edited based on focus state
			switch m.focus {
			case common.FocusX:
				currentText = &m.currentX
			case common.FocusH:
				currentText = &m.currentH
			case common.FocusLowerLimit:
				currentText = &m.currentLowerLimit
			case common.FocusUpperLimit:
				currentText = &m.currentUpperLimit
			case common.FocusTolerance:
				currentText = &m.currentTolerance
			case common.FocusNPoints:
				currentText = &m.currentNPoints
			default:
				return m, nil // Should not happen if focus is managed correctly
			}

			// Handle input for the focused field
			switch msg.String() {
			case "enter":
				// If editing NPoints, attempt to parse and update selectedIntegrationDegree for display consistency.
				if m.focus == common.FocusNPoints {
					if nVal, err := strconv.Atoi(*currentText); err == nil {
						// Check if this value is actually in our options for selection cursor later
						isOption := false
						for _, opt := range m.integrationDegreeOptions {
							if opt.Value == nVal {
								isOption = true
								break
							}
						}
						if isOption {
							m.selectedIntegrationDegree = nVal
						}
						// If not an option, selectedIntegrationDegree might not reflect currentNPoints if user manually typed
						// This is okay, as currentNPoints is the source of truth for calculation.
					}
				}
				m.focus = common.FocusNone // Clear focus
				return m, nil
			case "backspace":
				if len(*currentText) > 0 {
					*currentText = (*currentText)[:len(*currentText)-1]
				}
				return m, nil
			default:
				char := msg.String()
				// Basic validation for numeric input (digits, one decimal, optional leading minus)
				if (char >= "0" && char <= "9") || (char == "." && !strings.Contains(*currentText, ".")) || (char == "-" && len(*currentText) == 0) {
					*currentText += char
				}
				return m, nil
			}
		}

		// Handle state transitions and actions based on key presses
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
		case common.StateSelectIntegrationMethod:
			return m.updateSelectIntegrationMethod(msg)
		case common.StateSelectIntegrationDegree:
			return m.updateSelectIntegrationDegree(msg)
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
		m.previousState = common.StateMainMenu // Record current state before transitioning
		selectedItem := m.mainMenuChoices[m.cursor]
		switch selectedItem {
		case "Derivação Numérica":
			m.state = common.StateDerivationMenu
			m.derivationCursor = 0 // Reset cursor for the new menu
			m.err = nil            // Clear previous errors/results
			m.result = ""
		case "Integração Numérica":
			m.state = common.StateIntegrationMenu
			m.integrationCursor = 0 // Reset cursor for the new menu
			m.err = nil             // Clear previous errors/results
			m.result = ""
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
		m.cursor = 0 // Set main menu cursor to "Derivação Numérica"
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
		m.previousState = m.state // Record current state
		switch choice {
		case "Filosofia":
			m.state = common.StateSelectPhilosophy
			m.selectionCursor = 0 // Reset selection cursor
			for i, phil := range m.philosophyOptions {
				if phil == m.selectedDerivationPhilosophy {
					m.selectionCursor = i; break
				}
			}
		case "Ordem do Erro":
			m.state = common.StateSelectErrorOrder
			m.selectionCursor = 0
			m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy) // Refresh options
			for i, eo := range m.currentErrorOrderOptions {
				if eo.Value == m.selectedDerivationErrorOrder {
					m.selectionCursor = i; break
				}
			}
		case "Ordem da Derivada":
			m.state = common.StateSelectDerivativeOrder
			m.selectionCursor = 0
			for i, do := range m.derivativeOrderOptions {
				if do.Value == m.selectedDerivationOrder {
					m.selectionCursor = i; break
				}
			}
		case "Função":
			m.state = common.StateSelectFunction
			m.selectionCursor = 0
			for i, fd := range m.functionDefinitions {
				if fd.ID == m.selectedFunctionDef.ID {
					m.selectionCursor = i; break
				}
			}
		case "Ponto x":
			m.focus = common.FocusX
		case "Passo h":
			m.focus = common.FocusH
		case "Calcular":
			m.performDerivation()
			m.state = common.StateResult // Transition to show result/error
		case "Voltar":
			m.state = common.StateMainMenu
			m.cursor = 0 // Back to "Derivação Numérica" on main menu
			m.resetDerivationInputs()
		}
	}
	return m, nil
}

func (m *MainModel) updateIntegrationMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q":
		m.state = common.StateMainMenu
		m.cursor = 1 // Set main menu cursor to "Integração Numérica"
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
		m.previousState = m.state // Record current state
		switch choice {
		case "Método":
			m.state = common.StateSelectIntegrationMethod
			m.selectionCursor = 0
			for i, methodOpt := range m.integrationMethodOptions {
				if methodOpt.Value == m.selectedIntegrationMethod {
					m.selectionCursor = i; break
				}
			}
		case "Grau/Pontos (n)":
			// Option 1: Go to a selection screen for n
			m.state = common.StateSelectIntegrationDegree
			m.selectionCursor = 0
			currentNValue, _ := strconv.Atoi(m.currentNPoints) // Use currentNPoints to set cursor
			found := false
			for i, degOpt := range m.integrationDegreeOptions {
				if degOpt.Value == currentNValue {
					m.selectionCursor = i
					found = true
					break
				}
			}
			if !found && len(m.integrationDegreeOptions)>0 { m.selectionCursor = 0} // Default if not found
			// Option 2: Direct input (would be m.focus = common.FocusNPoints)
		case "Tolerância":
			m.focus = common.FocusTolerance
		case "Limite Inferior (a)":
			m.focus = common.FocusLowerLimit
		case "Limite Superior (b)":
			m.focus = common.FocusUpperLimit
		case "Função":
			m.state = common.StateSelectFunction
			m.selectionCursor = 0
			for i, fd := range m.functionDefinitions {
				if fd.ID == m.selectedFunctionDef.ID {
					m.selectionCursor = i; break
				}
			}
		case "Calcular":
			m.performIntegration()
			m.state = common.StateResult // Transition to show result/error
		case "Voltar":
			m.state = common.StateMainMenu
			m.cursor = 1 // Back to "Integração Numérica" on main menu
			m.resetIntegrationInputs()
		}
	}
	return m, nil
}

// All updateSelect<...> functions for Derivation
func (m *MainModel) updateSelectPhilosophy(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateDerivationMenu // Return to derivation menu
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {m.selectionCursor--}
	case "down", "j":
		if m.selectionCursor < len(m.philosophyOptions)-1 {m.selectionCursor++}
	case "enter":
		newPhilosophy := m.philosophyOptions[m.selectionCursor]
		if m.selectedDerivationPhilosophy != newPhilosophy {
			m.selectedDerivationPhilosophy = newPhilosophy
			m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy) // Update error orders based on new philosophy
		}
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

func (m *MainModel) updateSelectErrorOrder(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if len(m.currentErrorOrderOptions) == 0 { // No options available
		if msg.String() == "enter" || msg.String() == "q" || msg.String() == "ctrl+c" {
			m.state = common.StateDerivationMenu; return m, nil
		}
		return m, nil
	}
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateDerivationMenu; return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {m.selectionCursor--}
	case "down", "j":
		if m.selectionCursor < len(m.currentErrorOrderOptions)-1 {m.selectionCursor++}
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
		m.state = common.StateDerivationMenu; return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {m.selectionCursor--}
	case "down", "j":
		if m.selectionCursor < len(m.derivativeOrderOptions)-1 {m.selectionCursor++}
	case "enter":
		m.selectedDerivationOrder = m.derivativeOrderOptions[m.selectionCursor].Value
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

// All updateSelect<...> functions for Integration
func (m *MainModel) updateSelectIntegrationMethod(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateIntegrationMenu; return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {m.selectionCursor--}
	case "down", "j":
		if m.selectionCursor < len(m.integrationMethodOptions)-1 {m.selectionCursor++}
	case "enter":
		m.selectedIntegrationMethod = m.integrationMethodOptions[m.selectionCursor].Value
		// Future: Update m.integrationDegreeOptions if they are method-dependent
		// For now, selectedIntegrationDegree and currentNPoints might need resetting or validation
		// For simplicity, we assume the current selectedIntegrationDegree/currentNPoints might become invalid
		// and user has to re-select degree if method changes significantly.
		// Or, try to keep currentNPoints if it's valid for the new method.
		// For now, no change to N on method change, validation happens at calculation.
		m.state = common.StateIntegrationMenu
	}
	return m, nil
}

func (m *MainModel) updateSelectIntegrationDegree(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if len(m.integrationDegreeOptions) == 0 { // No degree options
		if msg.String() == "enter" || msg.String() == "q" || msg.String() == "ctrl+c" {
			m.state = common.StateIntegrationMenu; return m, nil
		}
		return m, nil
	}
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateIntegrationMenu; return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {m.selectionCursor--}
	case "down", "j":
		if m.selectionCursor < len(m.integrationDegreeOptions)-1 {m.selectionCursor++}
	case "enter":
		if m.selectionCursor >= 0 && m.selectionCursor < len(m.integrationDegreeOptions) {
			m.selectedIntegrationDegree = m.integrationDegreeOptions[m.selectionCursor].Value
			m.currentNPoints = strconv.Itoa(m.selectedIntegrationDegree) // Sync currentNPoints string
		}
		m.state = common.StateIntegrationMenu
	}
	return m, nil
}

// Shared updateSelectFunction
func (m *MainModel) updateSelectFunction(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = m.previousState // Return to the menu that called this selection screen
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {m.selectionCursor--}
	case "down", "j":
		if m.selectionCursor < len(m.functionDefinitions)-1 {m.selectionCursor++}
	case "enter":
		m.selectedFunctionDef = m.functionDefinitions[m.selectionCursor]
		m.state = m.previousState // Return to the calling menu
	}
	return m, nil
}

// Shared updateResultScreen
func (m *MainModel) updateResultScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "enter":
		// Return to the menu from which "Calcular" was pressed (stored in m.previousState)
		// Or, if error is very specific, could redirect to specific part of that menu.
		if m.previousState == common.StateDerivationMenu || m.previousState == common.StateIntegrationMenu {
			m.state = m.previousState
		} else {
			// Fallback if previousState wasn't set correctly before a calculation state.
			// This might happen if StateResult is reached from an unexpected path.
			// Default to main menu.
			m.state = common.StateMainMenu
		}
		m.err = nil    // Clear error/result
		m.result = ""
		// Reset cursor for the menu being returned to
		if m.state == common.StateMainMenu {m.cursor = 0}
		// Cursors for submenus (derivationCursor, integrationCursor) are typically reset when entering those menus or handled by their own update logic.
		// For now, we don't explicitly reset sub-menu cursors here, assuming they are at a reasonable state or will be reset upon re-entry.
	}
	return m, nil
}

func (m *MainModel) View() string {
	return View(m) // Forward to the View function in view.go
}

func (m *MainModel) performDerivation() {
	m.err = nil; m.result = "" // Clear previous results

	x, errX := strconv.ParseFloat(strings.TrimSpace(m.currentX), 64)
	if errX != nil { m.err = fmt.Errorf("Ponto x '%s': %w", m.currentX, errX); return }
	h, errH := strconv.ParseFloat(strings.TrimSpace(m.currentH), 64)
	if errH != nil { m.err = fmt.Errorf("Passo h '%s': %w", m.currentH, errH); return }

	if h == 0 { m.err = common.ErrZeroValue; return }
	if m.selectedFunctionDef.Func == nil { m.err = fmt.Errorf("nenhuma função selecionada"); return }

	_, factoryErr := derivation.DerivacaoFactory(m.selectedDerivationPhilosophy, m.selectedDerivationErrorOrder)
	if factoryErr != nil { m.err = fmt.Errorf("combinação inválida: %s, O(h^%d): %w", m.selectedDerivationPhilosophy, m.selectedDerivationErrorOrder, factoryErr); return }

	derivator, err := derivation.NewDerivator(m.selectedDerivationPhilosophy, m.selectedDerivationErrorOrder)
	if err != nil { m.err = fmt.Errorf("criar derivador: %w", err); return }

	val, calcErr := derivator.Calculate(m.selectedFunctionDef.Func, x, h, m.selectedDerivationOrder)
	if calcErr != nil { m.err = fmt.Errorf("cálculo derivação: %w", calcErr); return }

	m.result = fmt.Sprintf("f%s(%.4f) ≈ %.6f", strings.Repeat("'", m.selectedDerivationOrder), x, val)
}

func (m *MainModel) performIntegration() {
	m.err = nil; m.result = "" // Clear previous results

	a, errA := strconv.ParseFloat(strings.TrimSpace(m.currentLowerLimit), 64)
	if errA != nil { m.err = fmt.Errorf("Limite Inferior (a) '%s': %w", m.currentLowerLimit, errA); return }
	b, errB := strconv.ParseFloat(strings.TrimSpace(m.currentUpperLimit), 64)
	if errB != nil { m.err = fmt.Errorf("Limite Superior (b) '%s': %w", m.currentUpperLimit, errB); return }
	tol, errTol := strconv.ParseFloat(strings.TrimSpace(m.currentTolerance), 64)
	if errTol != nil { m.err = fmt.Errorf("Tolerância '%s': %w", m.currentTolerance, errTol); return }

	if tol <= 0 { m.err = fmt.Errorf("tolerância deve ser > 0, got %s", m.currentTolerance); return }

	n, errN := strconv.Atoi(strings.TrimSpace(m.currentNPoints))
	if errN != nil { m.err = fmt.Errorf("Grau/Pontos (n) '%s': %w", m.currentNPoints, errN); return }

	// Validate n based on method (example validation, strategies might have more specific needs)
	if m.selectedIntegrationMethod == "Newton-Cotes" && (n < 1 || n > 3) {
		m.err = fmt.Errorf("para Newton-Cotes, n (grau) deve ser 1, 2, ou 3. Valor fornecido: %d", n); return
	}
	if m.selectedIntegrationMethod == "Gauss-Legendre" && (n < 1 || n > 5) { // Assuming current example Gauss supports 1-5 points
		m.err = fmt.Errorf("para Gauss-Legendre, n (pontos) deve ser entre 1 e 5 (neste exemplo). Valor fornecido: %d", n); return
	}
    if n <= 0 { // General catch-all if other checks missed
        m.err = fmt.Errorf("Grau/Pontos (n) deve ser um inteiro positivo. Valor fornecido: %d", n); return
    }


	if m.selectedFunctionDef.Func == nil { m.err = fmt.Errorf("nenhuma função selecionada para integração"); return }

	integrator, err := integration.NewIntegrator(m.selectedIntegrationMethod)
	if err != nil { m.err = fmt.Errorf("criar integrador para '%s': %w", m.selectedIntegrationMethod, err); return }

	val, calcErr := integrator.Calculate(m.selectedFunctionDef.Func, a, b, n, tol)
	if calcErr != nil { m.err = fmt.Errorf("cálculo integração: %w", calcErr); return }

	m.result = fmt.Sprintf("∫[%.3f, %.3f] f(x)dx ≈ %.8f (Método: %s, n: %d, tol: %.1E)", a, b, val, m.selectedIntegrationMethod, n, tol)
}

func (m *MainModel) resetDerivationInputs() {
	if len(m.philosophyOptions) > 0 {
		m.selectedDerivationPhilosophy = m.philosophyOptions[0]
	}
	m.updateAvailableErrorOrders(m.selectedDerivationPhilosophy) // This also sets a default selectedDerivationErrorOrder
	if len(m.derivativeOrderOptions) > 0 {
		m.selectedDerivationOrder = m.derivativeOrderOptions[0].Value
	}
	// m.selectedFunctionDef is shared, so not reset here.
	m.currentX = "1.0"
	m.currentH = "0.1"
	m.focus = common.FocusNone
	m.err = nil
	m.result = ""
	m.derivationCursor = 0 // Reset menu cursor
}

func (m *MainModel) resetIntegrationInputs() {
	if len(m.integrationMethodOptions) > 0 {
		m.selectedIntegrationMethod = m.integrationMethodOptions[0].Value
	}
	if len(m.integrationDegreeOptions) > 0 { // Reset degree selection and string input
		m.selectedIntegrationDegree = m.integrationDegreeOptions[0].Value
		m.currentNPoints = strconv.Itoa(m.integrationDegreeOptions[0].Value)
	} else { // Fallback if options somehow empty
		m.selectedIntegrationDegree = 1 // Default degree
		m.currentNPoints = "1"
	}
	m.currentLowerLimit = "0.0"
	m.currentUpperLimit = "1.0"
	m.currentTolerance = "0.0001"
	// m.selectedFunctionDef is shared.
	m.focus = common.FocusNone
	m.err = nil
	m.result = ""
	m.integrationCursor = 0 // Reset menu cursor
}
