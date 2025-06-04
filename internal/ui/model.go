package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/derivation"
	"github.com/ArtroxGabriel/numeric-methods-cli/internal/functions"

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

	philosophyOptions      []string
	errorOrderOptions      []ErrorOrderOption
	derivativeOrderOptions []DerivativeOrderOption
	functionDefinitions    []common.FunctionDefinition
	selectionCursor        int
	tempStringInput        string
}

func NewMainModel() *MainModel {
	philosophyOpts := []string{"Forward", "Backward", "Central"}
	errorOrderOpts := []ErrorOrderOption{
		{Display: "Ordem 1 (O(h))", Value: 1},
		{Display: "Ordem 2 (O(h^2))", Value: 2},
		{Display: "Ordem 3 (O(h^3))", Value: 3},
	}
	derivativeOrderOpts := []DerivativeOrderOption{
		{Display: "Primeira Derivada", Value: 1},
		{Display: "Segunda Derivada", Value: 2},
		{Display: "Terceira Derivada", Value: 3},
	}
	funcDefs := functions.GetFunctionDefinitions()

	return &MainModel{
		state:                 common.StateMainMenu,
		mainMenuChoices:       []string{"Derivação Numérica", "Integração Numérica", "Sair"},
		derivationMenuChoices: []string{"Filosofia", "Ordem do Erro", "Ordem da Derivada", "Função", "Ponto x", "Passo h", "Calcular", "Voltar"},

		philosophyOptions:      philosophyOpts,
		errorOrderOptions:      errorOrderOpts,
		derivativeOrderOptions: derivativeOrderOpts,
		functionDefinitions:    funcDefs,

		selectedDerivationPhilosophy: philosophyOpts[0],
		selectedDerivationErrorOrder: errorOrderOpts[0].Value,
		selectedDerivationOrder:      derivativeOrderOpts[0].Value,
		selectedFunctionDef:          funcDefs[0],
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
				if (char >= "0" && char <= "9") || (char == "." && !strings.Contains(m.currentX, ".") && m.focus == common.FocusX) || (char == "." && !strings.Contains(m.currentH, ".") && m.focus == common.FocusH) {
					if m.focus == common.FocusX {
						m.currentX += char
					} else if m.focus == common.FocusH {
						m.currentH += char
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
		case common.StateResult, common.StateIntegrationMenu:
			return m.updateResultScreen(msg)
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
			m.result = "Integração Numérica ainda não implementada."
		case "Sair":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MainModel) updateDerivationMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
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
			for i, eo := range m.errorOrderOptions {
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
			m.state = common.StateSelectFunction
			m.selectionCursor = 0
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
		m.state = common.StateDerivationMenu
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
		m.selectedDerivationPhilosophy = m.philosophyOptions[m.selectionCursor]
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

func (m *MainModel) updateSelectErrorOrder(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.state = common.StateDerivationMenu
		return m, nil
	case "up", "k":
		if m.selectionCursor > 0 {
			m.selectionCursor--
		}
	case "down", "j":
		if m.selectionCursor < len(m.errorOrderOptions)-1 {
			m.selectionCursor++
		}
	case "enter":
		m.selectedDerivationErrorOrder = m.errorOrderOptions[m.selectionCursor].Value
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
		m.selectedFunctionDef = m.functionDefinitions[m.selectionCursor]
		m.state = common.StateDerivationMenu
	}
	return m, nil
}

func (m *MainModel) updateResultScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "enter":
		m.state = common.StateMainMenu
		m.result = ""
		m.err = nil
		m.cursor = 0
	}
	return m, nil
}

func (m *MainModel) View() string {
	s := strings.Builder{}

	switch m.state {
	case common.StateMainMenu:
		s.WriteString("Bem-vindo aos Métodos Numéricos!\n\n")
		for i, choice := range m.mainMenuChoices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}
		s.WriteString("\n(Pressione 'q' para sair, 'Enter' para selecionar)")

	case common.StateDerivationMenu:
		s.WriteString("Menu de Derivação Numérica:\n\n")
		for i, choice := range m.derivationMenuChoices {
			cursor := " "
			if m.derivationCursor == i {
				cursor = ">"
			}
			line := fmt.Sprintf("%s %s", cursor, choice)
			switch choice {
			case "Filosofia":
				line += fmt.Sprintf(": %s", m.selectedDerivationPhilosophy)
			case "Ordem do Erro":
				line += fmt.Sprintf(": O(h^%d)", m.selectedDerivationErrorOrder)
			case "Ordem da Derivada":
				line += fmt.Sprintf(": %da", m.selectedDerivationOrder)
			case "Função":
				name := "Nenhuma"
				if m.selectedFunctionDef.Func != nil {
					name = m.selectedFunctionDef.Name
				}
				line += fmt.Sprintf(": %s", name)
			case "Ponto x":
				inputX := m.currentX
				if m.focus == common.FocusX {
					inputX += "_"
				}
				line += fmt.Sprintf(": %s", inputX)
			case "Passo h":
				inputH := m.currentH
				if m.focus == common.FocusH {
					inputH += "_"
				}
				line += fmt.Sprintf(": %s", inputH)
			}
			s.WriteString(line + "\n")
		}
		s.WriteString("\n(Navegue com ↑/↓, 'Enter' para selecionar/editar, 'q' para sair)")
		if m.focus == common.FocusX {
			s.WriteString("\n[EDITANDO Ponto x: Digite o valor e pressione Enter para confirmar]")
		} else if m.focus == common.FocusH {
			s.WriteString("\n[EDITANDO Passo h: Digite o valor e pressione Enter para confirmar]")
		}

	case common.StateSelectPhilosophy:
		s.WriteString("Selecione a Filosofia:\n\n")
		for i, phil := range m.philosophyOptions {
			cursor := " "
			if m.selectionCursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, phil))
		}
		s.WriteString("\n('Enter' para confirmar, 'q' para voltar)")

	case common.StateSelectErrorOrder:
		s.WriteString("Selecione a Ordem do Erro:\n\n")
		for i, eo := range m.errorOrderOptions {
			cursor := " "
			if m.selectionCursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, eo.Display))
		}
		s.WriteString("\n('Enter' para confirmar, 'q' para voltar)")

	case common.StateSelectDerivativeOrder:
		s.WriteString("Selecione a Ordem da Derivada:\n\n")
		for i, do := range m.derivativeOrderOptions {
			cursor := " "
			if m.selectionCursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, do.Display))
		}
		s.WriteString("\n('Enter' para confirmar, 'q' para voltar)")

	case common.StateSelectFunction:
		s.WriteString("Selecione a Função:\n\n")
		for i, fd := range m.functionDefinitions {
			cursor := " "
			if m.selectionCursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, fd.Name))
		}
		s.WriteString("\n('Enter' para confirmar, 'q' para voltar)")

	case common.StateResult:
		if m.err != nil {
			s.WriteString(fmt.Sprintf("Erro: %v\n", m.err))
		} else {
			s.WriteString(fmt.Sprintf("Resultado: %s\n", m.result))
		}
		s.WriteString("\n(Pressione 'Enter' para voltar ao menu principal, 'q' para sair)")

	case common.StateIntegrationMenu:
		s.WriteString(m.result)
		s.WriteString("\n(Pressione 'Enter' para voltar ao menu principal, 'q' para sair)")
	}

	return s.String()
}

