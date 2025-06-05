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
		s.WriteString(TitleStyle.Render("Integração Numérica"))
		s.WriteString("\n\n")
		s.WriteString(model.result) //
		s.WriteString(HelpStyle.Render("\n\n(Pressione 'Enter' para voltar ao menu principal, 'q' ou Ctrl+C para sair)"))
	}

	return DocStyle.Render(s.String())
}
