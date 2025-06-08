package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

func View(model *MainModel) string {
	s := strings.Builder{}

	switch model.state {
	case common.StateMainMenu:
		s.WriteString(TitleStyle.Render("Numeric Methods 2 CLI"))
		s.WriteString("

")
		for i, choice := range model.mainMenuChoices {
			s.WriteString(RenderListItem(choice, model.cursor == i))
			s.WriteString("
")
		}
		s.WriteString(HelpStyle.Render("
(Navegue com ↑/↓, 'Enter' para selecionar, 'q' ou Ctrl+C para sair)"))

	case common.StateDerivationMenu:
		s.WriteString(TitleStyle.Render("Menu de Derivação Numérica:"))
		s.WriteString("

")
		for i, choice := range model.derivationMenuChoices {
			var line string
			isSelectedItem := model.derivationCursor == i
			if isSelectedItem {
				line = SelectedItemStyle.Render(choice)
			} else {
				line = ListItemStyle.Render(choice)
			}

			valueStyle := InputValueStyle
			// Check if the current item is being focused for input
			isFocusedForInput := false
			if isSelectedItem {
				switch choice {
				case "Ponto x":
					if model.focus == common.FocusX {
						isFocusedForInput = true
					}
				case "Passo h":
					if model.focus == common.FocusH {
						isFocusedForInput = true
					}
				}
			}
			if isFocusedForInput {
				valueStyle = FocusedInputStyle
			}

			switch choice {
			case "Filosofia":
				line += fmt.Sprintf(": %s", valueStyle.Render(model.selectedDerivationPhilosophy))
			case "Ordem do Erro":
				errorOrderDisplay := "N/A"
				for _, eo := range model.currentErrorOrderOptions {
					if eo.Value == model.selectedDerivationErrorOrder {
						errorOrderDisplay = eo.Display; break
					}
				}
				line += fmt.Sprintf(": %s", valueStyle.Render(errorOrderDisplay))
			case "Ordem da Derivada":
				derivativeOrderDisplay := "N/A"
				for _, do := range model.derivativeOrderOptions {
					if do.Value == model.selectedDerivationOrder {
						derivativeOrderDisplay = do.Display; break
					}
				}
				line += fmt.Sprintf(": %s", valueStyle.Render(derivativeOrderDisplay))
			case "Função":
				name := "Nenhuma"
				if model.selectedFunctionDef.Func != nil { name = model.selectedFunctionDef.Name }
				line += fmt.Sprintf(": %s", valueStyle.Render(name))
			case "Ponto x":
				inputX := model.currentX
				if model.focus == common.FocusX && isSelectedItem { inputX += CursorStyle.Render("_") }
				line += fmt.Sprintf(": %s", valueStyle.Render(inputX))
			case "Passo h":
				inputH := model.currentH
				if model.focus == common.FocusH && isSelectedItem { inputH += CursorStyle.Render("_") }
				line += fmt.Sprintf(": %s", valueStyle.Render(inputH))
			}
			s.WriteString(line + "
")
		}
		s.WriteString(HelpStyle.Render("
(Navegue com ↑/↓, 'Enter' para selecionar/editar, 'q' para voltar, Ctrl+C para sair)"))
		// Display specific help message when an input field is focused
		if model.focus != common.FocusNone {
			var fieldName string
			switch model.focus {
			case common.FocusX: fieldName = "Ponto x"
			case common.FocusH: fieldName = "Passo h"
			}
			if fieldName != "" {
				s.WriteString(HelpStyle.Render(fmt.Sprintf("
[EDITANDO %s: Digite o valor e pressione Enter]", fieldName)))
			}
		}


	// NEW: StateIntegrationMenu rendering
	case common.StateIntegrationMenu:
		s.WriteString(TitleStyle.Render("Menu de Integração Numérica:"))
		s.WriteString("

")
		for i, choice := range model.integrationMenuChoices {
			var line string
			isSelectedItem := model.integrationCursor == i
			if isSelectedItem {
				line = SelectedItemStyle.Render(choice)
			} else {
				line = ListItemStyle.Render(choice)
			}

			valueStyle := InputValueStyle
			isFocusedForInput := false
			if isSelectedItem {
				switch choice {
				case "Grau/Pontos (n)":
					// Assuming direct input focus is common.FocusNPoints
					// If selection is primary, this might not apply or apply differently
					if model.focus == common.FocusNPoints { isFocusedForInput = true }
				case "Tolerância":
					if model.focus == common.FocusTolerance { isFocusedForInput = true }
				case "Limite Inferior (a)":
					if model.focus == common.FocusLowerLimit { isFocusedForInput = true }
				case "Limite Superior (b)":
					if model.focus == common.FocusUpperLimit { isFocusedForInput = true }
				}
			}
			if isFocusedForInput { valueStyle = FocusedInputStyle }

			switch choice {
			case "Método":
				methodDisplay := "N/A"
				for _, methOpt := range model.integrationMethodOptions {
					if methOpt.Value == model.selectedIntegrationMethod {
						methodDisplay = methOpt.Display; break
					}
				}
				line += fmt.Sprintf(": %s", valueStyle.Render(methodDisplay))
			case "Grau/Pontos (n)":
				// Display currentNPoints, which is the string value for input
				// selectedIntegrationDegree is the integer value from selection
				displayN := model.currentNPoints
				if model.focus == common.FocusNPoints && isSelectedItem { displayN += CursorStyle.Render("_") }
				line += fmt.Sprintf(": %s", valueStyle.Render(displayN))
			case "Tolerância":
				inputTol := model.currentTolerance
				if model.focus == common.FocusTolerance && isSelectedItem { inputTol += CursorStyle.Render("_") }
				line += fmt.Sprintf(": %s", valueStyle.Render(inputTol))
			case "Limite Inferior (a)":
				inputA := model.currentLowerLimit
				if model.focus == common.FocusLowerLimit && isSelectedItem { inputA += CursorStyle.Render("_") }
				line += fmt.Sprintf(": %s", valueStyle.Render(inputA))
			case "Limite Superior (b)":
				inputB := model.currentUpperLimit
				if model.focus == common.FocusUpperLimit && isSelectedItem { inputB += CursorStyle.Render("_") }
				line += fmt.Sprintf(": %s", valueStyle.Render(inputB))
			case "Função":
				name := "Nenhuma"
				if model.selectedFunctionDef.Func != nil { name = model.selectedFunctionDef.Name }
				line += fmt.Sprintf(": %s", valueStyle.Render(name))
			}
			s.WriteString(line + "
")
		}
		s.WriteString(HelpStyle.Render("
(Navegue com ↑/↓, 'Enter' para selecionar/editar, 'q' para voltar, Ctrl+C para sair)"))
		if model.focus != common.FocusNone {
			var fieldName string
			switch model.focus {
			case common.FocusNPoints: fieldName = "Grau/Pontos (n)"
			case common.FocusTolerance: fieldName = "Tolerância"
			case common.FocusLowerLimit: fieldName = "Limite Inferior (a)"
			case common.FocusUpperLimit: fieldName = "Limite Superior (b)"
			}
			if fieldName != "" {
				s.WriteString(HelpStyle.Render(fmt.Sprintf("
[EDITANDO %s: Digite o valor e pressione Enter]", fieldName)))
			}
		}

	case common.StateSelectPhilosophy:
		s.WriteString(TitleStyle.Render("Selecione a Filosofia (Derivação):"))
		s.WriteString("

")
		for i, phil := range model.philosophyOptions {
			s.WriteString(RenderListItem(phil, model.selectionCursor == i))
			s.WriteString("
")
		}
		s.WriteString(HelpStyle.Render("
('Enter' para confirmar, 'q' para voltar)"))

	case common.StateSelectErrorOrder:
		s.WriteString(TitleStyle.Render("Selecione a Ordem do Erro (Derivação):"))
		s.WriteString("

")
		if len(model.currentErrorOrderOptions) == 0 {
			s.WriteString(ErrorStyle.Render("Nenhuma ordem de erro disponível para a filosofia selecionada."))
		} else {
			for i, eo := range model.currentErrorOrderOptions {
				s.WriteString(RenderListItem(eo.Display, model.selectionCursor == i))
				s.WriteString("
")
			}
		}
		s.WriteString(HelpStyle.Render("
('Enter' para confirmar, 'q' para voltar)"))

	case common.StateSelectDerivativeOrder:
		s.WriteString(TitleStyle.Render("Selecione a Ordem da Derivada:"))
		s.WriteString("

")
		for i, do := range model.derivativeOrderOptions {
			s.WriteString(RenderListItem(do.Display, model.selectionCursor == i))
			s.WriteString("
")
		}
		s.WriteString(HelpStyle.Render("
('Enter' para confirmar, 'q' para voltar)"))

	// NEW: StateSelectIntegrationMethod
	case common.StateSelectIntegrationMethod:
		s.WriteString(TitleStyle.Render("Selecione o Método de Integração:"))
		s.WriteString("

")
		for i, im := range model.integrationMethodOptions {
			s.WriteString(RenderListItem(im.Display, model.selectionCursor == i))
			s.WriteString("
")
		}
		s.WriteString(HelpStyle.Render("
('Enter' para confirmar, 'q' para voltar)"))

	// NEW: StateSelectIntegrationDegree
	case common.StateSelectIntegrationDegree:
		s.WriteString(TitleStyle.Render("Selecione o Grau/Pontos (n) para Integração:"))
		s.WriteString("

")
		// Note: model.integrationDegreeOptions is static for now.
		// Could be dynamically filtered based on model.selectedIntegrationMethod in a future enhancement.
		if len(model.integrationDegreeOptions) == 0 {
		    s.WriteString(ErrorStyle.Render("Nenhuma opção de grau/pontos disponível."))
		} else {
		    for i, ido := range model.integrationDegreeOptions {
			    s.WriteString(RenderListItem(ido.Display, model.selectionCursor == i))
			    s.WriteString("
")
		    }
		}
		s.WriteString(HelpStyle.Render("
('Enter' para confirmar, 'q' para voltar)"))

	case common.StateSelectFunction:
		title := "Selecione a Função"
		// Optionally, make title more specific based on model.previousState
		if model.previousState == common.StateDerivationMenu {
			title += " (para Derivação)"
		} else if model.previousState == common.StateIntegrationMenu {
			title += " (para Integração)"
		}
		s.WriteString(TitleStyle.Render(title + ":"))

		s.WriteString("

")
		for i, fd := range model.functionDefinitions {
			s.WriteString(RenderListItem(fd.Name, model.selectionCursor == i))
			s.WriteString("
")
		}
		s.WriteString(HelpStyle.Render("
('Enter' para confirmar, 'q' para voltar)"))

	case common.StateResult:
		if model.err != nil {
			s.WriteString(ErrorStyle.Render(fmt.Sprintf("Erro: %v
", model.err)))
		} else {
			s.WriteString(ResultStyle.Render(fmt.Sprintf("Resultado: %s
", model.result)))
		}
		s.WriteString(HelpStyle.Render("
(Pressione 'Enter' para voltar ao menu anterior, 'q' ou Ctrl+C para sair)"))
	}

	return DocStyle.Render(s.String())
}