func (m *MainModel) performDerivation() {
	m.err = nil
	m.result = ""

	x, err := strconv.ParseFloat(m.currentX, 64)
	if err != nil {
		m.err = fmt.Errorf("valor inválido para x: %w. Use '.' como separador decimal.", err)
		return
	}
	h, err := strconv.ParseFloat(m.currentH, 64)
	if err != nil {
		m.err = fmt.Errorf("valor inválido para h: %w. Use '.' como separador decimal.", err)
		return
	}

	if m.selectedFunctionDef.Func == nil {
		m.err = fmt.Errorf("nenhuma função selecionada")
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
		m.err = fmt.Errorf("erro no cálculo: %w", calcErr)
		return
	}
	m.result = fmt.Sprintf(
		"Derivada (%s, Ordem Erro %d, Ordem Derivada %d) da função %s em x=%.4f com h=%.4f: %.6f",
		m.selectedDerivationPhilosophy,
		m.selectedDerivationErrorOrder,
		m.selectedDerivationOrder,
		m.selectedFunctionDef.Name,
		x,
		h,
		calculatedVal,
	)
}

func (m *MainModel) resetDerivationInputs() {
	if len(m.philosophyOptions) > 0 {
		m.selectedDerivationPhilosophy = m.philosophyOptions[0]
	}
	if len(m.errorOrderOptions) > 0 {
		m.selectedDerivationErrorOrder = m.errorOrderOptions[0].Value
	}
	if len(m.derivativeOrderOptions) > 0 {
		m.selectedDerivationOrder = m.derivativeOrderOptions[0].Value
	}
	if len(m.functionDefinitions) > 0 {
		m.selectedFunctionDef = m.functionDefinitions[0]
	} else {
		m.selectedFunctionDef = common.FunctionDefinition{}
	}
	m.currentX = ""
	m.currentH = ""
	m.focus = common.FocusNone
	m.err = nil
	m.result = ""
	m.derivationCursor = 0
}
