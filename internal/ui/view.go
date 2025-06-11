package ui

import (
	"fmt"
	"strings"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/common"
)

func View(model *MainModel) string {
	s := strings.Builder{}

	switch model.state {
	case common.StateMainMenu:
		s.WriteString(TitleStyle.Render("Numeric Methods 2 CLI"))
		s.WriteString("\n\n")
		for i, choice := range model.mainMenuChoices {
			s.WriteString(RenderListItem(choice, model.cursor == i))
			s.WriteString("\n")
		}
		s.WriteString(HelpStyle.Render("\n(Navegue com ↑/↓, 'Enter' para selecionar, 'q' ou Ctrl+C para sair)"))

	case common.StateDerivationMenu:
		s.WriteString(TitleStyle.Render("Menu de Derivação Numérica:"))
		s.WriteString("\n\n")
		for i, choice := range model.derivationMenuChoices {
			var line string
			if model.derivationCursor == i {
				line = SelectedItemStyle.Render(choice)
			} else {
				line = ListItemStyle.Render(choice)
			}

			valueStyle := InputValueStyle
			if model.derivationCursor == i && (choice == "Ponto x" || choice == "Passo h") && model.focus != common.FocusNone {
				valueStyle = FocusedInputStyle
			}

			switch choice {
			case "Filosofia":
				line += fmt.Sprintf(": %s", valueStyle.Render(model.selectedDerivationPhilosophy))
			case "Ordem do Erro":
				errorOrderDisplay := "N/A"
				for _, eo := range model.currentErrorOrderOptions {
					if eo.Value == model.selectedDerivationErrorOrder {
						errorOrderDisplay = eo.Display
						break
					}
				}
				line += fmt.Sprintf(": %s", valueStyle.Render(errorOrderDisplay))
			case "Ordem da Derivada":
				derivativeOrderDisplay := "N/A"
				for _, do := range model.derivativeOrderOptions {
					if do.Value == model.selectedDerivationOrder {
						derivativeOrderDisplay = do.Display
						break
					}
				}
				line += fmt.Sprintf(": %s", valueStyle.Render(derivativeOrderDisplay))
			case "Função":
				name := "Nenhuma"
				if model.selectedFunctionDef.Func != nil {
					name = model.selectedFunctionDef.Name
				}
				line += fmt.Sprintf(": %s", valueStyle.Render(name))
			case "Ponto x":
				inputX := model.currentX
				if model.focus == common.FocusX {
					inputX += CursorStyle.Render("_")
					line += fmt.Sprintf(": %s", FocusedInputStyle.Render(inputX))
				} else {
					line += fmt.Sprintf(": %s", InputValueStyle.Render(inputX))
				}
			case "Passo h":
				inputH := model.currentH
				if model.focus == common.FocusH {
					inputH += CursorStyle.Render("_")
					line += fmt.Sprintf(": %s", FocusedInputStyle.Render(inputH))
				} else {
					line += fmt.Sprintf(": %s", InputValueStyle.Render(inputH))
				}
			}
			s.WriteString(line + "\n")
		}
		s.WriteString(HelpStyle.Render("\n(Navegue com ↑/↓, 'Enter' para selecionar/editar, 'q' para voltar, Ctrl+C para sair)"))
		switch model.focus {
		case common.FocusX:
			s.WriteString(HelpStyle.Render("\n[EDITANDO Ponto x: Digite o valor e pressione Enter para confirmar]"))
		case common.FocusH:
			s.WriteString(HelpStyle.Render("\n[EDITANDO Passo h: Digite o valor e pressione Enter para confirmar]"))
		}

	case common.StateSelectPhilosophy:
		s.WriteString(TitleStyle.Render("Selecione a Filosofia:"))
		s.WriteString("\n\n")
		for i, phil := range model.philosophyOptions {
			s.WriteString(RenderListItem(phil, model.selectionCursor == i))
			s.WriteString("\n")
		}
		s.WriteString(HelpStyle.Render("\n('Enter' para confirmar, 'q' para voltar)"))

	case common.StateSelectErrorOrder:
		s.WriteString(TitleStyle.Render("Selecione a Ordem do Erro:"))
		s.WriteString("\n\n")
		if len(model.currentErrorOrderOptions) == 0 {
			s.WriteString(ErrorStyle.Render("Nenhuma ordem de erro disponível para a filosofia selecionada."))
		} else {
			for i, eo := range model.currentErrorOrderOptions {
				s.WriteString(RenderListItem(eo.Display, model.selectionCursor == i))
				s.WriteString("\n")
			}
		}
		s.WriteString(HelpStyle.Render("\n('Enter' para confirmar, 'q' para voltar)"))

	case common.StateSelectDerivativeOrder:
		s.WriteString(TitleStyle.Render("Selecione a Ordem da Derivada:"))
		s.WriteString("\n\n")
		for i, do := range model.derivativeOrderOptions {
			s.WriteString(RenderListItem(do.Display, model.selectionCursor == i))
			s.WriteString("\n")
		}
		s.WriteString(HelpStyle.Render("\n('Enter' para confirmar, 'q' para voltar)"))

	case common.StateSelectFunction:
		s.WriteString(TitleStyle.Render("Selecione a Função:"))
		s.WriteString("\n\n")
		for i, fd := range model.functionDefinitions {
			s.WriteString(RenderListItem(fd.Name, model.selectionCursor == i))
			s.WriteString("\n")
		}
		s.WriteString(HelpStyle.Render("\n('Enter' para confirmar, 'q' para voltar)"))

	case common.StateResult:
		if model.err != nil {
			s.WriteString(ErrorStyle.Render(fmt.Sprintf("Erro: %v\n", model.err)))
		} else {
			s.WriteString(ResultStyle.Render(fmt.Sprintf("Resultado: %s\n", model.result)))
		}
		s.WriteString(HelpStyle.Render("\n(Pressione 'Enter' para voltar ao menu, 'q' ou Ctrl+C para sair)"))

	case common.StateIntegrationMenu:
		s.WriteString(TitleStyle.Render("Menu de Integração Numérica:"))
		s.WriteString("\n\n")
		for i, choiceText := range model.integrationMenuChoices {
			var line string
			isSelected := model.integrationCursor == i

			if isSelected {
				line = SelectedItemStyle.Render(choiceText)
			} else {
				line = ListItemStyle.Render(choiceText)
			}

			valueStr := ""
			currentValueStyle := InputValueStyle // Default style for values

			// Determine if the current item is an input field and apply focused style if necessary
			isCurrentChoiceFocusableInput := (choiceText == "Limite Inferior (a)" && model.focus == common.FocusIntegrationA) ||
				(choiceText == "Limite Superior (b)" && model.focus == common.FocusIntegrationB) ||
				(choiceText == "Num de Subintervalos/Ordem (n)" && model.focus == common.FocusIntegrationN)

			if isSelected && isCurrentChoiceFocusableInput {
				currentValueStyle = FocusedInputStyle
			}
			//else if isSelected && model.focus != common.FocusNone &&
			//	(choiceText == "Limite Inferior (a)" || choiceText == "Limite Superior (b)" || choiceText == "Num de Subintervalos/Ordem (n)") {
			// If another input is focused, but this one is selected (cursor is on it), keep standard value style
			// This case might be redundant if selection follows focus, but good for clarity
			//}

			switch choiceText {
			case "Método":
				valueStr = model.integrationMethodsOptions[model.selectedIntegrationMethod].Display
			case "Função":
				if model.selectedFunctionDef.Func != nil {
					valueStr = model.selectedFunctionDef.Name
				} else {
					valueStr = "Nenhuma"
				}
			case "Limite Inferior (a)":
				valueStr = model.currentA
				if model.focus == common.FocusIntegrationA {
					valueStr += CursorStyle.Render("_")
				}
			case "Limite Superior (b)":
				valueStr = model.currentB
				if model.focus == common.FocusIntegrationB {
					valueStr += CursorStyle.Render("_")
				}
			case "Num de Subintervalos/Ordem (n)":
				valueStr = model.currentN
				if model.focus == common.FocusIntegrationN {
					valueStr += CursorStyle.Render("_")
				}
			}

			if choiceText == "Calcular" || choiceText == "Voltar" {
				s.WriteString(line + "\n")
			} else {
				// For items that display a value, apply the determined style to the value part
				s.WriteString(fmt.Sprintf("%s: %s\n", line, currentValueStyle.Render(valueStr)))
			}
		}
		s.WriteString(HelpStyle.Render("\n(Navegue com ↑/↓, 'Enter' para selecionar/editar, 'q' para voltar, Ctrl+C para sair)"))
		switch model.focus {
		case common.FocusIntegrationA:
			s.WriteString(HelpStyle.Render("\n[EDITANDO Limite Inferior (a): Digite o valor e pressione Enter para confirmar]"))
		case common.FocusIntegrationB:
			s.WriteString(HelpStyle.Render("\n[EDITANDO Limite Superior (b): Digite o valor e pressione Enter para confirmar]"))
		case common.FocusIntegrationN:
			s.WriteString(HelpStyle.Render("\n[EDITANDO Num de Subintervalos/Ordem (n): Digite o valor e pressione Enter para confirmar]"))
		}

	case common.StateSelectIntegrationMethod:
		s.WriteString(TitleStyle.Render("Selecione o Método de Integração:"))
		s.WriteString("\n\n")
		for i, methodName := range model.integrationMethodsOptions {
			s.WriteString(RenderListItem(methodName.Display, model.selectionCursor == i))
			s.WriteString("\n")
		}
		s.WriteString(HelpStyle.Render("\n('Enter' para confirmar, 'q' para voltar)"))
	}

	return DocStyle.Render(s.String())
}
